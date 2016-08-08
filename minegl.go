package main

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)


func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

const (
	WIDTH int = 720
	HALF float32 = 360
	SCALE float32 = 45
)

func handleClick(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
	if action != glfw.Press {
		return
	}

	xpos, ypos := w.GetCursorPos()

	if button == glfw.MouseButton2 {
		flagClick(uint(xpos/float64(SCALE)), uint(ypos/float64(SCALE)))
	}

	if button == glfw.MouseButton1 {
		revealClick(uint(xpos/float64(SCALE)), uint(ypos/float64(SCALE)))
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
		glfw.WaitEvents()
	}
}

func setupScene() {
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
}

func destroyScene() {
}

func drawSquare(x float32, y float32, w float32, color int, count int) {

	x1 := (x - HALF) / HALF
	y1 := (HALF - y) / HALF
	x2 := ((x + w) - HALF) / HALF
	y2 := (HALF - (y + w)) / HALF

	countBlockSize := 3.0 / HALF
	countBlockSpace := 4.0 / HALF

	//color 0 = white, 1 = blue, 2 = pink

	switch {
		case color == 0:
			gl.Color4f(1, 1, 1, 1)
		case color == 1:
			if lastState {
				gl.Color4f(0.25, 0.88, 0.82, 1)
			} else {
				gl.Color4f(0, 0, 1, 1)
			}
		case color == 2:
			gl.Color4f(1, 0.7, 0.7, 1)
	}

	gl.Begin(gl.QUADS)

	gl.Vertex3f(x1, y1, 1)
	gl.Vertex3f(x2, y1, 1)
	gl.Vertex3f(x2, y2, 1)
	gl.Vertex3f(x1, y2, 1)

	if color == 0 && count > 0 {
		gl.Color4f(0, 0, 0, 1)

		for i := 0; i < count; i++ {
			fc := float32(i + 2)

			y3 := (y1 + y2) / 2

			gl.Vertex3f(x1 + (fc * countBlockSpace), y3, 1)
			gl.Vertex3f(x1 + (fc * countBlockSpace) + countBlockSize, y3, 1)
			gl.Vertex3f(x1 + (fc * countBlockSpace) + countBlockSize, y3 - countBlockSize, 1)
			gl.Vertex3f(x1 + (fc * countBlockSpace), y3 - countBlockSize, 1)
		} 
	}

	gl.End()
}

func drawScene() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.LoadIdentity()

	for x := uint(0); x < GRIDLENGTH; x++ {
		for y := uint(0); y < GRIDLENGTH; y++ {
			color := 1

			if flagged[x][y] {
				color = 2
			}
			if revealed[x][y] {
				color = 0
			}

			drawSquare(float32(x) * SCALE, float32(y) * SCALE, SCALE, color, grid[x][y])
		}
	}
}