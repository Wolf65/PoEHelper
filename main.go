package main

import (
	"fmt"
	imgui "github.com/AllenDang/cimgui-go"
	"github.com/hpcloud/tail"
	"github.com/sirupsen/logrus"
	"os"
	"poehelper/config"
	"poehelper/fonts"
	"poehelper/misc"
	"poehelper/window"
	"regexp"
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
			misc.Log().Warn("Font folder does not exist")
			if os.MkdirAll(config.App.Vars.FontDirectory, 0755) != nil {
				misc.Log().WithFields(logrus.Fields{
					"err": err,
				}).Error("Font folder mkdirAll")
			}
			misc.Log().Info("Font folder create")
			misc.CheckAndDownloadFont()
		} else {
			misc.Log().WithFields(logrus.Fields{
				"err": err,
			}).Warn("Font folder")
		}
	} else {
		misc.Log().Debug("Font folder exists")
		misc.CheckAndDownloadFont()
	}

	//lab
	if _, err := os.Stat(config.App.Vars.LabDirectory); err != nil {
		if os.IsNotExist(err) {
			misc.Log().Warn("Lab folder does not exist")
			if os.MkdirAll(config.App.Vars.LabDirectory, 0755) != nil {
				misc.Log().WithFields(logrus.Fields{
					"err": err,
				}).Error("Lab folder mkdirAll")
			}
			misc.Log().Info("Lab folder create")
		} else {
			misc.Log().WithFields(logrus.Fields{
				"err": err,
			}).Warn("Lab folder")
		}
	} else {
		misc.Log().Debug("Lab folder exists")
	}
}

func main() {
	backend := imgui.CreateBackend(imgui.NewGLFWBackend())
	backend.SetBgColor(imgui.NewVec4(0.5, 0.5, 0.5, 1.0))
	backend.SetWindowHint(0x00020004, 0)
	backend.CreateWindow(config.App.Info.ProjectName, 500, 500, imgui.GLFWWindowFlagsFloating)
	backend.SetWindowPos(50, 50)
	backend.SetTargetFPS(60)
	x, y := backend.DisplaySize()

	imgui.CurrentIO().SetIniFilename("imgui.ini")
	config.App.Vars.DisplaySize = imgui.Vec2{X: float32(x), Y: float32(y)}

	imgui.CurrentIO().SetConfigFlags(imgui.CurrentIO().ConfigFlags() & ^imgui.ConfigFlagsDockingEnable)
	imgui.CurrentIO().SetConfigViewportsNoAutoMerge(true)

	fonts.AppendDefaultFont(imgui.CurrentIO())

	regexpDisplayMessage, _ := regexp.Compile(`^\d{4}/\d{2}/\d{2}\s\d{2}\:\d{2}\:\d{2}\s\d{9}\scffb0719\s\[\w{4}\s\w{6}\s\d{5}\]\s(.*)$`)
	go func() {
		pathLog := "C:\\Program Files (x86)\\Steam\\steamapps\\common\\Path of Exile\\logs\\Client.txt"
		t, err := tail.TailFile(pathLog, tail.Config{
			ReOpen:   true,
			Follow:   true,
			Poll:     true,
			Location: &tail.SeekInfo{Offset: 0, Whence: 2},
			Logger:   tail.DiscardingLogger,
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

	backend.SetCloseCallback(func(b imgui.Backend) {
		fmt.Println("window is closing")
	})

	backend.Run(loop)
}

func loop() {
	debug()

	metricsWindow(&config.App.Vars.IsOpenMetrics)
	demoWindow(&config.App.Vars.IsOpenDemo)

	window.DockbarWindow(&config.App.Dockbar.IsOpen)
	window.SettingsWindow(&config.App.Setting.IsOpen)
	window.PinWindow(&config.App.Pin.IsOpen)
	window.LabMapWindow(&config.App.LabMap.IsOpen)
	window.Changelog(&config.App.Changelog.IsOpen)
	window.LabCompassWindow(&config.App.LabCompass.IsOpen)
	window.TradeWindow(&config.App.Trade.IsOpen)

	misc.FileBrowserWindow(&config.App.FileBrowser.IsOpen)
}

func debug() {
}

func metricsWindow(isOpen *bool) {
	if *isOpen {
		imgui.ShowMetricsWindowV(&config.App.Vars.IsOpenMetrics)
	}
}

func demoWindow(isOpen *bool) {
	if *isOpen {
		imgui.ShowDemoWindowV(&config.App.Vars.IsOpenDemo)
	}
}
