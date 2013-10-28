package main

import (
	//"github.com/davecgh/go-spew/spew"
	"math"
)

const (
	SNAKE_DIRECTION_UP Direction = iota
	SNAKE_DIRECTION_RIGHT
	SNAKE_DIRECTION_DOWN
	SNAKE_DIRECTION_LEFT
)

type Direction byte

func (direction Direction) angle() int {
	return int(direction) * 90
}

type Snake struct {
	direction Direction
	body      Body
}

type Body []Node

type Node struct {
	x int
	y int
}

func (body Body) Contains(node Node) bool {
	for _, n := range body {
		if node == n {
			return true
		}
	}
	return false
}

func (snake *Snake) Move(screenSize ScreenSize, fruit Node) CharacterStatus {
	newHead := snake.NewHead(screenSize)
	if snake.body.Contains(newHead) {
		return CHARACTER_STATUS_DEAD
	} else if newHead == fruit {
		snake.GrowInScreenSize(screenSize)
		return CHARACTER_STATUS_GROW
	} else {
		snake.MoveInScreenSize(screenSize)
		return CHARACTER_STATUS_MOVE
	}
}

func (snake *Snake) MoveInScreenSize(screenSize ScreenSize) {
	snake.body = snake.body[1:]
	head := snake.NewHead(screenSize)
	snake.body = append(snake.body, head)
}

func (snake *Snake) GrowInScreenSize(screenSize ScreenSize) {
	head := snake.NewHead(screenSize)
	snake.body = append(snake.body, head)
}

func (snake *Snake) NewHead(screenSize ScreenSize) Node {
	head := snake.head()
	var newHead Node

	switch snake.direction {
	case SNAKE_DIRECTION_RIGHT:
		newHead = Node{x: head.x + 1, y: head.y}
	case SNAKE_DIRECTION_DOWN:
		newHead = Node{x: head.x, y: head.y + 1}
	case SNAKE_DIRECTION_LEFT:
		newHead = Node{x: head.x - 1, y: head.y}
	case SNAKE_DIRECTION_UP:
		newHead = Node{x: head.x, y: head.y - 1}
	}

	if screenSize.width > 0 {
		if newHead.x < 0 { // over left edge
			newHead.x = screenSize.width - 1
		} else if newHead.x >= screenSize.width { // over right edge
			newHead.x = screenSize.width - newHead.x
		}
	}

	if screenSize.height > 0 {
		if newHead.y < 0 { // over top edge
			newHead.y = screenSize.height - 1
		} else if newHead.y >= screenSize.height { // over bottom edge
			newHead.y = screenSize.height - newHead.y
		}
	}
	return newHead
}

func (snake *Snake) Len() int {
	return len(snake.body)
}

func (snake *Snake) head() Node {
	return snake.body[snake.Len()-1]
}

func (snake *Snake) Turn(direction Direction) {
	// You can't turn to opposite direction
	angle := float64(direction.angle() - snake.direction.angle())
	if math.Abs(angle) == 180.0 {
		return
	}

	snake.direction = direction
}

func (snake *Snake) Draw() {
	for _, node := range snake.body {
		DrawPoint(node.x, node.y, COLOR_CHARACTER)
	}
}

func (snake *Snake) Body() Body {
	return snake.body
}

// TODO: init with position, direction & length
func NewSnake() *Snake {
	snake := Snake{}

	// give default
	snake.direction = SNAKE_DIRECTION_RIGHT
	snake.body = Body{Node{x: 0, y: 0}, Node{x: 1, y: 0}}
	return &snake
}
