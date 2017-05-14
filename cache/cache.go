package cache

import (
	"io/ioutil"
	"log"
	"time"

	rediscache "github.com/go-redis/cache"
	"github.com/go-redis/redis"
	msgpack "gopkg.in/vmihailenco/msgpack.v2"
)

func init() {
	log.SetOutput(ioutil.Discard)
}

// Cache interface
type Cache interface {
	Close() error
	Get(key string, result interface{}) (err error)
	Put(key string, obj interface{}) (err error)
}

// Redis cache main type
type Redis struct {
	ring  *redis.Ring
	codec *rediscache.Codec
}

// New redis cache
func New(url string) Cache {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server": url,
		},
	})
	codec := &rediscache.Codec{
		Redis: ring,
		Marshal: func(v interface{}) ([]byte, error) {
			return msgpack.Marshal(v)
		},
		Unmarshal: func(b []byte, v interface{}) error {
			return msgpack.Unmarshal(b, v)
		},
	}
	return &Redis{
		ring:  ring,
		codec: codec,
	}
}

// Close the redis connections
func (c *Redis) Close() error {
	return c.ring.Close()
}

// Get a key from cache
func (c *Redis) Get(key string, result interface{}) (err error) {
	return c.codec.Get(key, result)
}

// Put cache something
func (c *Redis) Put(key string, obj interface{}) (err error) {
	return c.codec.Set(&rediscache.Item{
		Key:        key,
		Object:     obj,
		Expiration: time.Hour * 24 * 30, // 1mo
	})
}
