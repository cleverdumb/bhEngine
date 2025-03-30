package engine

import (
	"math"
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	windowWidth  = 800
	windowHeight = 800
)

var window *glfw.Window

func toScr(x, y float32) (float32, float32) {
	return (x/windowWidth)*2 - 1, (windowHeight-y)/windowHeight*2 - 1
}

// var game *Game

// func drawCircle(x, y, radius float32, segments int) {
// 	gl.Begin(gl.LINE_LOOP)
// 	for i := 0; i < segments; i++ {
// 		angle := 2 * math.Pi * float64(i) / float64(segments)
// 		xPos := x + radius*float32(math.Cos(angle))
// 		yPos := y + radius*float32(math.Sin(angle))
// 		gl.Vertex2f(xPos, yPos)
// 	}
// 	gl.End()
// }

func fillCircle(x, y, radius float32, segments int) {
	gl.Begin(gl.TRIANGLE_FAN)
	gl.Vertex2f(x, y) // Center of the circle
	for i := 0; i <= segments; i++ {
		angle := 2 * math.Pi * float64(i) / float64(segments)
		xPos := x + radius*float32(math.Cos(angle))
		yPos := y + radius*float32(math.Sin(angle))
		gl.Vertex2f(xPos, yPos)
	}
	gl.End()
}

// import "fmt"

func InitRender() {
	// game = g
	runtime.LockOSThread()

	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)

	var err error
	window, err = glfw.CreateWindow(windowWidth, windowHeight, "GLFW Circle", nil, nil)
	// window = w
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}

	gl.Ortho(-1, 1, -1, 1, -1, 1) // Set orthographic projection

	// for !window.ShouldClose() {
	// 	gl.Clear(gl.COLOR_BUFFER_BIT)

	// 	gl.Color3f(1.0, 0.0, 0.0) // Red circle
	// 	fillCircle(0, 0, 0.5, 100)

	// 	window.SwapBuffers()
	// 	glfw.PollEvents()
	// }
}

func RenderFrame(g *Game) {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.Color3f(1.0, 0.0, 0.0)
	px, py := toScr(g.Pl.X, g.Pl.Y)
	fillCircle(px, py, 0.02, 100)

	for _, e := range g.Entities {
		renderEntity(&e)
	}
	// fmt.Println(g.Pl)

	window.SwapBuffers()
	glfw.PollEvents()
}

func renderEntity(e *Entity) {
	switch e.Shape {
	case Circle:
		gl.Color3f(0.0, 0.0, 1.0)
		ex, ey := toScr(e.X, e.Y)
		fillCircle(ex, ey, 2*(num(e.D[R]))/windowWidth, 100)
	}
}
