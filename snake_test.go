package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSnake(t *testing.T) {
	var snake = NewSnake()
	assert.Equal(t, snake.Direction, SNAKE_DIRECTION_RIGHT, "")
}
