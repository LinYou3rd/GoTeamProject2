package raise_project

import (
	"github.com/go-redis/redis"
)

var db *redis.Client

func Init() {
	db = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       1,
	})

	_, err := db.Ping().Result()
	if err != nil {
		panic(err)
	}

}
