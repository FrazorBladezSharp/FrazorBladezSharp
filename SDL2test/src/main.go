package main

import (
	"fmt"
	"github.com/FrazorBladezSharp/SDL2test/AmberLib/amberCore"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/vulkan-go/asche"
	"github.com/vulkan-go/vulkan"
)

const WindowWidth int32 = 640
const WindowHeight int32 = 480
const WindowTitle string = "SDL2_Vulkan"

func checkErr(err error) {
	if err == nil {
		return
	}

	panic(err)
}

func main() {

	var err error

	// setup system to support vulkan
	err = sdl.Init(sdl.INIT_EVERYTHING)
	checkErr(err)

	// init vulkan
	err = sdl.VulkanLoadLibrary("")
	checkErr(err)

	vulkan.SetGetInstanceProcAddr(
		sdl.VulkanGetVkGetInstanceProcAddr())

	err = vulkan.Init()
	checkErr(err)

	// app system
	mainApplication := amberCore.NewApplication(
		"SDL2test",
		false)

	// create a window
	window, err := sdl.CreateWindow(
		WindowTitle,
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		WindowWidth,
		WindowHeight,
		sdl.WINDOW_VULKAN)

	checkErr(err)

	mainApplication.WindowHandle = window

	// init the GFX platform
	platform, err := asche.NewPlatform(mainApplication)
	checkErr(err)

	mainApplication.VulkanSwapchainDimensions(
		uint32(WindowWidth),
		uint32(WindowHeight))

	swapchain := mainApplication.Context().SwapchainDimensions()

	fmt.Printf("Vulkan: %s with %+v swapchain\n",
		WindowTitle,
		swapchain)

	// Delta Time
	//fpsDelay := time.Second / 60
	//fpsTicker := time.NewTicker(fpsDelay)
	//start := time.Now()
	//frames := 0

	// main loop
	running := true
	for running {

		// poll events
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}

		// systemize the vulkan workflow
	}

	// clean up
	platform.Destroy()
	mainApplication.Destroy()
	sdl.VulkanUnloadLibrary()
	sdl.Quit()

}
