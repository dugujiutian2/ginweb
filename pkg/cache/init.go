package cache

import (
	"encoding/json"
	"errors"
	"github.com/go-redis/redis"
	"github.com/hero1s/ginweb/pkg/log"
	"time"
)

var (
	MemCache   Cache
	RedisCache Cache
	Redis      *redis.Client
)

func InitRedis(host, password string) bool {
	Redis = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0,
	})
	pong, err := Redis.Ping().Result()
	if err != nil {
		log.Error(err.Error())
		return false
	}
	log.Info("redis ping rep:%v", pong)
	return true
}

//发布消息
func PublishMessage(channel string, data interface{}) {
	Redis.Publish(channel, data)
}

//接受消息
func SubscribeMessage(channel string, msg_func func(msg *redis.Message)) {
	pubsub := Redis.Subscribe(channel)
	_, err := pubsub.Receive()
	if err != nil {
		return
	}
	ch := pubsub.Channel()
	for msg := range ch {
		log.Debug("接受到消息:%v-->%v", msg.Channel, msg.Payload)
		msg_func(msg)
	}
}

func InitCache(host, password, defaultKey string) error {
	var err error
	MemCache, err = NewCache("memory", `{"interval":60}`)
	if err != nil {
		return err
	}
	RedisCache, err = NewCache("redis",
		`{"conn":"`+host+`", "password":"`+password+`", "key":"`+defaultKey+`"}`)
	return err
}

func SetCache(cc Cache, key string, value interface{}, timeout time.Duration) error {
	data, err := EncodeJson(value)
	if err != nil {
		return err
	}
	if cc == nil {
		return errors.New("cc is nil")
	}

	defer func() {
		if r := recover(); r != nil {
			//fmt.Println("set cache error caught: %v\n", r)
			cc = nil
		}
	}()
	return Put(key, data, timeout)
}

func GetCache(cc Cache, key string, to interface{}) error {
	if cc == nil {
		return errors.New("cc is nil")
	}

	defer func() {
		if r := recover(); r != nil {
			//fmt.Println("get cache error caught: %v\n", r)
			cc = nil
		}
	}()

	data := Get(key)
	if data == nil {
		return errors.New("Cache不存在")
	}
	// log.Pinkln(data)
	return DecodeJson(data.([]byte), to)

}

func DelCache(cc Cache, key string) error {
	if cc == nil {
		return errors.New("cc is nil")
	}

	defer func() {
		if r := recover(); r != nil {
			//fmt.Println("get cache error caught: %v\n", r)
			cc = nil
		}
	}()

	return Delete(key)
}

func IsExist(cc Cache, key string) bool {
	if cc == nil {
		return false
	}
	defer func() {
		if r := recover(); r != nil {
			//fmt.Println("get cache error caught: %v\n", r)
			cc = nil
		}
	}()
	return IsExist(key)
}

// increase cached int value by key, as a counter.
func Incr(cc Cache, key string) error {
	if cc == nil {
		return errors.New("cc is nil")
	}
	defer func() {
		if r := recover(); r != nil {
			//fmt.Println("get cache error caught: %v\n", r)
			cc = nil
		}
	}()
	return Incr(key)
}

// decrease cached int value by key, as a counter.
func Decr(cc Cache, key string) error {
	if cc == nil {
		return errors.New("cc is nil")
	}
	defer func() {
		if r := recover(); r != nil {
			//fmt.Println("get cache error caught: %v\n", r)
			cc = nil
		}
	}()
	return Decr(key)
}

// clear all cache.
func ClearAll(cc Cache, ) error {
	if cc == nil {
		return errors.New("cc is nil")
	}
	defer func() {
		if r := recover(); r != nil {
			//fmt.Println("get cache error caught: %v\n", r)
			cc = nil
		}
	}()
	return ClearAll()
}

// 用json进行数据编码
//
func EncodeJson(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

// -------------------
// Decode
// 用json进行数据解码
//
func DecodeJson(data []byte, to interface{}) error {
	return json.Unmarshal(data, to)
}

//将一个数据结构转填充另一个数据结构
func ChangeStructByEncodeJson(from interface{}, to interface{}) error {
	data, err := EncodeJson(from)
	if err != nil{
		return err
	}
	return DecodeJson(data, to)
}
