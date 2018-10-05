package redis

import (
	"github.com/go-redis/redis"
)

type ENV struct {
	Client *redis.Client
	pub    *redis.Client
	sub    *redis.Client
}

func (env *ENV) Connect(addr string) error {
	opts := &redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	}
	client := redis.NewClient(opts)
	pub := redis.NewClient(opts)
	sub := redis.NewClient(opts)
	env.Client = client
	env.pub = pub
	env.sub = sub
	return nil
}

func (env *ENV) Close() {
	env.Client.Close()
	env.pub.Close()
	env.sub.Close()
}

func (env *ENV) Publish(channel string, message interface{}) (int64, error) {
	res, err := env.pub.Publish(channel, message).Result()
	return res, err
}

func (env *ENV) Subscribe(channel string, message interface{}) *redis.PubSub {
	pubsub := env.pub.Subscribe(channel)
	_, err := pubsub.Receive()
	if err != nil {
		panic(err)
	}
	return pubsub
}
