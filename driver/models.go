package driver

import (
	"github.com/globalsign/mgo/bson"
)

type User struct {
	ID            bson.ObjectId "_id,omitempty"
	Username      string        "username,omitempty"
	UsernameLower string        "usernameLower,omitempty"
	CPU           int           "cpu,omitempty"
}

type Room struct {
	ID     string "_id,omitempty"
	Status string "status,omitempty"
	Active bool   "active,omitempty"
}

type Object struct {
	ID    bson.ObjectId          "_id,omitempty"
	Type  string                 "type"
	Room  string                 "room"
	X     int                    "x"
	Y     int                    "y"
	Props map[string]interface{} ",inline"
}

type Terrain struct {
	ID      bson.ObjectId "_id,omitempty"
	Type    string        "type"
	Room    string        "room"
	Terrain string        "terrain"
}

type BodyPart struct {
	Type string
	Hits int
}

type Energy struct {
	Energy         int
	EnergyCapacity int
}

type Creep struct {
	Object
	Resources
	Name      string
	Body      []BodyPart
	ActionLog ActionLog
	Fatigue   int
	AgeTime   int
	InterRoom InterRoom
}

type Mineral struct {
	Object
	MineralType   string
	Density       int
	MineralAmount int
}

type Resources struct {
	EnergyCapacity int "energyCapacity"
	Energy         int "energy"
	Power          int "power"
	G              int
	H              int
	K              int
	L              int
	O              int
	U              int
	X              int
	Z              int
	GH             int
	GO             int
	KH             int
	KO             int
	LH             int
	LO             int
	UH             int
	OH             int
	UL             int
	UO             int
	ZH             int
	ZK             int
	ZO             int
	GH2O           int
	LHO            int
	ZH2O           int
	KHO2           int
	UH2O           int
	LH2O           int
	GHO2           int
	ZHO2           int
	UHO2           int
	KH2O           int
	XUHO2          int
	XKHO2          int
	XGHO2          int
	XLHO2          int
	XUH2O          int
	XZHO2          int
	XKH2O          int
	XZH2O          int
	XLH2O          int
	XGH2O          int
}

type Source struct {
	Object
	Energy
	TicksToRegeneration  int
	InvaderHarvested     int
	NextRegenerationTime int
}

type Controller struct {
	Object
	Level             int
	Sign              Sign
	Reservation       Reservation
	Progress          int
	ProgressTotal     int
	DowngradeTime     int
	SafeMode          int
	SafeModeAvailable int
	SafeModeCooldown  int
	UpgradeBlocked    bool
}

type Sign struct {
}

type Reservation struct {
}

type ActionLog struct {
}

type InterRoom struct {
}

type Structure struct {
	Hits    int
	HitsMax int
}

type UserOwned struct {
	User               string
	NotifyWhenAttacked bool
}

type Spawn struct {
	Object
	UserOwned
	Energy
	Name          string `bson:"name,omitempty"`
	NextSpawnTime int
	Spawning      struct{}
	Off           bool
}

type Container struct {
	Object
	Energy
	Resources
	NextDecayTime int
}

type Lab struct {
	Object
	UserOwned
	CooldownTime    int
	MineralCapacity int
}

type PowerSpawn struct {
	Object
	UserOwned
	Energy        int
	Capacity      int
	PowerCapacity int
}

type Observer struct {
	Object
	UserOwned
	ObserveRoom string
}

type ConstructionSite struct {
	Object
	UserOwned
	StructureType string
	Progress      int
	ProgressTotal int
}

type Nuke struct {
	Object
	LandTime       int
	LaunchRoomName string
}

type Tombstone struct {
	Object
	DecayTime        int
	DeathTime        int
	CreepId          string
	CreepName        string
	CreepTicksToLive int
	CreepBody        []BodyPart
	CreepSaying      string
}

type Terminal struct {
	Object
	UserOwned
}
