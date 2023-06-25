package main

import (
	"fmt"
	"poehelper/config"
	"poehelper/fonts"
	"poehelper/window"
	"strings"
	"time"

	imgui "github.com/AllenDang/cimgui-go"
	"github.com/hpcloud/tail"
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
	backend := imgui.CreateBackend(imgui.NewGLFWBackend())
	backend.SetBgColor(imgui.NewVec4(0.45, 0.55, 0.6, 1.0))
	backend.CreateWindow(config.App.Info.ProjectName, 1200, 900, 0)
	backend.SetTargetFPS(60)
	x, y := backend.DisplaySize()
	config.App.Vars.DisplaySize = imgui.Vec2{X: float32(x), Y: float32(y)}

	imgui.CurrentIO().SetConfigFlags(imgui.CurrentIO().ConfigFlags() & ^imgui.ConfigFlagsDockingEnable)

	// imgui.CurrentIO().SetIniFilename(Tststr)

	// fmt.Println(imgui.CurrentIO().IniFilename())
	fonts.AppendDefaultFont(imgui.CurrentIO())

	config.App.Vars.ItemSpacing = imgui.CurrentIO().Ctx().Style().ItemSpacing()

	refreshTicket := time.NewTicker(500 * time.Millisecond)
	refreshTicketDone := make(chan bool)

	go func() {
		for {
			select {
			case <-refreshTicketDone:
				return
			case <-refreshTicket.C:
				imgui.GetBackend().Refresh()
			}
		}
	}()
	defer func() {
		refreshTicket.Stop()
		refreshTicketDone <- true
	}()

	go func() {
		pathLog := "C:\\Program Files (x86)\\Steam\\steamapps\\common\\Path of Exile\\logs\\Client.txt"
		t, err := tail.TailFile(pathLog, tail.Config{
			ReOpen:   true,
			Follow:   true,
			Poll:     true,
			Location: &tail.SeekInfo{Offset: 0, Whence: 2},
		})
		if err != nil {
			fmt.Println(err)
		}

		for {
			select {
			// // if the channel is done, then exit the loop
			// case <-ctx.Done():
			// 	t.Stop()
			// 	t.Cleanup()
			// 	return
			// // get the next log line and echo it out
			case line := <-t.Lines:
				if line != nil {
					// hash cffb0719 print display message
					s := strings.Split(line.Text, " ")
					// r, _ := regexp.Match("cffb0719", []byte(line.Text))
					if s[4] == "cffb0719" {
						config.TestLog = line.Text
					}
				}
			}
		}

		// for line := range t.Lines {
		// 	fmt.Println(line.Text)
		// }
	}()
	backend.Run(loop)
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
