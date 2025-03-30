package main

import (
	"fmt"

	E "example.com/engine"
)

func main() {
	fmt.Println("Hello world")
	e := E.InitEngine()

	e.Configure(E.InputMode, E.Snap)
	e.Configure(E.PlayerSpd, float32(7))

	e.Spawn(E.Entity{100, 100, E.Circle})

	endChan := make(chan int)
	e.BeginGameLoop(endChan)

	// fmt.Println(e)
}
