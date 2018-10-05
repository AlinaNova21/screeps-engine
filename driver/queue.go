package driver

import (
	"fmt"
	"time"
)

type Queue struct {
	name       string
	pending    string
	processing string
}

func GetQueue(name string) Queue {
	q := Queue{}
	q.name = name
	q.pending = name + "Pending"
	q.processing = name + "Processing"
	return q
}

func (q *Queue) Fetch() (string, error) {
	for {
		res := env.Client.RPopLPush(q.pending, q.processing)
		err := res.Err()
		item := res.Val()
		fmt.Sprintln("Item: %T", item)
		if err != nil && err.Error() != "redis: nil" {
			return item, err
		}
		if item != "" {
			return item, err
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func (q *Queue) Done(id string) error {
	_, err := env.Client.LRem(q.processing, 0, id).Result()
	if err != nil {
		return err
	}
	err = q.emit("done", nil)
	return err
}

func (q *Queue) emit(channel string, data interface{}) error {
	fullChannel := fmt.Sprintf("queue_%s_%s", q.name, channel)
	_, err := env.Publish(fullChannel, data)
	return err
}
