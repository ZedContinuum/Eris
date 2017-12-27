package Eris

import (
	"fmt"
	"github.com/ZedContinuum/Eris/Events"
	"github.com/ZedContinuum/mobile/app"
	"github.com/ZedContinuum/mobile/event/key"
	"github.com/ZedContinuum/mobile/event/lifecycle"
	"github.com/ZedContinuum/mobile/event/paint"
	"github.com/ZedContinuum/mobile/event/touch"
	"github.com/ZedContinuum/mobile/gl"
	"os"
)

//The Constants declare the name and the current Version of the Engine
const (
	NAME    = "Eris Game Framework"
	VERSION = "0.10"
	TAG     = "-alpha"
)

//Returns a string of the current version of this library
func Version() string {
	return fmt.Sprintf("%s Game Engine v%s%s", NAME, VERSION, TAG)
}

//The Master Data type contains all of the top-level data organized
type Eris struct {
	Title string

	//Private Variables
	isRunning bool
	isVisible bool
	glctx     gl.Context
	app       app.App

	deltaTime    float32
	SceneManager eScene

	components []Engineer
}

//Just a fun name for the "Create" function
//Returns a pointer reference to The Master type
func Awaken() (*Eris, error) {
	tm := &Eris{}
	tm.components = make([]Engineer, 0)
	tm.components = append(tm.components, new(eTime))
	tm.SceneManager = eScene{}
	return tm, nil
}

//Starts the Engine
//Your application no longer has control
func (tm *Eris) Start() error {
	tm.isRunning = true
	var err error
	for i := range tm.components {
		tm.components[i].Init(tm)
	}
	err = tm.SceneManager.Init(tm)
	if err != nil {
		return startError(err)
	}
	app.Main(tm.master)
	os.Exit(1)
	return nil
}

//This is the main function
func (tm *Eris) master(a app.App) {
	tm.app = a
	for e := range a.Events() {
		switch e := a.Filter(e).(type) {
		case lifecycle.Event:
			switch e.Crosses(lifecycle.StageVisible) {
			case lifecycle.CrossOn:
				//This is to check if we have already initialized the game or not.  isRunning is only set once
				// If isRunning is set to true, then we will need to call the OnVisible function on components to ensure stability
				if tm.isRunning {
					for i := range tm.components {
						tm.components[i].OnVisible()
					}
				}
				tm.glctx, _ = e.DrawContext.(gl.Context)
				tm.isVisible = true
				tm.SceneManager.NextScene()
				a.Send(paint.Event{})
			case lifecycle.CrossOff:
				tm.isVisible = false
			}
			if e.Crosses(lifecycle.StageAlive) == lifecycle.CrossOff {
				tm.CleanUp()
				return
			}
		case paint.Event:
			if tm.glctx == nil || e.External {
				continue
			}
			tm.update()
			tm.draw()
			a.Publish()
			if tm.isVisible {
				a.Send(paint.Event{})
			}
		case touch.Event:
			fmt.Printf("Click at  X: %f | Y: %f\n", e.X, e.Y)
		case key.Event:
			fmt.Println("Pressed key: ", e.Code)
			if e.Code == key.CodeSpacebar {
				tm.Quit()
			}
		case Events.QuitEvent:
			quitEvent := lifecycle.Event{}
			quitEvent.From = lifecycle.StageAlive
			quitEvent.To = lifecycle.StageAlive - 1
			tm.app.Send(quitEvent)
		}
		if !tm.isRunning {
			break
		}
	}
}

//Main Update Loop
//Meant to update all scripts and objects
func (tm *Eris) update() {
	for i := range tm.components {
		tm.components[i].Update()
	}
	tm.SceneManager.Update()
}

//Main Draw Loop
//Meant for Drawing objects
func (tm *Eris) draw() {
	for i := range tm.components {
		tm.components[i].Draw()
	}
	tm.SceneManager.Draw()
}

func (tm *Eris) CleanUp() {
	fmt.Println("Clean up!")
	tm.glctx = nil
	for i := range tm.components {
		tm.components[i].CleanUp()
	}
	tm.SceneManager.CleanUp()
	tm.components = nil
}

//Sends the "Quit Event"
func (tm *Eris) Quit() {
	tm.app.Send(Events.QuitEvent{})
}

func startError(e error) error {
	return fmt.Errorf("Unable to start %s: %s", NAME, e)
}
