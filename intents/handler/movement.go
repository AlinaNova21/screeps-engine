package handler

import (
	"fmt"
	"github.com/ags131/screeps-engine/common"
	"github.com/ags131/screeps-engine/driver"
	"github.com/ags131/screeps-engine/utils"
	"github.com/globalsign/mgo/bson"
)

type Movement struct {
	movements []movementObj
	matrix    [50][50][]driver.Object
	objects   map[string]driver.Object
}

type movementObj struct {
	Object driver.Object
	Dx     int
	Dy     int
}

func MakeMovement() Movement {
	m := Movement{}
	m.movements = []movementObj{}
	m.matrix = [50][50][]driver.Object{}
	m.objects = make(map[string]driver.Object)
	return m
}

func (m *Movement) Add(object driver.Object, dx int, dy int) {
	nm := movementObj{
		Object: object, Dx: dx, Dy: dy,
	}
	m.movements = append(m.movements, nm)
	newX := object.X + dx
	newY := object.Y + dy
	if newX >= 50 {
		newX = 49
	}
	if newY >= 50 {
		newY = 49
	}
	if newX <= 0 {
		newX = 0
	}
	if newY <= 0 {
		newY = 0
	}
	m.matrix[newX][newY] = append(m.matrix[newX][newY], object)
	object.X = newX
	object.Y = newY
	m.objects[object.ID.Hex()] = object
}

func (m *Movement) Exec() {
	// fmt.Println("Exec", m.movements)

}

func (m *Movement) Execute(object driver.Object, scope IntentScope) {
	var fatigueRate = 2
	move, ok := m.objects[object.ID.Hex()]
	if !ok {
		return
	}
	if utils.CheckTerrain(scope.Terrain, move.X, move.Y, common.TERRAIN_MASK_SWAMP) {
		fatigueRate = 10
	}
	// TODO: Road decay/fatigue adjust
	// TODO: Construction Site Stomping
	// TODO: Fatigue calc
	fatigue := 1 * fatigueRate
	var toEdge = move.X == 0 || move.X == 49 || move.Y == 0 || move.Y == 49
	var fromEdge = object.X == 0 || object.X == 49 || object.Y == 0 || object.Y == 49
	if toEdge && !fromEdge {
		fatigue = 0
	}
	query := bson.M{"_id": object.ID}
	doc := bson.M{"$set": bson.M{"x": move.X, "y": move.Y, "fatigue": fatigue}}
	fmt.Println(object.X, object.Y, move.X, move.Y)
	scope.ObjectsBulk.Update(query, doc)
}
