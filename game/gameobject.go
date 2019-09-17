package game

type Store map[string]float64

type GameObjectType string

const (
	Spawn            GameObjectType = "spawn"
	Creep            GameObjectType = "creep"
	Controller       GameObjectType = "controller"
	ConstructionSite GameObjectType = "construction_site"
	Road             GameObjectType = "road"
	PowerCreep       GameObjectType = "power_creep"
)

func NewGameObject(id string, objectType GameObjectType, room string, x, y int8) *GameObject {
	return &GameObject{
		ID:    id,
		Type:  objectType,
		Room:  room,
		X:     x,
		Y:     y,
		Props: make(map[string]interface{}),
	}
}

type GameObject struct {
	ID            string
	Room          string
	X, Y          int8
	Type          GameObjectType
	Props         map[string]interface{}
	Store         Store
	StoreCapacity int64
}
