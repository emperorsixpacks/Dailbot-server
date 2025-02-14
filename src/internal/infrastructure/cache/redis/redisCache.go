package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/emperorsixpacks/dailbot/src/pkg/logger"
	"github.com/emperorsixpacks/dailbot/src/pkg/utils"
	"github.com/redis/go-redis/v9"
)

type KeyPath interface{}

func NewRedisCache(appConfig utils.DBSettings) (*Redis, error) {

	redisConnAddr := fmt.Sprintf("%s:%s", appConfig.Host, appConfig.Port)
	intDB, err := strconv.Atoi(appConfig.DB)
	if err != nil {
		intDB = 0
	}
	options := &redis.Options{
		Addr:     redisConnAddr,
		Password: appConfig.Password,
		DB:       intDB,
	}
	client := redis.NewClient(options)
	err = client.Ping(context.Background()).Err()
	if err != nil {
		logger.DefaultLogger.Printf("could not connect on %s \n%v", redisConnAddr, err)
		return nil, err
	}
	return &Redis{client}, nil
}

func mapToStruct(i interface{}, o interface{}) error {
	newStrVal, err := json.Marshal(i)
	if err != nil {
		return err
	}
	if err = json.Unmarshal([]byte(newStrVal), o); err != nil {
		return err
	}
	return nil
}
func returnJSONKey(key KeyPath) (string, error) {
	if itm, ok := key.(int); ok {
		if itm == 0 {
			key = []string{"$"}
		}
	}
	if str, ok := key.([]string); ok {
		return strings.Join(str, "."), nil
	}
	// log and crash server
	return "", fmt.Errorf("Invalid Key:%v", key)

}

// TODO look into making this a singleton
type Redis struct {
	rdb *redis.Client
}

func (this *Redis) clearDB() error {
	if err := this.rdb.FlushAll(context.Background()).Err(); err != nil {
		return err
	}
	return nil
}

// TODO try to make this simpler
func (this Redis) GetJSON(item string, k KeyPath, o interface{}) error {
	// NOTE this works
	val, err := this.getJSON(item, k)
	if err != nil {
		return err
	}
	strMapping, ok := val.([]interface{})
	if !ok {
		return errors.New("internal error")
	}
	err = mapToStruct(strMapping[0], &o)
	if err != nil {
		return err
	}
	return nil
}

// this is a low level method, from here, we can perform things like deleting a single key or updating a single key
func (this Redis) getJSON(item string, key KeyPath) (interface{}, error) {
	_key, err := returnJSONKey(key)
	if err != nil {
		return nil, err
	}
	val, err := this.rdb.JSONGet(context.Background(), item, _key).Expanded()
	if err != nil {
		return nil, err
	}
	return val, nil
}

// we can even expand this further to get the data in a nestad json
// let us go ahead now and create some hidden methods to handle this
func (this Redis) SetJSON(item string, key KeyPath, value interface{}) error {
	if err := this.setJSON(item, key, value); err != nil {
		return err
	}

	return nil
}

// TODO look into making some of thise public
func (this Redis) setJSON(item string, key KeyPath, value interface{}) error {
	val, err := json.Marshal(value)
	if err != nil {
		return err
	}
	// TODO put this into a function
	_key, err := returnJSONKey(key)
	if err != nil {
		return err
	}
	err = this.rdb.JSONSet(context.Background(), item, _key, val).Err()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// Keeing it like this for now, later if needed we may need to go into nested objets to delete specific keys, but that should be from the client
func (this Redis) DeleteJSON(item string, key KeyPath, value interface{}) error {
	_key, err := returnJSONKey(key)
	if err != nil {
		return err
	}
	err = this.rdb.JSONDel(context.Background(), item, _key).Err()
	if err != nil {
		// log error here
		return err
	}
	// log here
	return nil
}
