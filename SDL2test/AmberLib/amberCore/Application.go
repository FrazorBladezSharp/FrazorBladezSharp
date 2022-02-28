package amberCore

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/vulkan-go/asche"
	"github.com/vulkan-go/vulkan"
	"log"
)

type Application struct {
	WindowHandle *sdl.Window

	name         string
	debugEnabled bool

	asche.BaseVulkanApp
}

func NewApplication(name string, enableDebug bool) *Application {
	a := &Application{
		name:         name,
		debugEnabled: enableDebug,
	}
	return a
}

func (a *Application) VulkanSurface(instance vulkan.Instance) (surface vulkan.Surface) {
	surfPtr, err := a.WindowHandle.VulkanCreateSurface(instance)
	if err != nil {
		log.Println("vulkan error:", err)
		return vulkan.NullSurface
	}

	surf := vulkan.SurfaceFromPointer(uintptr(surfPtr))
	return surf
}

func (a *Application) VulkanAppName() string {
	return a.name
}

func (a *Application) VulkanLayers() []string {
	return []string{
		// "VK_LAYER_GOOGLE_threading",
		// "VK_LAYER_LUNARG_parameter_validation",
		// "VK_LAYER_LUNARG_object_tracker",
		// "VK_LAYER_LUNARG_core_validation",
		// "VK_LAYER_LUNARG_api_dump",
		// "VK_LAYER_LUNARG_swapchain",
		// "VK_LAYER_GOOGLE_unique_objects",
	}
}

func (a *Application) VulkanDebug() bool {
	return a.debugEnabled
}

func (a *Application) VulkanDeviceExtensions() []string {
	return []string{
		"VK_KHR_swapchain",
	}
}

func (a *Application) VulkanSwapchainDimensions(width, height uint32) *asche.SwapchainDimensions {
	return &asche.SwapchainDimensions{
		Width:  width,
		Height: height,
		Format: vulkan.FormatB8g8r8a8Unorm,
	}
}

func (a *Application) VulkanInstanceExtensions() []string {
	extensions := a.WindowHandle.VulkanGetInstanceExtensions()
	if a.debugEnabled {
		extensions = append(extensions, "VK_EXT_debug_report")
	}
	return extensions
}

func (a *Application) Destroy() {
	err := a.WindowHandle.Destroy()
	if err != nil {
		fmt.Println("[Application] Destroy: Unable to destroy the Window Handle")
	}
}
