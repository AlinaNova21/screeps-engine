package main

import (
	"github.com/ags131/screeps-engine/common/rpc"
 "encoding/json"
 "log"
 "fmt"
)

func main() {
	s := rpc.NewServer()
	s.Methods["dbEnvGet"] = wrap(dbEnvGet)
	s.Methods["dbEnvSet"] = wrap(dbEnvSet)
	s.Listen(":3000")
}

func wrap(fn func([]interface{})(interface{},error)) rpc.Method {
	return func(f *rpc.Frame) (interface{}, error) {
		var raw interface{}
		err := json.Unmarshal(*f.Args, &raw)
		if err != nil {
			return nil, err
		}
		args := raw.([]interface{})
		return fn(args)			
	}
}

var fakeEnv = map[string]string{}

func dbEnvGet(args []interface{}) (interface{}, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("Wrong number of args: %d", len(args))
	}
	log.Printf("%v", args)
	return fakeEnv[args[0].(string)], nil
}

func dbEnvSet(args []interface{}) (interface{}, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("Wrong number of args: %d", len(args))
	}
	key, ok := args[0].(string)
	if !ok {
		return nil, fmt.Errorf("Key must be a string")
	}
	value, ok := args[1].(string)
	if !ok {
		return nil, fmt.Errorf("Value must be a string")
	}
	fakeEnv[key] = value
	return true, nil
}

type testResp struct {
	Hello string
}