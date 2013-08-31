package main

import (
	"github.com/nsf/termbox-go"
)

type ScreenSize struct {
	width  int
	height int
}

type Scene struct {
	snake Snake
	size  ScreenSize
}

func (scene *Scene) SetSize(width int, height int) {
	scene.size = ScreenSize{width: width, height: height}
}

func (scene *Scene) Draw() {
	termbox.Flush()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for _, node := range scene.snake.body {
		termbox.SetCell(node.x, node.y, ' ', termbox.ColorDefault, termbox.ColorRed)
	}
	scene.snake.Move()
}
