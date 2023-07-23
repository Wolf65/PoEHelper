package main

import (
	"fmt"
	"os"
	"poehelper/config"
	"poehelper/fonts"
	"poehelper/misc"
	"poehelper/window"
	"regexp"
	"time"

	imgui "github.com/AllenDang/cimgui-go"
	"github.com/hpcloud/tail"
	"github.com/sirupsen/logrus"
)

func init() {
	//check if the required folders and files are available
	executablePath, _ := os.Getwd()
	misc.Log().Debugf(".exe path: %s", executablePath)

	//font
	config.App.Vars.FontDirectory = fmt.Sprintf("%s\\fonts\\ttf\\", executablePath)
	config.App.Vars.LabDirectory = fmt.Sprintf("%s\\lab\\", executablePath)

	if _, err := os.Stat(config.App.Vars.FontDirectory); err != nil {
		if os.IsNotExist(err) {
			//Shows error if file not exists
			misc.Log().Info("Font folder does not exist")
			os.MkdirAll(config.App.Vars.FontDirectory, 0777)
			misc.Log().Info("Font folder create")
			misc.DownloadFile(config.App.Vars.FontDirectory, "fa-brands-400.ttf", "https://github.com/Wolf65/PoEHelper/raw/main/fonts/ttf/fa-brands-400.ttf")
			misc.Log().Info("fa-brands-400.ttf download")
			misc.DownloadFile(config.App.Vars.FontDirectory, "fa-solid-900.ttf", "https://github.com/Wolf65/PoEHelper/raw/main/fonts/ttf/fa-solid-900.ttf")
			misc.Log().Info("fa-solid-900.ttf download")
			misc.DownloadFile(config.App.Vars.FontDirectory, "JetBrainsMono-Medium.ttf", "https://github.com/Wolf65/PoEHelper/raw/main/fonts/ttf/JetBrainsMono-Medium.ttf")
			misc.Log().Info("JetBrainsMono-Medium.ttf download")
		} else {
			misc.Log().Warnf("Font folder err: %s", err)
		}
	} else {
		misc.Log().Info("Font folder exists")
	}

	//lab
	if _, err := os.Stat(config.App.Vars.LabDirectory); err != nil {
		if os.IsNotExist(err) {
			//Shows error if file not exists
			misc.Log().Info("Lab folder does not exist")
			os.MkdirAll(config.App.Vars.LabDirectory, 0777)
			misc.Log().Info("Lab folder create")
		} else {
			misc.Log().Warnf("Lab folder err: %s", err)
		}
	} else {
		misc.Log().Info("Lab folder exists")
	}
}

func main() {

	config.App.Vars.Backend = imgui.CreateBackend(imgui.NewGLFWBackend())
	config.App.Vars.Backend.SetBgColor(imgui.NewVec4(0.5, 0.5, 0.5, 1.0))
	config.App.Vars.Backend.CreateWindow(config.App.Info.ProjectName, 1200, 900, 0)
	config.App.Vars.Backend.SetTargetFPS(60)
	x, y := config.App.Vars.Backend.DisplaySize()

	config.App.Vars.DisplaySize = imgui.Vec2{X: float32(x), Y: float32(y)}

	imgui.CurrentIO().SetConfigFlags(imgui.CurrentIO().ConfigFlags() & ^imgui.ConfigFlagsDockingEnable)

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
	regexpDisplayMessage, _ := regexp.Compile(`^\d{4}/\d{2}/\d{2}\s\d{2}\:\d{2}\:\d{2}\s\d{9}\scffb0719\s\[\w{4}\s\w{6}\s\d{5}\]\s(.*)$`)
	go func() {
		pathLog := "C:\\Program Files (x86)\\Steam\\steamapps\\common\\Path of Exile\\logs\\Client.txt"
		t, err := tail.TailFile(pathLog, tail.Config{
			ReOpen:   true,
			Follow:   true,
			Poll:     true,
			Location: &tail.SeekInfo{Offset: 0, Whence: 2},
		})
		if err != nil {
			misc.Log().WithFields(logrus.Fields{
				"err": err,
			}).Error("Watch PoE Client.txt")
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
					res := regexpDisplayMessage.FindAllStringSubmatch(line.Text, -1)
					if res != nil {
						fmt.Println(res[0][1])
						config.TestLog = res[0][1]
					}
				}
			}
		}

		// for line := range t.Lines {
		// 	fmt.Println(line.Text)
		// }
	}()
	config.App.Vars.Backend.Run(loop)
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
