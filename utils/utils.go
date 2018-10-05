package utils

import (
	"github.com/ags131/screeps-engine/driver"
	"strconv"
)

type Point struct {
	x int
	y int
}

var offsetsByDirection [8][2]int

func init() {
	offsetsByDirection = [8][2]int{
		[2]int{0, -1},
		[2]int{1, -1},
		[2]int{1, 0},
		[2]int{1, 1},
		[2]int{0, 1},
		[2]int{-1, 1},
		[2]int{-1, 0},
		[2]int{-1, -1},
	}

}

func GetOffsetByDirection(dir int) (int, int) {
	x := offsetsByDirection[dir-1][0]
	y := offsetsByDirection[dir-1][1]
	return x, y
}

func CheckTerrain(terrain driver.Terrain, x int, y int, mask int) bool {
	code, err := strconv.ParseInt(string([]rune(terrain.Terrain)[(y*50)+x]), 10, 0)
	if err != nil {
		panic(err)
	}
	return (int(code) & mask) > 0
}
