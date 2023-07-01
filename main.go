package main

import (
	"fmt"
	"os"
	"poehelper/config"
	"poehelper/platform"
	"poehelper/renderer"
	"poehelper/window"
	"time"

	imgui "github.com/AllenDang/cimgui-go"
)

var (
	Tststr = "test.ini"
)

// заменить этим

// func (self IO) SetIniFilename(v string) {
// 	vArg, vFin := wrapString(v)
// 	C.wrap_ImGuiIO_SetIniFilename(self.handle(), vArg)

// 	vFin()
// }

func main() {
	context := imgui.CreateContext()

	defer context.Destroy()
	io := imgui.CurrentIO()

	platform, err := platform.NewGLFW(io, platform.GLFWClientAPIOpenGL3)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}
	defer platform.Dispose()

	renderer, err := renderer.NewOpenGL3(io)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}
	defer renderer.Dispose()

	Run(platform, renderer)

	// backend := imgui.CreateBackend(imgui.NewGLFWBackend())
	// backend.SetBgColor(imgui.NewVec4(0.45, 0.55, 0.6, 1.0))
	// backend.CreateWindow(config.App.Info.ProjectName, 1200, 900, 0)
	// backend.SetTargetFPS(60)
	// x, y := backend.DisplaySize()
	// config.App.Vars.DisplaySize = imgui.Vec2{X: float32(x), Y: float32(y)}

	// imgui.CurrentIO().SetConfigFlags(imgui.CurrentIO().ConfigFlags() & ^imgui.ConfigFlagsDockingEnable)

	// // imgui.CurrentIO().SetIniFilename(Tststr)

	// // fmt.Println(imgui.CurrentIO().IniFilename())
	// fonts.AppendDefaultFont(imgui.CurrentIO())

	// config.App.Vars.ItemSpacing = imgui.CurrentIO().Ctx().Style().ItemSpacing()

	// refreshTicket := time.NewTicker(500 * time.Millisecond)
	// refreshTicketDone := make(chan bool)

	// go func() {
	// 	for {
	// 		select {
	// 		case <-refreshTicketDone:
	// 			return
	// 		case <-refreshTicket.C:
	// 			imgui.GetBackend().Refresh()
	// 		}
	// 	}
	// }()
	// defer func() {
	// 	refreshTicket.Stop()
	// 	refreshTicketDone <- true
	// }()

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

// Platform covers mouse/keyboard/gamepad inputs, cursor shape, timing, windowing.
type Platform interface {
	// ShouldStop is regularly called as the abort condition for the program loop.
	ShouldStop() bool
	// ProcessEvents is called once per render loop to dispatch any pending events.
	ProcessEvents()
	// DisplaySize returns the dimension of the display.
	DisplaySize() [2]float32
	// FramebufferSize returns the dimension of the framebuffer.
	FramebufferSize() [2]float32
	// NewFrame marks the begin of a render pass. It must update the cimgui IO state according to user input (mouse, keyboard, ...)
	NewFrame()
	// PostRender marks the completion of one render pass. Typically this causes the display buffer to be swapped.
	PostRender()
	// ClipboardText returns the current text of the clipboard, if available.
	ClipboardText() (string, error)
	// SetClipboardText sets the text as the current text of the clipboard.
	SetClipboardText(text string)
}

type clipboard struct {
	platform Platform
}

func (board clipboard) Text() (string, error) {
	return board.platform.ClipboardText()
}

func (board clipboard) SetText(text string) {
	board.platform.SetClipboardText(text)
}

// Renderer covers rendering cimgui draw data.
type Renderer interface {
	// PreRender causes the display buffer to be prepared for new output.
	PreRender(clearColor [3]float32)
	// Render draws the provided cimgui draw data.
	Render(displaySize [2]float32, framebufferSize [2]float32, drawData imgui.DrawData)
}

const (
	millisPerSecond = 1000
	sleepDuration   = time.Millisecond * 25
)

func Run(p Platform, r Renderer) {
	clearColor := [3]float32{0, 0, 0}
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
		<-time.After(sleepDuration)
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
