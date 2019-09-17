package creep

import (
	"fmt"
	"math"

	"github.com/ags131/screeps-engine/game"
	"github.com/ags131/screeps-engine/intents/handler"
)

func init() {
	handler.RegisterIntent("creep", "upgradeController", UpgradeController)
}

func UpgradeController(object *game.GameObject, intent handler.IntentData, scope handler.IntentScope) {
	if object.Props["spawning"] == true {
		return
	}
	if object.Props["fatigue"].(int) > 0 || object.Props["energy"].(int) <= 0 {
		return
	}

	target, ok := scope.RoomObjects[intent["id"]]
	if !ok || target.Type != "controller" {
		return
	}
	if math.Abs(target.X-object.X) > 3 || math.Abs(target.Y-object.Y) > 3 {
		return
	}
	if target.Props["level"].(int) == 0 || target.Props["user"] != object.Props["user"] {
		return
	}
	if target.Props["upgradeBlocked"].(int) > 0 && target.Props["upgradeBlocked"].(int) > scope.GameTime {
		return
	}

	// TODO: Complete upgradeController coding

	fmt.Println("upgradeController")
}
