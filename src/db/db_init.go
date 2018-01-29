package db

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis"
)

const dbURLVariable = "REDIS_URL"

var url string

func Test(id string) {
	client := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}

	b := client.HSet("hash", "f", "v")
	if b.Err() != nil {
		log.Fatal(b.Err())
	}

	log.Println("set!")
}

func init() {
	url = os.Getenv(dbURLVariable)
	if url == "" {
		panic(fmt.Sprintf("Failed to read port from variable '%v'", dbURLVariable))
	}
}
