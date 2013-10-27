package main

import (
	//"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScene(t *testing.T) {
}

func TestSceneAvailableNodes(t *testing.T) {
	t.SkipNow()
	var snake = NewSnake()
	var scene = Scene{size: ScreenSize{5, 5}, character: snake}
	assert.Equal(t, scene.AvailableNodes(), []Node{Node{x: 0, y: 0}})
}
