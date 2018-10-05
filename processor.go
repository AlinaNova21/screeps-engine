package main

import (
	"fmt"
	"github.com/ags131/screeps-engine/driver"
	_ "github.com/ags131/screeps-engine/intents"
	"github.com/ags131/screeps-engine/intents/handler"
)

func ProcessRoom(room string) {
	defer finishRoom(room)
	movement := handler.MakeMovement()
	objectsBulk := driver.GetObjectsBulk()
	gameTime := driver.GetGameTime()
	terrain, err := driver.GetTerrain(room)
	if err != nil {
		return
	}

	roomIntents, err := driver.GetIntentsByRoom(room)
	if err != nil {
		return
	}
	objects, err := driver.GetObjectsByRoom(room)
	if err != nil {
		panic(err)
	}
	keyedObjects := make(map[string]driver.Object)

	for _, obj := range objects {
		keyedObjects[obj.ID.Hex()] = obj
	}

	scope := handler.IntentScope{
		Terrain:     terrain,
		Movement:    &movement,
		ObjectsBulk: objectsBulk,
		GameTime:    gameTime,
	}

	for userId, userIntents := range roomIntents {
		user, err := driver.GetUser(userId)
		if err != nil {
			continue
		}
		fmt.Println("User: ", user.Username)
		for objId, intents := range userIntents {
			obj := keyedObjects[objId]
			fmt.Println("Object Type: ", obj.Type)
			fmt.Println("Intents: ", intents)
			for intentType, intent := range intents {
				handler.ProcessIntent(obj, intentType, intent, scope)
			}
			movement.Execute(obj, scope)
		}
	}
	for _, obj := range objects {
		handler.ProcessIntent(obj, "tick", nil, scope)
	}
	_, err = objectsBulk.Run()
	if err != nil {
		panic(err)
	}
}

func finishRoom(room string) {
	q := driver.GetQueue("rooms")
	err := q.Done(room)
	if err != nil {
		panic(err)
	}
}
