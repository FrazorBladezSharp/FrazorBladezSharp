package main

import (
	"fmt"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/veandco/go-sdl2/sdl"
)

func errorCheck(err error) {

	if err != nil {
		panic(err)
	}

}

func main() {

	var winTitle = "Go-SDL2 + Go-GL"
	var winWidth, winHeight int32 = 800, 600
	var window *sdl.Window
	var context sdl.GLContext
	var event sdl.Event
	var running bool
	var err error

	err = sdl.Init(sdl.INIT_EVERYTHING)
	errorCheck(err)

	window, err = sdl.CreateWindow(
		winTitle,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		winWidth,
		winHeight,
		sdl.WINDOW_OPENGL)
	errorCheck(err)

	context, err = window.GLCreateContext()
	errorCheck(err)

	err = gl.Init()
	errorCheck(err)

	gl.Enable(gl.DEPTH_TEST)
	gl.ClearColor(0.05, 0.05, 0.10, 1.0)
	gl.ClearDepth(1)
	gl.DepthFunc(gl.LEQUAL)
	gl.Viewport(0, 0, winWidth, winHeight)

	running = true

	for running {

		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {

			switch t := event.(type) {

			case *sdl.QuitEvent:
				running = false

			case *sdl.MouseMotionEvent:
				fmt.Printf(
					"[%d ms] MouseMotion\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n",
					t.Timestamp,
					t.Which,
					t.X,
					t.Y,
					t.XRel,
					t.YRel)

			}

		}

		drawgl()

		window.GLSwap()

		sdl.Delay(10)

	}

	// cleanup
	sdl.GLDeleteContext(context)

	err = window.Destroy()
	errorCheck(err)

	sdl.Quit()
}

func drawgl() {

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.Begin(gl.TRIANGLES)

	gl.Color3f(1.0, 0.0, 0.0)
	gl.Vertex2f(-0.5, -0.5)
	gl.Color3f(0.0, 1.0, 0.0)
	gl.Vertex2f(0.5, -0.5)
	gl.Color3f(0.0, 0.0, 1.0)
	gl.Vertex2f(0.0, 0.5)

	gl.End()

}
