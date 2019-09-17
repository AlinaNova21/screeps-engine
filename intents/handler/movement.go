package handler

import (
	"fmt"
	"math"

	"github.com/ags131/screeps-engine/common"
	"github.com/ags131/screeps-engine/game"
)

type Movement struct {
	movements []movementObj
	matrix    [50][50][]game.GameObject
	objects   map[string]game.GameObject
}

type movementObj struct {
	Object game.GameObject
	Dx     int8
	Dy     int8
}

func MakeMovement() Movement {
	m := Movement{}
	m.movements = []movementObj{}
	m.matrix = [50][50][]game.GameObject{}
	m.objects = make(map[string]game.GameObject)
	return m
}

func (m *Movement) Add(object game.GameObject, dx int8, dy int8) {
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
	m.objects[object.ID] = object
}

func (m *Movement) Execute(object *game.GameObject, scope IntentScope) {
	var fatigueRate = 2
	move, ok := m.objects[object.ID]
	if !ok {
		return
	}
	if scope.Room.Terrain.Check(move.X, move.Y, game.Swamp) {
		fatigueRate = 10
	}

	road := scope.Room.ObjectByTypeAt(game.Road, move.X, move.Y)
	if road != nil {
		fatigueRate = 1
		ndt, ok := road.Props["nextDecayTime"].(int64)
		if !ok {
			ndt = 0
		}
		if object.Type == game.PowerCreep {
			ndt -= int64(common.ROAD_WEAROUT_POWER_CREEP)
		} else {
			ndt -= int64(common.ROAD_WEAROUT) * int64(len(object.Props["body"].([]game.BodyPartType)))
		}
		road.Props["nextDecayTime"] = ndt
	}

	roomController := scope.Room.ObjectByType(game.Controller)
	roomOwner := ""
	safeMode := false
	if roomController != nil {
		if uid, ok := roomController.Props["user"].(string); ok {
			roomOwner = uid
		}
		if v, ok := roomController.Props["safeMode"]; ok {
			safeMode = v.(int64) > scope.World.GameTime
		}
	}

	if roomController == nil || roomOwner == scope.User.ID || !safeMode {
		site := scope.Room.ObjectByTypeAt(game.ConstructionSite, move.X, move.Y)
		if site != nil {
			if site.Props["user"].(string) != scope.User.ID {
				delete(scope.Room.Objects, site.ID)
				if site.Props["progress"].(int) > 1 {
					// TODO: site stomping - create energy
					//  require('./creeps/_create-energy')(constructionSite.x, constructionSite.y,
					//    constructionSite.room, Math.floor(constructionSite.progress/2), 'energy', scope);
				}
			}
		}
	}

	fatigue := 0
	if object.Type == game.Creep {
		body := object.Props["body"].([]game.BodyPart)
		for _, part := range body {
			if part.Type != game.Move && part.Type != game.Carry {
				fatigue++
			}
		}
		fatigue += calcResourceWeight(object)
		fatigue *= fatigueRate
	}
	var toEdge = move.X == 0 || move.X == 49 || move.Y == 0 || move.Y == 49
	var fromEdge = object.X == 0 || object.X == 49 || object.Y == 0 || object.Y == 49
	if toEdge && !fromEdge {
		fatigue = 0
	}
	object.X = move.X
	object.Y = move.Y
	object.Props["fatigue"] = fatigue
	// TODO: Investigate _add-fatigue ref: https://github.com/screeps/engine/blob/ptr/src/processor/intents/movement.js#L242-L251
	fmt.Println(object.X, object.Y, move.X, move.Y)
}

func calcResourceWeight(object *game.GameObject) int {
	weight := 0
	totalCarry := float64(0)
	store := object.Props["store"].(game.Store)
	for _, amt := range store {
		totalCarry += amt
	}
	for _, part := range object.Props["body"].([]game.BodyPart) {
		if totalCarry == 0 {
			break
		}
		if part.Type != game.Carry || part.Hits == 0 {
			continue
		}
		boost := float64(1)
		if part.Boost != "" {
			boost = common.BOOSTS[game.Carry][part.Boost]["capacity"]
		}
		totalCarry -= math.Min(totalCarry, float64(common.CARRY_CAPACITY)*boost)
		weight++
	}
	return weight
}
