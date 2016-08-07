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
	window, err := glfw.CreateWindow(600, 600, "Cube", nil, nil)
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

func drawScene() {
	gl.Clear(gl.COLOR_BUFFER_BIT)

	gl.LoadIdentity()

	gl.Color4f(1, 0, 0, 1)

	gl.Begin(gl.QUADS)

	gl.Vertex3f(0, 600, 1)
	gl.Vertex3f(600, 600, 1)
	gl.Vertex3f(600, 0, 1)
	gl.Vertex3f(0, 0, 1)

	gl.End()

	gl.Color4f(0, 1, 0, 1)

	gl.Begin(gl.QUADS)

	gl.Vertex3f(0, 300, 1)
	gl.Vertex3f(300, 300, 1)
	gl.Vertex3f(300, 0, 1)
	gl.Vertex3f(0, 0, 1)

	gl.End()

	gl.Color4f(0, 0, 1, 1)

	gl.Begin(gl.QUADS)

	gl.Vertex3f(600, 300, 1)
	gl.Vertex3f(300, 300, 1)
	gl.Vertex3f(300, 600, 1)
	gl.Vertex3f(600, 600, 1)

	gl.End()
}