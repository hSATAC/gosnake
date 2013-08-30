package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSnake(t *testing.T) {
	var snake = NewSnake()
	assert.Equal(t, snake.Direction, SNAKE_DIRECTION_RIGHT, "")
	assert.Equal(t, len(snake.Body), 2, "")
	spew.Dump("test")
}

func TestSnakeLen(t *testing.T) {
	var snake = NewSnake()
	snake.Body = []Node{Node{x: 0, y: 0}, Node{x: 1, y: 1}}
	assert.Equal(t, snake.Len(), 2, "")
}

func TestSnakeMove(t *testing.T) {
	var snake = NewSnake()
	snake.Move()
	assert.Equal(t, snake.Len(), 2, "")
	assert.Equal(t, snake.Body[0], Node{x: 1, y: 0})
	assert.Equal(t, snake.Body[1], Node{x: 2, y: 0})
}

func TestSnakeMoveDown(t *testing.T) {
	var snake = NewSnake()
	snake.Direction = SNAKE_DIRECTION_DOWN
	snake.Move()
	assert.Equal(t, snake.Len(), 2, "")
	assert.Equal(t, snake.Body[0], Node{x: 1, y: 0})
	assert.Equal(t, snake.Body[1], Node{x: 1, y: 1})
}

func TestSnakeMoveLeft(t *testing.T) {
	var snake = NewSnake()
	snake.Direction = SNAKE_DIRECTION_LEFT
	snake.Body = []Node{Node{x: 5, y: 5}, Node{x: 4, y: 5}}
	snake.Move()
	assert.Equal(t, snake.Len(), 2, "")
	assert.Equal(t, snake.Body[0], Node{x: 4, y: 5})
	assert.Equal(t, snake.Body[1], Node{x: 3, y: 5})
}

func TestSnakeMoveUp(t *testing.T) {
	var snake = NewSnake()
	snake.Direction = SNAKE_DIRECTION_UP
	snake.Body = []Node{Node{x: 5, y: 5}, Node{x: 4, y: 5}}
	snake.Move()
	assert.Equal(t, snake.Len(), 2, "")
	assert.Equal(t, snake.Body[0], Node{x: 4, y: 5})
	assert.Equal(t, snake.Body[1], Node{x: 4, y: 4})
}
