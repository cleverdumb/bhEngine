package main

import (
	"fmt"

	"example.com/engine"
)

func main() {
	fmt.Println("Hello world")
	e := engine.InitEngine()

	e.Configure(engine.InputMode, engine.Snap)
	e.Configure(engine.PlayerSpd, float32(7))

	endChan := make(chan int)
	e.BeginGameLoop(endChan)

	// fmt.Println(e)
}
