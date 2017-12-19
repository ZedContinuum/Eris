package Eris

//The interface that defines
type Engineer interface {
	Init(m *Eris) error
	Update()
	Draw()
	CleanUp()
	Reset()
	OnVisible()
}

type EngineComponent struct {
	Engineer
	master *Eris
}

func (c *EngineComponent) Init(m *Eris) error {
	c.master = m
	return nil
}

func (c *EngineComponent) Update() {

}

func (c EngineComponent) Draw() {

}

func (c *EngineComponent) CleanUp() {

}

func (c *EngineComponent) Reset() {

}

func (c *EngineComponent) OnVisible() {

}
