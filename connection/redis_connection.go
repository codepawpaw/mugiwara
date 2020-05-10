package connection

import (
	"sync"

	"github.com/go-redis/redis"
)

type Redis struct {
	*redis.Client
}

var redisInstance *Redis
var once sync.Once

func GetRedis() *Redis {
	once.Do(func() {
		redisInstance = &Redis{
			Client: redis.NewClient(&redis.Options{
				Addr:     "redis:6379",
				Password: "", // no password set
				DB:       0,  // use default DB
			}),
		}
	})

	return redisInstance
}
