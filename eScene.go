package Eris

import (
	"fmt"
	"github.com/ZedContinuum/mobile/gl"
)

type eScene struct {
	EngineComponent
	activeScene scenery
	nextScene   scenery
}

func (s *eScene) Init(m *Eris) error {
	s.master = m
	if s.nextScene == nil {
		return fmt.Errorf("No Initial Scene set, please initialize your first scene and set via master.SceneManager.FirstScene()\n")
	} else {
		s.LoadScene(s.nextScene)
	}
	return nil
}

func (s *eScene) FirstScene(first scenery) {
	s.nextScene = first
}

func (s *eScene) LoadScene(nex scenery) {
	s.nextScene = nex
	s.nextScene.init(s)
}

func (s *eScene) NextScene() {
	s.activeScene = s.nextScene
	s.nextScene = nil
	s.activeScene.Start()
}

func (s *eScene) Update() {
	s.activeScene.Update(s.master.deltaTime)
}

func (s *eScene) Draw() {
	s.activeScene.draw(s.master.glctx)
}

func (s *eScene) OnVisible() {

}

func (s eScene) GetMaster() *Eris {
	return s.master
}

type scenery interface {
	init(m *eScene)
	Start()
	Update(float32)
	draw(gl.Context)
	CleanUp()
	OnPause()
	OnResume()
}

type Scene struct {
	scenery
	Manager    *eScene
	Background Color
	Name       string
}

func (s *Scene) init(m *eScene) {
	s.Manager = m
}

func (s Scene) draw(glctx gl.Context) {
	glctx.ClearColor(s.Background.R, s.Background.G, s.Background.B, s.Background.A)
	glctx.Clear(gl.COLOR_BUFFER_BIT)

}
