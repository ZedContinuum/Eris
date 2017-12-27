package Eris

type objecter interface {
}

type GameObject struct {
	objecter

	Name     string
	Parent   *GameObject
	Children []*GameObject

	Components []*GameComponent
}
