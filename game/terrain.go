package game

type TerrainType byte

const (
	Plain TerrainType = 0
	Wall
	Swamp
)

const TerrainSize = 2500

type Terrain []TerrainType

func NewTerrain() *Terrain {
	t := make(Terrain, TerrainSize)
	return &t
}

func (t Terrain) Check(x, y int8, ty TerrainType) bool {
	return t[xyToI(x, y)]&ty > 0
}

func (t Terrain) Fill(tt TerrainType) {
	for i := 0; i < TerrainSize; i++ {
		t[i] = tt
	}
}

func xyToI(x, y int8) int16 {
	return int16(y*50 + x)
}

func xyFromI(i int16) (x, y int8) {
	x = int8(i % 50)
	y = int8(i / 50)
	return
}
