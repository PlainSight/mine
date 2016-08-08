package main

import (
	"log"
	"runtime"
	"bytes"
	"image"
	"image/draw"
	_ "image/png"

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

var (
	numbers uint32
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

	numbers = loadNumbersTexture()
	defer gl.DeleteTextures(1, &numbers)

	setupScene()

	window.SetMouseButtonCallback(handleClick);

	for !window.ShouldClose() {
		drawScene()
		window.SwapBuffers()
		glfw.WaitEvents()
	}
}

func loadNumbersTexture() uint32 {
	imageBytes, _ := numbersPngBytes()

	img, _, _ := image.Decode(bytes.NewReader(imageBytes))

	rgba := image.NewRGBA(img.Bounds())

	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	var texture uint32
	gl.Enable(gl.TEXTURE_2D)
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))

	return texture
}

func setupScene() {
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
}

func destroyScene() {
}

func drawNumber(x float32, y float32, w float32, number int) {

	margin := w / 4

	x1 := ((x + margin) - HALF) / HALF
	y1 := (HALF - (y + margin)) / HALF
	x2 := (((x - margin) + w) - HALF) / HALF
	y2 := (HALF - ((y - margin) + w)) / HALF

	gl.BindTexture(gl.TEXTURE_2D, numbers)

	gl.Color4f(1, 1, 1, 1)

	gl.Begin(gl.QUADS)

	txmin := float32(number) / 10.0
	txmax := float32(number + 1) / 10.0

	gl.TexCoord2f(txmin, 0)
	gl.Vertex3f(x1, y1, 1)
	gl.TexCoord2f(txmax, 0)
	gl.Vertex3f(x2, y1, 1)
	gl.TexCoord2f(txmax, 1)
	gl.Vertex3f(x2, y2, 1)
	gl.TexCoord2f(txmin, 1)
	gl.Vertex3f(x1, y2, 1)

	gl.End()
}

func drawSquare(x float32, y float32, w float32, color int, count int) {

	x1 := (x - HALF) / HALF
	y1 := (HALF - y) / HALF
	x2 := ((x + w) - HALF) / HALF
	y2 := (HALF - (y + w)) / HALF

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

	gl.End()

	if color == 0 && count > 0 {
		drawNumber(x, y, w, count)
	}
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