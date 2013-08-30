package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSnake(t *testing.T) {
	var snake = NewSnake()
	assert.Equal(t, snake.Direction, SNAKE_DIRECTION_RIGHT, "")
	assert.Equal(t, len(snake.Body), 1, "")
}

func TestSnakeLen(t *testing.T) {
	var snake = NewSnake()
	snake.Body = []Node{Node{x: 0, y: 0}, Node{x: 1, y: 1}}
	assert.Equal(t, snake.Len(), 2, "")
}
