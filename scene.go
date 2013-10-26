package main

import (
	"github.com/nsf/termbox-go"
)

type Character interface {
	Move()
	MoveInScreenSize(screenSize ScreenSize)
	Turn(direction Direction)
	Draw()
}

type ScreenSize struct {
	width  int
	height int
}

type Scene struct {
	character Character
	size      ScreenSize
}

func (scene *Scene) SetSize(width int, height int) {
	width = width / 2
	scene.size = ScreenSize{width: width, height: height}
}

func (scene *Scene) Draw() {
	termbox.Flush()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	scene.character.Draw()
	//scene.character.Move()
	scene.character.MoveInScreenSize(scene.size)
}
