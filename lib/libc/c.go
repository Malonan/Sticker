package c

/*
  Copyright 2023 Malonan & 3JoB

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.
*/

import (
	"context"

	"github.com/go-redis/redis/v9"

	"sticker/lib/libF"
)

var (
	rd *redis.Client
	kc = libF.F
	ctx = context.Background()
)

func init() {
	rd = redis.NewClient(&redis.Options{
		Addr:     kc().String("cache.addr"),
		Password: kc().String("cache.pwd"), // no password set
		DB:       kc().Int("cache.db"),     // use default DB
	})
	if err := rd.Conn().Ping(ctx).Err(); err != nil {
		panic(err)
	}
}

func Do() *redis.Client {
	return rd
}
