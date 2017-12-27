package Eris

import (
	"github.com/ZedContinuum/mobile/gl"
)

type componenter interface {
	Start()
	Update(deltaTime float32)
	Draw(glctx gl.Context)
}

type GameComponent struct {
	componenter
	Name string
}

func (gc *GameComponent) Start() {
}

func (gc *GameComponent) Update(deltaTime float32) {
}

func (gc *GameComponent) Draw(glctx gl.Context) {
}
