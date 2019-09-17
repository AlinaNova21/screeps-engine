package common

import "github.com/ags131/screeps-engine/game"

var (
	OBSTACLE_OBJECT_TYPES    = []string{"spawn", "creep", "source", "mineral", "controller", "constructedWall", "extension", "link", "storage", "tower", "observer", "powerSpawn", "powerBank", "lab", "terminal", "nuker"}
	TERRAIN_MASK_WALL        = 1
	TERRAIN_MASK_SWAMP       = 2
	TERRAIN_MASK_LAVA        = 4
	CARRY_CAPACITY           = 50
	ROAD_WEAROUT             = 1
	ROAD_WEAROUT_POWER_CREEP = 100
	BOOSTS                   = map[game.BodyPartType]map[string]map[string]float64{
		game.Work: {
			"UO": {
				"harvest": 3,
			},
			"UHO2": {
				"harvest": 5,
			},
			"XUHO2": {
				"harvest": 7,
			},
			"LH": {
				"build":  1.5,
				"repair": 1.5,
			},
			"LH2O": {
				"build":  1.8,
				"repair": 1.8,
			},
			"XLH2O": {
				"build":  2,
				"repair": 2,
			},
			"ZH": {
				"dismantle": 2,
			},
			"ZH2O": {
				"dismantle": 3,
			},
			"XZH2O": {
				"dismantle": 4,
			},
			"GH": {
				"upgradeController": 1.5,
			},
			"GH2O": {
				"upgradeController": 1.8,
			},
			"XGH2O": {
				"upgradeController": 2,
			},
		},
		game.Attack: {
			"UH": {
				"attack": 2,
			},
			"UH2O": {
				"attack": 3,
			},
			"XUH2O": {
				"attack": 4,
			},
		},
		game.RangedAttack: {
			"KO": {
				"rangedAttack":     2,
				"rangedMassAttack": 2,
			},
			"KHO2": {
				"rangedAttack":     3,
				"rangedMassAttack": 3,
			},
			"XKHO2": {
				"rangedAttack":     4,
				"rangedMassAttack": 4,
			},
		},
		game.Heal: {
			"LO": {
				"heal":       2,
				"rangedHeal": 2,
			},
			"LHO2": {
				"heal":       3,
				"rangedHeal": 3,
			},
			"XLHO2": {
				"heal":       4,
				"rangedHeal": 4,
			},
		},
		game.Carry: {
			"KH": {
				"capacity": 2,
			},
			"KH2O": {
				"capacity": 3,
			},
			"XKH2O": {
				"capacity": 4,
			},
		},
		game.Move: {
			"ZO": {
				"fatigue": 2,
			},
			"ZHO2": {
				"fatigue": 3,
			},
			"XZHO2": {
				"fatigue": 4,
			},
		},
		game.Tough: {
			"GO": {
				"damage": .7,
			},
			"GHO2": {
				"damage": .5,
			},
			"XGHO2": {
				"damage": .3,
			},
		},
	}
)
