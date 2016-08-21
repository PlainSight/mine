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

var (
	WIDTH int = 720
	HEIGHT int = 720
	XSCALE float32 = float32(WIDTH) / 16
	YSCALE float32 = float32(HEIGHT) / 16
	numbers uint32
)

func handleResize(w *glfw.Window, width int, height int) {
	WIDTH = width
	HEIGHT = height
	XSCALE = float32(WIDTH / 16)
	YSCALE = float32(HEIGHT / 16)
	setupScene()
	drawScene()
}

func handleClick(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
	if action != glfw.Press {
		return
	}

	xpos, ypos := w.GetCursorPos()

	if button == glfw.MouseButton2 {
		flagClick(uint(xpos/float64(XSCALE)), uint(ypos/float64(YSCALE)))
	}

	if button == glfw.MouseButton1 {
		revealClick(uint(xpos/float64(XSCALE)), uint(ypos/float64(YSCALE)))
	}
}

func setup() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	window, err := glfw.CreateWindow(WIDTH, HEIGHT, "Cube", nil, nil)
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
	window.SetSizeCallback(handleResize);

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
	gl.Ortho(0, float64(WIDTH), float64(HEIGHT), 0, -1, 1)	
	gl.Viewport(0, 0, int32(WIDTH), int32(HEIGHT))	
}

func destroyScene() {
}

func drawNumber(x float32, y float32, w float32, h float32, number int) {

	xmargin := w / 4
	ymargin := h / 4

	x1 := x + xmargin
	y1 := y + ymargin
	x2 := x - xmargin + w
	y2 := y - ymargin + h

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

func drawSquare(x float32, y float32, w float32, h float32, color int, count int) {

	x1 := x
	y1 := y
	x2 := x + w
	y2 := y + h

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
		drawNumber(x, y, w, h, count)
	}
}

func drawScene() {
	gl.Clear(gl.COLOR_BUFFER_BIT)

	for x := uint(0); x < GRIDLENGTH; x++ {
		for y := uint(0); y < GRIDLENGTH; y++ {
			color := 1

			if flagged[x][y] {
				color = 2
			}
			if revealed[x][y] {
				color = 0
			}

			drawSquare(float32(x) * XSCALE, float32(y) * YSCALE, XSCALE, YSCALE, color, grid[x][y])
		}
	}
}
