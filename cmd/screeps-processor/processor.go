package main

import (
	"fmt"

	"github.com/ags131/screeps-engine/driver"
	"github.com/ags131/screeps-engine/game"
	_ "github.com/ags131/screeps-engine/intents"
	"github.com/ags131/screeps-engine/intents/handler"
)

func ProcessRoom(roomName string) {
	defer finishRoom(roomName)
	room := world.Rooms[roomName]
	movement := handler.MakeMovement()
	objectsBulk := driver.GetObjectsBulk()
	gameTime, err := driver.GetGameTime()
	if err != nil {
		return
	}
	terrain, err := driver.GetTerrain(roomName)
	if err != nil {
		return
	}

	roomIntents, err := driver.GetIntentsByRoom(roomName)
	if err != nil {
		return
	}
	objects, err := driver.GetObjectsByRoom(roomName)
	if err != nil {
		panic(err)
	}
	objs, err := driver.GetObjectsByRoom(room.Name)
	if err != nil {
		panic(err)
	}
	for _, o := range objs {
		obj := game.NewGameObject(o.ID.Hex(), game.GameObjectType(o.Type), room.Name, o.X, o.Y)
		obj.Props = o.Props
		room.Objects[obj.ID] = *obj
	}

	scope := handler.IntentScope{
		Movement:    &movement,
		ObjectsBulk: objectsBulk,
	}

	for userId, userIntents := range roomIntents {
		user, err := driver.GetUser(userId)
		// scope.User :=
		if err != nil {
			continue
		}
		fmt.Println("User: ", user.Username)
		for objId, intents := range userIntents {
			obj := room.Objects[objId]
			fmt.Println("Object Type: ", obj.Type)
			fmt.Println("Intents: ", intents)
			for intentType, intent := range intents {
				handler.ProcessIntent(&obj, intentType, intent, scope)
			}
			movement.Execute(&obj, scope)
			room.Objects[objId] = obj
		}
	}
	for _, obj := range objects {
		handler.ProcessIntent(&obj, "tick", nil, scope)
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
