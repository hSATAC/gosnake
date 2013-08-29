package main

const (
	SNAKE_DIRECTION_UP = iota
	SNAKE_DIRECTION_RIGHT
	SNAKE_DIRECTION_LEFT
	SNAKE_DIRECTION_DOWN
)

type Snake struct {
	Direction int
	Body      [][]int
}

func NewSnake() *Snake {
	snake := Snake{Direction: SNAKE_DIRECTION_RIGHT}
	return &snake
}
