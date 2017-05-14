package cache

import (
	"time"

	msgpack "gopkg.in/vmihailenco/msgpack.v2"

	rediscache "github.com/go-redis/cache"
	"github.com/go-redis/redis"
)

// Redis cache main type
type Redis struct {
	ring  *redis.Ring
	codec *rediscache.Codec
}

// New redis cache
func New(url string) *Redis {
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
