package driver

import (
	"encoding/json"
	"github.com/ags131/screeps-engine/mongo"
	"github.com/ags131/screeps-engine/redis"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"strconv"
)

var db mongo.DB
var env redis.ENV

type Bulk = mgo.Bulk

func Connect(module string) error {
	db = mongo.DB{}
	err := db.Connect("127.0.0.1", "screeps")
	if err != nil {
		panic(err)
	}
	env = redis.ENV{}
	env.Connect("127.0.0.1:6379")
	return nil
}

func Close() {
	db.Close()
	env.Close()
}

func GetIntentsByRoom(room string) (map[string]map[string]map[string]map[string]interface{}, error) {
	result := make(map[string]map[string]map[string]map[string]interface{})
	res := env.Client.HGetAll("roomIntents:" + room)
	err := res.Err()
	if err != nil {
		return result, err
	}
	val := res.Val()
	for id, envIntents := range val {
		var tmp map[string]map[string]map[string]interface{}
		json.Unmarshal([]byte(envIntents), &tmp)
		result[id] = tmp
	}
	return result, err
}

func GetUsers() ([]User, error) {
	var results []User
	err := db.DB.C("users").Find(bson.M{}).All(&results)
	return results, err
}

func GetUser(id string) (User, error) {
	result := User{}
	err := db.DB.C("users").Find(bson.M{"_id": id}).One(&result)
	return result, err
}

func GetRooms() ([]Room, error) {
	var results []Room
	err := db.DB.C("rooms").Find(bson.M{}).All(&results)
	return results, err
}

func GetRoom(id string) (Room, error) {
	result := Room{}
	err := db.DB.C("rooms").Find(bson.M{"_id": id}).One(&result)
	return result, err
}

func GetObjectsByRoom(room string) ([]Object, error) {
	var results []Object
	err := db.DB.C("rooms.objects").Find(bson.M{"room": room}).All(&results)
	return results, err
}

func GetTerrain(room string) (Terrain, error) {
	var result Terrain
	err := db.DB.C("rooms.terrain").Find(bson.M{"room": room}).One(&result)
	return result, err
}

func GetGameTime() (int64, error) {
	res, err := env.Client.Get("gameTime").Result()
	if err != nil {
		return 0, err
	}
	num, err := strconv.ParseInt(res, 10, 64)
	return num, err
}

func GetObjectsBulk() *Bulk {
	return db.DB.C("rooms.objects").Bulk()
}

// func UpdateObject(id bson.ObjectId, data interface{}) error {
//err := db.DB.C("rooms.objects").Update(bson.M{"_id": id})
// }
