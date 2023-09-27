package config

import (
	imgui "github.com/AllenDang/cimgui-go"
)

var (
	TestLog        string
	Backend        imgui.Backend[imgui.GLFWWindowFlags]
	SelectMenuItem = "General"
)

type application struct {
	Info info

	Vars vars

	Setting     setting
	Pin         pin
	Dockbar     dockbar
	LabMap      labMap
	LabCompass  labCompass
	Changelog   changelog
	Trade       trade
	FileBrowser fileBrowser
}
type info struct {
	ProjectName    string
	ProjectVersion string
}
type vars struct {
	ExecutablePath string
	FontDirectory  string
	LabDirectory   string
	//
	BaseButton     imgui.Vec2
	BaseLongButton imgui.Vec2
	SquareButton   imgui.Vec2
	//
	IsOpenMetrics bool
	IsOpenDemo    bool
}
type setting struct {
	Title       string
	IsOpen      bool
	IsPinned    bool
	WindowSize  imgui.Vec2
	WindowFlags imgui.WindowFlags
	//
	SelectPage string
}
type pin struct {
	Title       string
	IsOpen      bool
	WindowSize  imgui.Vec2
	WindowFlags imgui.WindowFlags
}
type dockbar struct {
	Title       string
	IsOpen      bool
	IsPinned    bool
	WindowFlags imgui.WindowFlags
}
type labMap struct {
	Title       string
	IsOpen      bool
	IsPinned    bool
	WindowSize  imgui.Vec2
	WindowFlags imgui.WindowFlags
}
type labCompass struct {
	Title       string
	IsOpen      bool
	IsPinned    bool
	WindowSize  imgui.Vec2
	WindowFlags imgui.WindowFlags
}
type changelog struct {
	Title       string
	IsOpen      bool
	IsPinned    bool
	WindowSize  imgui.Vec2
	WindowFlags imgui.WindowFlags

	ChangelogPath string
}
type trade struct {
	Title       string
	IsOpen      bool
	IsPinned    bool
	WindowSize  imgui.Vec2
	WindowFlags imgui.WindowFlags
}
type fileBrowser struct {
	Title       string
	IsOpen      bool
	WindowSize  imgui.Vec2
	WindowFlags imgui.WindowFlags
}

//init

var App = application{
	Info: info{
		ProjectName:    "PoEHelper",
		ProjectVersion: "0.0.1",
	},
	//
	Vars: vars{
		ExecutablePath: "",
		FontDirectory:  "",
		LabDirectory:   "",

		BaseButton:     imgui.Vec2{X: 80, Y: 19},
		BaseLongButton: imgui.Vec2{X: 100, Y: 19},
		SquareButton:   imgui.Vec2{X: 46, Y: 46},
		//
		IsOpenMetrics: false,
		IsOpenDemo:    false,
	},
	//
	Setting: setting{
		Title:       "Settings",
		IsOpen:      false,
		IsPinned:    false,
		WindowSize:  imgui.Vec2{X: 700, Y: 500},
		WindowFlags: imgui.WindowFlagsNoResize,
		//
		SelectPage: "pageGeneral",
	},
	Pin: pin{
		Title:       "Pin",
		IsOpen:      false,
		WindowSize:  imgui.Vec2{X: 250, Y: 77},
		WindowFlags: imgui.WindowFlagsNoResize | imgui.WindowFlagsNoCollapse,
	},
	Dockbar: dockbar{
		Title:       "Dockbar",
		IsOpen:      true,
		IsPinned:    true,
		WindowFlags: imgui.WindowFlagsNoResize | imgui.WindowFlagsNoTitleBar | imgui.WindowFlagsNoMove,
	},
	LabMap: labMap{
		Title:       "Lab map",
		IsOpen:      false,
		IsPinned:    false,
		WindowSize:  imgui.Vec2{X: 700, Y: 300},
		WindowFlags: imgui.WindowFlagsNoResize,
	},
	LabCompass: labCompass{
		Title:       "Lab compass",
		IsOpen:      false,
		IsPinned:    false,
		WindowSize:  imgui.Vec2{X: 168, Y: 200},
		WindowFlags: imgui.WindowFlagsNoResize,
	},
	Changelog: changelog{
		Title:       "Changelog",
		IsOpen:      false,
		IsPinned:    false,
		WindowSize:  imgui.Vec2{X: 500, Y: 600},
		WindowFlags: imgui.WindowFlagsNoResize,

		ChangelogPath: "Changelog.txt",
	},
	Trade: trade{
		Title:       "Trade",
		IsOpen:      true,
		IsPinned:    true,
		WindowSize:  imgui.Vec2{X: 400, Y: 200},
		WindowFlags: imgui.WindowFlagsNoResize, // | imgui.WindowFlagsNoMove,
	},
	FileBrowser: fileBrowser{
		Title:       "File browser",
		IsOpen:      false,
		WindowSize:  imgui.Vec2{X: 500, Y: 300},
		WindowFlags: imgui.WindowFlagsNoCollapse | imgui.WindowFlagsNoResize,
	},
}
