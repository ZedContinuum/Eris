package main

import (
	"fmt"
	"github.com/ZedContinuum/Eris"
)

type myScene struct {
	Eris.Scene
	state int
}

func (m *myScene) Start() {
	m.Background.A = 1.0
	m.state = 1
	fmt.Println("Starting!")
}

func (m *myScene) Update(delta float32) {
	switch m.state {
	case 0:
		m.Background.R += delta
		if m.Background.R >= 1.0 {
			m.Background.R = 1.0
			m.state = 1
		}
	case 1:
		m.Background.G += delta
		if m.Background.G >= 1.0 {
			m.Background.G = 1.0
			m.state = 2
		}
	case 2:
		m.Background.B += delta
		if m.Background.B >= 1.0 {
			m.Background.B = 1.0
			m.state = 3
		}
	case 3:
		m.Background.R -= delta
		if m.Background.R <= 0.0 {
			m.Background.R = 0.0
			m.state = 4
		}
	case 4:
		m.Background.G -= delta
		if m.Background.G <= 0.0 {
			m.Background.G = 0.0
			m.state = 5
		}
	case 5:
		m.Background.B -= delta
		if m.Background.B <= 0.0 {
			m.Background.B = 0.0
			m.state = 0
		}
	}
}

func (m *myScene) CleanUp() {

}
