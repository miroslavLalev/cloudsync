package internal

import (
	"fmt"
	"os"

	"github.com/miroslavLalev/cloudsync/src/db/errors"
	"github.com/miroslavLalev/cloudsync/src/objects"

	"github.com/go-redis/redis"
)

const dbURLVariable = "REDIS_URL"

const (
	usersPrefix = "users."
)

type RedisController struct {
	client *redis.Client
}

func GetController() *RedisController {
	return controller
}

var controller *RedisController

func (rc *RedisController) AddSpace(space *objects.Space) error {
	s, err := rc.GetSpace(space.Owner())
	if err != nil {
		return err
	}

	if s != "" {
		return &errors.SpaceAlreadyExists{}
	}

	cmd := rc.client.Set(usersPrefix+space.Owner(), space.Id(), 0)
	return cmd.Err()
}

func (rc *RedisController) GetSpace(user string) (string, error) {
	sc := rc.client.Get(usersPrefix + user)
	if sc.Err() == redis.Nil {
		// No key exists
		return "", nil
	}

	return sc.String(), sc.Err()
}

func (rc *RedisController) DeleteSpace(user string) error {
	ic := rc.client.Del(usersPrefix + user)
	if ic.Err() == redis.Nil {
		return nil
	}

	return ic.Err()
}

func init() {
	url := os.Getenv(dbURLVariable)
	if url == "" {
		panic(fmt.Sprintf("Failed to read port from variable '%v'", dbURLVariable))
	}

	controller = &RedisController{
		client: redis.NewClient(&redis.Options{
			Addr:     url,
			Password: "",
			DB:       0,
		}),
	}
}
