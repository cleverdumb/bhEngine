package main

import (
	"fmt"

	E "example.com/engine"
)

func main() {
	fmt.Println("Hello world")
	e := E.InitEngine()

	e.Configure(E.InputMode, E.Snap)
	e.Configure(E.PlayerSpd, 10)
	e.Configure(E.ScreenBound, E.BoundX|E.BoundY)

	testEnemy := e.NewEntity(200, 200, E.Circle, E.PropMap{E.R: float32(30)})

	e.Spawn(testEnemy)

	testEnemy.CB[E.CollidePL] = func(pl *E.Player) {
		fmt.Println("Collide")
	}
	// e.Spawn(E.Entity{X: 200, Y: 300, Shape: E.Circle, D: E.PropMap{E.R: float32(20)}})

	endChan := make(chan int)
	e.BeginGameLoop(endChan)

	// fmt.Println(e)
}
