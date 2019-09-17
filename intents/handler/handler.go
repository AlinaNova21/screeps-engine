package handler

import (
	"github.com/ags131/screeps-engine/driver"
	"github.com/ags131/screeps-engine/game"
)

type IntentData = map[string]interface{}
type IntentFunc = func(*game.GameObject, IntentData, IntentScope)

var objectTypes map[game.GameObjectType]map[string]IntentFunc

type IntentScope struct {
	World       *game.World
	User        *game.User
	Room        *game.Room
	Movement    *Movement
	ObjectsBulk *driver.Bulk
}

func init() {
	objectTypes = make(map[game.GameObjectType]map[string]IntentFunc, 0)
}

func RegisterIntent(objectType game.GameObjectType, intent string, fn IntentFunc) {
	if objectTypes[objectType] == nil {
		objectTypes[objectType] = make(map[string]IntentFunc, 0)
	}
	objectTypes[objectType][intent] = fn
}

func ProcessIntent(object *game.GameObject, intentType string, intent IntentData, scope IntentScope) {
	if objectTypes[object.Type] != nil && objectTypes[object.Type][intentType] != nil {
		fn := objectTypes[object.Type][intentType]
		fn(object, intent, scope)
	}
}
