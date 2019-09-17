package creep

import (
	"fmt"

	"github.com/ags131/screeps-engine/driver"
	"github.com/ags131/screeps-engine/game"
	"github.com/ags131/screeps-engine/intents/handler"
	"github.com/ags131/screeps-engine/utils"
)

func init() {
	handler.RegisterIntent("creep", "move", Move)
}

func Move(object *game.GameObject, intent handler.IntentData, scope handler.IntentScope) {
	if object.Props["spawning"] == true {
		return
	}
	if object.Props["fatigue"].(int) > 0 {
		return
	}
	body := object.Props["body"].([]interface{})
	canMove := false
	for _, rbp := range body {
		bp := rbp.(map[string]interface{})
		canMove = canMove || (bp["type"] == "move" && bp["hits"].(int) > 0)
	}
	if intent["direction"] == nil {
		return
	}
	dir := int(intent["direction"].(float64))
	dx, dy := utils.GetOffsetByDirection(dir)
	if object.X+dx < 0 || object.X+dx > 49 || object.Y+dy < 0 || object.Y+dy > 49 {
		return
	}
	targetObjects := []driver.Object{}
	for _, obj := range scope.RoomObjects {
		if obj.X == object.X && obj.Y == object.Y {
			targetObjects = append(targetObjects, obj)
		}
	}
	// TODO Check for Obstacles, etc
	scope.Movement.Add(object, dx, dy)
	fmt.Println("move", intent["direction"], canMove)
}
