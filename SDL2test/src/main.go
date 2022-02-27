package main

import (
	"fmt"
	"github.com/FrazorBladezSharp/SDL2test/AmberLib/amberCore"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/vulkan-go/asche"
	"github.com/vulkan-go/vulkan"
)

const WINDOW_WIDTH int32 = 640
const WINDOW_HEIGHT int32 = 480
const WINDOW_TITLE string = "SDL2_Vulkan"

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

	mainApplication := amberCore.NewApplication()

	// create a window
	window, err := sdl.CreateWindow(
		WINDOW_TITLE,
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		WINDOW_WIDTH,
		WINDOW_HEIGHT,
		sdl.WINDOW_VULKAN)

	checkErr(err)

	mainApplication.WindowHandle = window

	// init the GFX platform
	platform, err := asche.NewPlatform(mainApplication)
	checkErr(err)

	mainApplication.VulkanSwapchainDimensions(
		uint32(WINDOW_WIDTH),
		uint32(WINDOW_HEIGHT))

	swapchain := mainApplication.Context().SwapchainDimensions()

	fmt.Printf("Vulkan: %s with %+v swapchain\n",
		WINDOW_TITLE,
		swapchain)

	//window.Show()
	// Delta Time
	// main loop

	// systemize the vulkan workflow

	// clean up

	mainApplication.Destroy()
	platform.Destroy()
	err = window.Destroy()
	if err != nil {
		fmt.Println("[main] clean up: Unable to destroy window.")
	}
	sdl.VulkanUnloadLibrary()
	sdl.Quit()

}
