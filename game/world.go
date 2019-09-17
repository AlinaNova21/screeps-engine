package game

type World struct {
	GameTime int64
	Rooms    map[string]Room
}

func NewWorld() *World {
	return &World{
		GameTime: 0,
		Rooms:    make(map[string]Room, 0),
	}
}
