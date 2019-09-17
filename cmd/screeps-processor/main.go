package main

import (
	"github.com/ags131/screeps-engine/driver"
	"github.com/ags131/screeps-engine/game"
)

var world *game.World

func main() {
	err := driver.Connect("processor")
	if err != nil {
		panic(err)
	}
	defer driver.Close()

	world = game.NewWorld()
	rooms, err := driver.GetRooms()
	if err != nil {
		panic(err)
	}
	for _, r := range rooms {
		room := game.NewRoom(r.ID, game.RoomStatus(r.Status))
		room.Active = r.Active
		world.Rooms[room.Name] = *room
	}
	q := driver.GetQueue("rooms")
	for {
		item, err := q.Fetch()
		if err != nil {
			panic(err)
		}
		go ProcessRoom(item)

	}
}
