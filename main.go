package main

import (
	"github.com/ags131/screeps-engine/driver"
)

func main() {
	err := driver.Connect("processor")
	if err != nil {
		panic(err)
	}
	defer driver.Close()

	q := driver.GetQueue("rooms")
	for {
		item, err := q.Fetch()
		if err != nil {
			panic(err)
		}
		go ProcessRoom(item)
		// ProcessRoom(item)
	}
}
