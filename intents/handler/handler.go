package handler

import "github.com/ags131/screeps-engine/driver"

type IntentData = map[string]interface{}
type IntentFunc = func(driver.Object, IntentData, IntentScope)

var objectTypes map[string]map[string]IntentFunc

type IntentScope struct {
	User        driver.User
	Terrain     driver.Terrain
	RoomObjects map[string]driver.Object
	Movement    *Movement
	ObjectsBulk *driver.Bulk
	GameTime    int
}

func init() {
	objectTypes = make(map[string]map[string]IntentFunc)
}

func RegisterIntent(object string, intent string, fn IntentFunc) {
	if objectTypes[object] == nil {
		objectTypes[object] = make(map[string]IntentFunc)
	}
	objectTypes[object][intent] = fn
}

func ProcessIntent(object driver.Object, intentType string, intent IntentData, scope IntentScope) {
	if objectTypes[object.Type] != nil && objectTypes[object.Type][intentType] != nil {
		fn := objectTypes[object.Type][intentType]
		fn(object, intent, scope)
	}
}
