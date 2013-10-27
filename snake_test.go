package main

import (
	//"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSnake(t *testing.T) {
	var snake = NewSnake()
	assert.Equal(t, snake.direction, Direction(SNAKE_DIRECTION_RIGHT))
	assert.Equal(t, len(snake.body), 2)
}

func TestSnakeLen(t *testing.T) {
	var snake = NewSnake()
	snake.body = Body{Node{x: 0, y: 0}, Node{x: 1, y: 1}}
	assert.Equal(t, snake.Len(), 2)
}

func TestSnakeMove(t *testing.T) {
	var snake = NewSnake()
	snake.Move()
	assert.Equal(t, snake.Len(), 2)
	assert.Equal(t, snake.body[0], Node{x: 1, y: 0})
	assert.Equal(t, snake.body[1], Node{x: 2, y: 0})
}

func TestSnakeMoveDown(t *testing.T) {
	var snake = NewSnake()
	snake.direction = SNAKE_DIRECTION_DOWN
	snake.Move()
	assert.Equal(t, snake.Len(), 2)
	assert.Equal(t, snake.body[0], Node{x: 1, y: 0})
	assert.Equal(t, snake.body[1], Node{x: 1, y: 1})
}

func TestSnakeMoveLeft(t *testing.T) {
	var snake = NewSnake()
	snake.direction = SNAKE_DIRECTION_LEFT
	snake.body = Body{Node{x: 5, y: 5}, Node{x: 4, y: 5}}
	snake.Move()
	assert.Equal(t, snake.Len(), 2)
	assert.Equal(t, snake.body[0], Node{x: 4, y: 5})
	assert.Equal(t, snake.body[1], Node{x: 3, y: 5})
}

func TestSnakeMoveUp(t *testing.T) {
	var snake = NewSnake()
	snake.direction = SNAKE_DIRECTION_UP
	snake.body = Body{Node{x: 5, y: 5}, Node{x: 4, y: 5}}
	snake.Move()
	assert.Equal(t, snake.Len(), 2)
	assert.Equal(t, snake.body[0], Node{x: 4, y: 5})
	assert.Equal(t, snake.body[1], Node{x: 4, y: 4})
}

func TestSnakeMoveOutOfRightEdge(t *testing.T) {
	var snake = NewSnake()
	var screenSize = ScreenSize{3, 3}
	snake.MoveInScreenSize(screenSize)
	assert.Equal(t, snake.Len(), 2)
	assert.Equal(t, snake.body[0], Node{x: 1, y: 0})
	assert.Equal(t, snake.body[1], Node{x: 2, y: 0})

	snake.MoveInScreenSize(screenSize)
	assert.Equal(t, snake.Len(), 2)
	assert.Equal(t, snake.body[0], Node{x: 2, y: 0})
	assert.Equal(t, snake.body[1], Node{x: 0, y: 0})
}

func TestSnakeMoveOutOfTopEdge(t *testing.T) {
	var snake = NewSnake()
	var screenSize = ScreenSize{3, 3}
	snake.Turn(SNAKE_DIRECTION_UP)
	snake.MoveInScreenSize(screenSize)
	assert.Equal(t, snake.Len(), 2)
	assert.Equal(t, snake.body[0], Node{x: 1, y: 0})
	assert.Equal(t, snake.body[1], Node{x: 1, y: 2})

	snake.MoveInScreenSize(screenSize)
	assert.Equal(t, snake.Len(), 2)
	assert.Equal(t, snake.body[0], Node{x: 1, y: 2})
	assert.Equal(t, snake.body[1], Node{x: 1, y: 1})
}

func TestSnakeMoveOutOfBottomEdge(t *testing.T) {
	var snake = NewSnake()
	var screenSize = ScreenSize{3, 3}

	snake.Turn(SNAKE_DIRECTION_DOWN)
	snake.MoveInScreenSize(screenSize)

	snake.MoveInScreenSize(screenSize)
	assert.Equal(t, snake.Len(), 2)
	assert.Equal(t, snake.body[0], Node{x: 1, y: 1})
	assert.Equal(t, snake.body[1], Node{x: 1, y: 2})

	snake.MoveInScreenSize(screenSize)
	assert.Equal(t, snake.Len(), 2)
	assert.Equal(t, snake.body[0], Node{x: 1, y: 2})
	assert.Equal(t, snake.body[1], Node{x: 1, y: 0})
}

func TestSnakeMoveOutOfLeftEdge(t *testing.T) {
	var snake = NewSnake()
	var screenSize = ScreenSize{3, 3}

	snake.Turn(SNAKE_DIRECTION_DOWN)
	snake.MoveInScreenSize(screenSize)

	snake.Turn(SNAKE_DIRECTION_LEFT)
	snake.MoveInScreenSize(screenSize)
	assert.Equal(t, snake.Len(), 2)
	assert.Equal(t, snake.body[0], Node{x: 1, y: 1})
	assert.Equal(t, snake.body[1], Node{x: 0, y: 1})

	snake.MoveInScreenSize(screenSize)
	assert.Equal(t, snake.Len(), 2)
	assert.Equal(t, snake.body[0], Node{x: 0, y: 1})
	assert.Equal(t, snake.body[1], Node{x: 2, y: 1})
}

func TestSnakeNewHead(t *testing.T) {
	var snake = NewSnake()
	var screenSize = ScreenSize{3, 3}

	assert.Equal(t, snake.NewHead(screenSize), Node{x: 2, y: 0})

	snake.MoveInScreenSize(screenSize)
	assert.Equal(t, snake.NewHead(screenSize), Node{x: 0, y: 0})
}

func TestSnakeGrow(t *testing.T) {
	var snake = NewSnake()
	var screenSize = ScreenSize{3, 3}

	snake.GrowInScreenSize(screenSize)

	assert.Equal(t, snake.Len(), 3)
	assert.Equal(t, snake.body[0], Node{x: 0, y: 0})
	assert.Equal(t, snake.body[1], Node{x: 1, y: 0})
	assert.Equal(t, snake.body[2], Node{x: 2, y: 0})

	snake.Turn(SNAKE_DIRECTION_DOWN)
	snake.MoveInScreenSize(screenSize)
	assert.Equal(t, snake.Len(), 3)
	assert.Equal(t, snake.body[0], Node{x: 1, y: 0})
	assert.Equal(t, snake.body[1], Node{x: 2, y: 0})
	assert.Equal(t, snake.body[2], Node{x: 2, y: 1})
}

func TestSnakeTurn(t *testing.T) {
	var snake = NewSnake()
	snake.Turn(SNAKE_DIRECTION_UP)
	assert.Equal(t, snake.direction, Direction(SNAKE_DIRECTION_UP))
}

func TestSnakeUTurn(t *testing.T) {
	var snake = NewSnake()
	snake.direction = SNAKE_DIRECTION_DOWN
	snake.Turn(SNAKE_DIRECTION_UP)
	assert.Equal(t, snake.direction, Direction(SNAKE_DIRECTION_DOWN))
}
