package game

import (
	"fmt"
	"strconv"
	"strings"
)

type RoomStatus string

const (
	Normal      RoomStatus = "normal"
	OutOfBounds RoomStatus = "out of bounds"
)

type Room struct {
	Name    string
	Status  RoomStatus
	Active  bool
	X       int16
	Y       int16
	Objects map[string]GameObject
	Terrain *Terrain
}

func NewRoom(name string, status RoomStatus) *Room {
	x, y := RoomNameToXY(name)
	return &Room{
		Name:    name,
		Status:  status,
		Active:  true,
		X:       x,
		Y:       y,
		Objects: make(map[string]GameObject, 0),
		Terrain: NewTerrain(),
	}
}

func (r *Room) ObjectByType(t GameObjectType) *GameObject {
	for _, obj := range r.Objects {
		if obj.Type == t {
			return &obj
		}
	}
	return nil
}
func (r *Room) ObjectByTypeAt(t GameObjectType, x, y int8) *GameObject {
	for _, obj := range r.Objects {
		if obj.Type == t && obj.X == x && obj.Y == y {
			return &obj
		}
	}
	return nil
}

func (r *Room) ObjectsByType(t GameObjectType) []GameObject {
	ret := make([]GameObject, 0)
	for _, obj := range r.Objects {
		if obj.Type == t {
			ret = append(ret, obj)
		}
	}
	return ret
}

func RoomNameFromXY(x, y int16) string {
	dirX := "E"
	dirY := "S"
	if x < 0 {
		dirX = "W"
		x = -x - 1
	}
	if y < 0 {
		dirY = "N"
		y = -y - 1
	}
	return fmt.Sprintf("%s%d%s%d", dirX, x, dirY, y)
}

func RoomNameToXY(name string) (x, y int16) {
	name = strings.ToUpper(name)
	dirX := name[0]
	dirY := ' '
	x = 0
	y = 0
	for i, char := range name[1:] {
		if char == 'N' || char == 'S' {
			dirY = char
			xv, _ := strconv.ParseInt(name[1:i], 10, 16)
			x = int16(xv)
			yv, _ := strconv.ParseInt(name[i+1:], 10, 16)
			y = int16(yv)
			break
		}
	}
	if dirX == 'W' {
		x = -x - 1
	}
	if dirY == 'N' {
		y = -y - 1
	}
	return
}
