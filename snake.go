package main

const (
	SNAKE_DIRECTION_UP = iota
	SNAKE_DIRECTION_RIGHT
	SNAKE_DIRECTION_LEFT
	SNAKE_DIRECTION_DOWN
)

type Snake struct {
	Direction int
	Body      []Node
}

type Node struct {
	x int
	y int
}

func NewSnake() *Snake {
	snake := Snake{}

	// give default
	snake.Direction = SNAKE_DIRECTION_RIGHT
	snake.Body = append(snake.Body, Node{x: 0, y: 0})
	return &snake
}

func (snake *Snake) Len() int {
	return len(snake.Body)
}
