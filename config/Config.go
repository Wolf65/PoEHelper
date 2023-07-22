package config

import (
	imgui "github.com/AllenDang/cimgui-go"
)

var (
	TestLog string
)

type application struct {
	Info info

	Vars vars

	Setting    setting
	Pin        pin
	Dockbar    dockbar
	LabMap     labMap
	LabCompass labCompass
	Changelog  changelog
}
type info struct {
	ProjectName    string
	ProjectVersion string
}
type vars struct {
	Backend imgui.Backend
	//
	ExecutablePath string
	FontDirectory  string
	LabDirectory   string
	//
	DisplaySize    imgui.Vec2
	BaseButton     imgui.Vec2
	BaseLongButton imgui.Vec2
	DockbarButton  imgui.Vec2
	ItemSpacing    imgui.Vec2
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
	WindowSize  imgui.Vec2
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

//init

var App = application{
	Info: info{
		ProjectName:    "PoEHelper",
		ProjectVersion: "0.0.1 dev",
	},
	//
	Vars: vars{
		Backend: nil,
		// set in main.go DisplaySize

		ExecutablePath: "",
		FontDirectory:  "",
		LabDirectory:   "",

		BaseButton:     imgui.Vec2{X: 80, Y: 19},
		BaseLongButton: imgui.Vec2{X: 100, Y: 19},
		// set in main.go ItemSpacing
		DockbarButton: imgui.Vec2{X: 45, Y: 45},
		//
		IsOpenMetrics: false,
		IsOpenDemo:    false,
	},
	//
	Setting: setting{
		Title:       "Settings",
		IsOpen:      false,
		IsPinned:    false,
		WindowSize:  imgui.Vec2{X: 600, Y: 500},
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
		WindowSize:  imgui.Vec2{X: 380, Y: 61},
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
}
