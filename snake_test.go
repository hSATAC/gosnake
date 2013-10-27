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

func TestBodyContains(t *testing.T) {
	var body = Body{Node{x: 0, y: 0}, Node{x: 1, y: 1}}
	assert.Equal(t, body.Contains(Node{x: 1, y: 1}), true)
}

func TestBodyContainsFails(t *testing.T) {
	var body = Body{Node{x: 0, y: 0}, Node{x: 1, y: 1}}
	assert.Equal(t, body.Contains(Node{x: 1, y: 2}), false)
}
