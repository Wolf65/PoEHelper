package main

import (
	"fmt"
	"os"
	"poehelper/config"
	"poehelper/fonts"
	"poehelper/platform"
	"poehelper/renderer"
	"poehelper/window"
	"time"

	imgui "github.com/AllenDang/cimgui-go"
)

var (
	err error
)

func main() {
	context := imgui.CreateContext()

	defer context.Destroy()
	io := imgui.CurrentIO()

	config.App.Vars.RootPlatform, err = platform.NewGLFW(io, platform.GLFWClientAPIOpenGL3, 1200, 900, config.App.Info.ProjectName)

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}
	defer config.App.Vars.RootPlatform.Dispose()

	config.App.Vars.RootRenderer, err = renderer.NewOpenGL3(io)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}
	defer config.App.Vars.RootRenderer.Dispose()

	imgui.CurrentIO().SetConfigFlags(imgui.CurrentIO().ConfigFlags() | imgui.ConfigFlagsViewportsEnable)

	fonts.AppendDefaultFont()

	Run(config.App.Vars.RootPlatform, config.App.Vars.RootRenderer)

	// go func() {
	// 	pathLog := "C:\\Program Files (x86)\\Steam\\steamapps\\common\\Path of Exile\\logs\\Client.txt"
	// 	t, err := tail.TailFile(pathLog, tail.Config{
	// 		ReOpen:   true,
	// 		Follow:   true,
	// 		Poll:     true,
	// 		Location: &tail.SeekInfo{Offset: 0, Whence: 2},
	// 	})
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	// 	for {
	// 		select {
	// 		// // if the channel is done, then exit the loop
	// 		// case <-ctx.Done():
	// 		// 	t.Stop()
	// 		// 	t.Cleanup()
	// 		// 	return
	// 		// // get the next log line and echo it out
	// 		case line := <-t.Lines:
	// 			if line != nil {
	// 				// hash cffb0719 print display message
	// 				s := strings.Split(line.Text, " ")
	// 				// r, _ := regexp.Match("cffb0719", []byte(line.Text))
	// 				if s[4] == "cffb0719" {
	// 					config.TestLog = line.Text
	// 				}
	// 			}
	// 		}
	// 	}

	// 	// for line := range t.Lines {
	// 	// 	fmt.Println(line.Text)
	// 	// }
	// }()
	// backend.Run(loop)
}

func Run(p platform.Platform, r renderer.Renderer) {
	clearColor := [3]float32{0.5, 0.5, 0.5}

	for !p.ShouldStop() {
		p.ProcessEvents()

		// Signal start of a new frame
		p.NewFrame()
		imgui.NewFrame()

		loop()

		// Rendering
		imgui.Render() // This call only creates the draw data list. Actual rendering to framebuffer is done below.

		r.PreRender(clearColor)
		// A this point, the application could perform its own rendering...
		// app.RenderScene()

		r.Render(p.DisplaySize(), p.FramebufferSize(), imgui.CurrentDrawData())
		p.PostRender()

		// sleep to avoid 100% CPU usage for this demo
		<-time.After(25 * time.Millisecond)
	}
}

func loop() {
	debug()

	metricsWindow(config.App.Vars.IsOpenMetrics)
	demoWindow(config.App.Vars.IsOpenDemo)

	window.DockbarWindow(config.App.Dockbar.IsOpen)
	window.SettingsWindow(config.App.Setting.IsOpen)
	window.PinWindow(config.App.Pin.IsOpen)
	window.LabMapWindow(config.App.LabMap.IsOpen)
	window.Changelog(config.App.Changelog.IsOpen)
	window.LabCompassWindow(config.App.LabCompass.IsOpen)

}

func debug() {
	imgui.Text(fmt.Sprintf("%c", fonts.ICON_FA_500PX))
}

func metricsWindow(isOpen bool) {
	if isOpen {
		imgui.ShowMetricsWindowV(&config.App.Vars.IsOpenMetrics)
	}
}

func demoWindow(isOpen bool) {
	if isOpen {
		imgui.ShowDemoWindowV(&config.App.Vars.IsOpenDemo)
	}
}
