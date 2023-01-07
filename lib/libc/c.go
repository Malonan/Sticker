package c

import (
	"github.com/go-redis/redis/v9"

	"sticker/lib/libF"
)

var (
	rd *redis.Client
	kc = libF.F
)

func init() {
	rd = redis.NewClient(&redis.Options{
		Addr:     kc().String("cache.addr"),
		Password: kc().String("cache.pwd"), // no password set
		DB:       kc().Int("cache.db"),     // use default DB
	})
}

func Do() *redis.Client {
	return rd
}
