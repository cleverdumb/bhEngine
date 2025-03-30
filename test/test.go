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

	e.Spawn(E.Entity{X: 100, Y: 100, Shape: E.Circle, D: map[E.EntityProp]interface{}{E.R: float32(100)}})
	e.Spawn(E.Entity{X: 200, Y: 300, Shape: E.Circle, D: map[E.EntityProp]interface{}{E.R: float32(20)}})

	endChan := make(chan int)
	e.BeginGameLoop(endChan)

	// fmt.Println(e)
}
