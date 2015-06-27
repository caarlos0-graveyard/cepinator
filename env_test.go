package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const key = "WHATEVER"

func TestGetenv(t *testing.T) {
	value := "bla"
	os.Setenv(key, value)
	defer os.Unsetenv(key)
	assert.Equal(t, value, Getenv(key, "default"))
}

func TestGetenvUnseted(t *testing.T) {
	value := "default"
	assert.Equal(t, value, Getenv(key, value))
}
