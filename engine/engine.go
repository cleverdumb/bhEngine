package engine

import (
	"math"

	"github.com/go-gl/glfw/v3.3/glfw"
)

type Game struct {
	Pl       Player
	Config   map[ConfigureType]interface{}
	Entities []Entity
}

type Player struct {
	X float32
	Y float32
	R float32
}

// var pl Player
var keyDown = make(map[glfw.Key]bool)
var mousePos = [2]float32{0, 0}

// var game Game

func InitEngine() Game {
	InitRender()

	mx, my := window.GetCursorPos()
	mousePos = [2]float32{float32(mx), float32(my)}

	g := Game{}
	g.Config = getDefaultGameConfig()
	g.Entities = make([]Entity, 0)
	pl := Player{400, 400, 10}
	g.Pl = pl

	window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		if action == glfw.Press {
			keyDown[key] = true
		} else if action == glfw.Release {
			keyDown[key] = false
		}
	})
	window.SetCursorPosCallback(func(w *glfw.Window, x float64, y float64) {
		mousePos[0] = float32(x)
		mousePos[1] = float32(y)
	})
	return g
}

func (e *Game) BeginGameLoop(c chan int) {
out:
	for !window.ShouldClose() {
		select {
		case <-c:
			defer glfw.Terminate()
			break out
		default:
			inputProcess(e)
			RenderFrame(e)

			for _, entity := range e.Entities {
				entity.CheckCB(e)
			}
			// time.Sleep(1 * time.Second)
		}
	}
}

// const spd = 7
// const followEpsilon = 10

func inputProcess(e *Game) {
	spd := num(e.Config[PlayerSpd])
	switch e.Config[InputMode].(ConfigureVal) {
	case WASD:
		if keyDown[glfw.KeyW] {
			if keyDown[glfw.KeyA] || keyDown[glfw.KeyD] {
				e.Pl.Y -= spd / math.Sqrt2
			} else {
				e.Pl.Y -= spd
			}
		}
		if keyDown[glfw.KeyA] {
			if keyDown[glfw.KeyW] || keyDown[glfw.KeyS] {
				e.Pl.X -= spd / math.Sqrt2
			} else {
				e.Pl.X -= spd
			}
		}
		if keyDown[glfw.KeyS] {
			if keyDown[glfw.KeyA] || keyDown[glfw.KeyD] {
				e.Pl.Y += spd / math.Sqrt2
			} else {
				e.Pl.Y += spd
			}
		}
		if keyDown[glfw.KeyD] {
			if keyDown[glfw.KeyW] || keyDown[glfw.KeyS] {
				e.Pl.X += spd / math.Sqrt2
			} else {
				e.Pl.X += spd
			}
		}
	case FollowMouse:
		dy := mousePos[1] - e.Pl.Y
		dx := mousePos[0] - e.Pl.X

		l := math.Hypot(float64(dx), float64(dy))

		if l < float64(spd) {
			e.Pl.X = mousePos[0]
			e.Pl.Y = mousePos[1]
		} else {
			dy /= float32(l)
			dx /= float32(l)

			e.Pl.X += dx * spd
			e.Pl.Y += dy * spd
		}
	case Snap:
		e.Pl.X = mousePos[0]
		e.Pl.Y = mousePos[1]
	}
	if e.Config[ScreenBound].(ConfigureVal)&BoundX > 0 {
		e.Pl.X = max(e.Pl.X, e.Pl.R)
		e.Pl.X = min(windowWidth-e.Pl.R, e.Pl.X)
	}

	if e.Config[ScreenBound].(ConfigureVal)&BoundY > 0 {
		e.Pl.Y = max(e.Pl.Y, e.Pl.R)
		e.Pl.Y = min(windowWidth-e.Pl.R, e.Pl.Y)
	}
}

func (e *Game) Configure(t ConfigureType, v interface{}) {
	// switch t {
	// case InputMode:
	// 	var inputFunc glfw.CharCallback
	// 	if v == WASD {
	// 		inputFunc = func(w *glfw.Window, c rune) {

	// 		}
	// 	}
	// 	window.SetCharCallback(inputFunc)
	// }
	e.Config[t] = v
}

func (g *Game) Spawn(e Entity) {
	g.Entities = append(g.Entities, e)
}

func (g *Game) NewEntity(x, y float32, shape Shape, data PropMap) Entity {
	return Entity{
		X:     x,
		Y:     y,
		Shape: shape,
		D:     data,
		CB:    make(CBMap),
	}
}
