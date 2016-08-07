package main

import (
	"log"
	"runtime"
	"fmt"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)


func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

const (
	WIDTH int = 600
	HALF float32 = 300
)

func handleClick(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
	if action != glfw.Press {
		return
	}

	xpos, ypos := w.GetCursorPos()

	if button == glfw.MouseButton2 {
		//right
		fmt.Printf("RIGHT %f %f\n", xpos, ypos)

	}

	if button == glfw.MouseButton1 {
		//left
		fmt.Printf("LEFT %f %f\n", xpos, ypos)
	}

}

func setup() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	window, err := glfw.CreateWindow(WIDTH, WIDTH, "Cube", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}


	setupScene()

	window.SetMouseButtonCallback(handleClick);

	for !window.ShouldClose() {
		drawScene()
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func setupScene() {
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
}

func destroyScene() {
}

func screenCoordToGrid(x int, y int) (x1 int, y1 int) {
	return x / GRIDLENGTH, y / GRIDLENGTH
}

func drawSquare(x int, y int, w int) {
	var x1, x2, y1, y2 float32

	x1 = (float32(x) - HALF) / HALF
	y1 = (HALF - float32(y)) / HALF
	x2 = (float32(x + w) - HALF) / HALF
	y2 = (HALF - float32(y + w)) / HALF

	gl.Color4f(1, 0, 0, 1)

	gl.Begin(gl.QUADS)

	gl.Vertex3f(x1, y1, 1)
	gl.Vertex3f(x2, y1, 1)
	gl.Vertex3f(x2, y2, 1)
	gl.Vertex3f(x1, y2, 1)

	gl.End()
}

func drawScene() {
	gl.Clear(gl.COLOR_BUFFER_BIT)

	gl.LoadIdentity()

	for i := 0; i < 12; i++ {
		drawSquare(i * 50, i*50, 50)
	}
	
}