package main

import (
	"fmt"

	"example.com/engine"
)

func main() {
	fmt.Println("Hello world")
	e := engine.InitEngine()

	e.Configure(engine.InputMode, engine.FOLLOWMOUSE)

	endChan := make(chan int)
	e.BeginGameLoop(endChan)

	// fmt.Println(e)
}
