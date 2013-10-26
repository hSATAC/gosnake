package main

import (
	//"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetSize(t *testing.T) {
	var scene = Scene{}
	scene.SetSize(100, 100)
	assert.Equal(t, scene.size.width, 50)
	assert.Equal(t, scene.size.height, 100)
}

func TestSetSizeOddNumber(t *testing.T) {
	var scene = Scene{}
	scene.SetSize(99, 99)
	assert.Equal(t, scene.size.width, 49)
	assert.Equal(t, scene.size.height, 99)
}
