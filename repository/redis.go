// repository/redis.go
package repository

import (
	"context"
	"encoding/json"
	"golang-redis/models"
	
	"github.com/go-redis/redis/v8"
)

var Cache *redis.Client

func SetupRedis() {
	Cache = redis.NewClient(&redis.Options{
		// docker-compose.ymlに指定したservice名+port
		Addr: "redis:6379",
		DB:   0,
	})
}

func GetUserList(uuid string) ([]models.User, error) {
	data, err := Cache.Get(context.Background(), uuid).Result()
	if err != nil {
		return nil, err
	}

	userList := new([]models.User)
	err = json.Unmarshal([]byte(data), userList)
	if err != nil {
		return nil, err
	}
	return *userList, nil
}