package server

import (
	"crud-challenge/config"
	"crud-challenge/storage"
	distributedlock "crud-challenge/utils/distributed-lock"
	"github.com/go-redis/redis/v8"
	"github.com/golobby/container/v3"
)

func InitDependencies() {
	err := storage.WithStorage()
	if err != nil {
		panic(err)
	}

	withDistributedLock()
}

func withDistributedLock() {
	container.Singleton(func() distributedlock.DistributedLock {
		redisClient := redis.NewClient(&redis.Options{
			Addr: config.Config.Redis.Address,
		})

		return distributedlock.New(redisClient)
	})
}