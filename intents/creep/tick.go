package creep

import (
	"github.com/ags131/screeps-engine/driver"
	"github.com/ags131/screeps-engine/intents/handler"
	"github.com/globalsign/mgo/bson"
)

func init() {
	handler.RegisterIntent("creep", "tick", Tick)
}

func Tick(object driver.Object, intent handler.IntentData, scope handler.IntentScope) {
	update := make(map[string]interface{})
	doUpdate := false
	if object.Props["spawning"] == true {
		return
	}
	if object.Props["fatigue"].(int) > 0 {
		update["fatigue"] = object.Props["fatigue"].(int) - 1
		doUpdate = true
	}
	if doUpdate {
		query := bson.M{"_id": object.ID}
		doc := bson.M{"$set": update}
		scope.ObjectsBulk.Update(query, doc)
	}
}
