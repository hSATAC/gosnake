package main

import (
	//"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScene(t *testing.T) {
	t.SkipNow()
}

func TestSceneAvailableNodes(t *testing.T) {
	var snake = NewSnake()
	var scene = Scene{size: ScreenSize{2, 2}, character: snake}
	assert.Equal(t, scene.availableNodes(), []Node{Node{x: 0, y: 1}, Node{x: 1, y: 1}})
}

func TestSceneRandomAvailableNode(t *testing.T) {
	var snake = NewSnake()
	var scene = Scene{size: ScreenSize{3, 1}, character: snake}
	assert.Equal(t, scene.randomAvailableNode(), Node{x: 2, y: 0})
}
