package main

import (
	//"github.com/nsf/termbox-go"
	"math/rand"
	"time"
)

type Character interface {
	MoveInScreenSize(screenSize ScreenSize)
	Turn(direction Direction)
	Draw()
	Body() Body
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
	scene.size = ScreenSize{width: width, height: height}
}

func (scene *Scene) Draw() {
	ClearScene()
	scene.character.Draw()
	scene.character.MoveInScreenSize(scene.size)
}

func (scene *Scene) availableNodes() (availableNodes []Node) {
	for x := 0; x < scene.size.width; x++ {
		for y := 0; y < scene.size.height; y++ {
			node := Node{x: x, y: y}
			if !scene.character.Body().Contains(node) {
				availableNodes = append(availableNodes, node)
			}
		}
	}

	return availableNodes
}

func (scene *Scene) randomAvailableNode() Node {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	nodes := scene.availableNodes()
	return nodes[r.Intn(len(nodes))]
}
