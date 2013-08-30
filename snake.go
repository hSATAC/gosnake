package main

import (
	"math"
)

const (
	SNAKE_DIRECTION_UP = iota
	SNAKE_DIRECTION_RIGHT
	SNAKE_DIRECTION_DOWN
	SNAKE_DIRECTION_LEFT
)

type Snake struct {
	Direction int
	Body      []Node
}

type Node struct {
	x int
	y int
}

func (snake *Snake) Move() {
	// TODO:if eat fruit, don't kill tail here.
	snake.Body = snake.Body[1:]

	head := snake.head()
	var new_head Node

	// TODO: deal over edge case here.
	switch snake.Direction {
	case SNAKE_DIRECTION_RIGHT:
		new_head = Node{x: head.x + 1, y: head.y}
	case SNAKE_DIRECTION_DOWN:
		new_head = Node{x: head.x, y: head.y + 1}
	case SNAKE_DIRECTION_LEFT:
		new_head = Node{x: head.x - 1, y: head.y}
	case SNAKE_DIRECTION_UP:
		new_head = Node{x: head.x, y: head.y - 1}
	}

	snake.Body = append(snake.Body, new_head)
}

func (snake *Snake) Len() int {
	return len(snake.Body)
}

func (snake *Snake) head() Node {
	return snake.Body[snake.Len()-1]
}

// TODO: see if there's a way to restrict parameter direction to const
func (snake *Snake) Turn(direction int) {
	angle := float64((direction - snake.Direction) * 90)
	if math.Abs(angle) != 180.0 { // You can't turn to opposite direction
		snake.Direction = direction
	}
}

func NewSnake() *Snake {
	snake := Snake{}

	// give default
	snake.Direction = SNAKE_DIRECTION_RIGHT
	snake.Body = []Node{Node{x: 0, y: 0}, Node{x: 1, y: 0}}
	return &snake
}
