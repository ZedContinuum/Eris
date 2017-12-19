package main

import (
	"fmt"
	"github.com/markedhero/flagit"
	"os"
	"zedcontinuum.com/Zed/Eris"
)

func main() {
	//Print the version of the engine, just cause
	fmt.Println(Eris.Version())

	//First, the engine must be started.
	/* This creates the Master Engine Object that controls the game
	 * OpenGL goes through its init process
	 * All libraries are loaded
	 * The configuration file is loaded
	 * The engine itself is initialized by creating the 3 base Engine Components
	 */
	master, err := Eris.Awaken()
	if err != nil {
		fmt.Println("A Fatal Error has occurred: ", err)
		return
	}

	//This is just to collect what flags are passed into the application.
	flags, err := parseFlags()
	if err != nil {
		fmt.Println("A Fatal Error has occurred: ", err)
		return
	}
	if flags.fullscreen {
		fmt.Println("You set the Fullscreen flag!")
	}

	//Set the initial Scene
	master.SceneManager.FirstScene(&myScene{})

	//This starts the engine's Main Loop
	err = master.Start()
	if err != nil {
		fmt.Println(err)
		return
	}

}

type flagData struct {
	fullscreen bool
}

func parseFlags() (*flagData, error) {
	flags := flagit.NewFlag()
	flagdata := &flagData{}

	flags.Bool(&flagdata.fullscreen, []string{"-f", "--fullscreen"}, "Sets the game fullscreen, overrides config file")

	_, err := flags.Parse(os.Args[1:])
	if err == flagit.ErrNoFlags {
		err = nil
	}
	return flagdata, err
}
