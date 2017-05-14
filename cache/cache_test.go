package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCache(t *testing.T) {
	var assert = assert.New(t)
	var cache = New(":6379")
	var key = "key"
	var value = "value"
	var result string
	defer func() {
		assert.NoError(cache.Close())
	}()
	cache.codec.Delete(key)
	assert.Error(cache.Get(key, &result))
	assert.NoError(cache.Put(key, value))
	assert.NoError(cache.Get(key, &result))
	assert.Equal(value, result)
}
