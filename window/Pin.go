package window

import (
	"poehelper/config"

	imgui "github.com/AllenDang/cimgui-go"
)

func PinWindow(isOpen bool) {
	if isOpen {
		imgui.SetNextWindowPos(config.App.Vars.DisplaySize.Div(2))
		imgui.SetNextWindowSize(config.App.Pin.WindowSize)
		imgui.BeginV(config.App.Pin.Title, nil, config.App.Pin.WindowFlags)
		textPin := "Position the windows as you like"
		textSize := imgui.CalcTextSize(textPin)
		imgui.Dummy(imgui.Vec2{X: (imgui.ContentRegionAvail().X-textSize.X)/2 - config.App.Vars.ItemSpacing.X, Y: 19})
		imgui.SameLine()
		imgui.Text(textPin)
		imgui.Dummy(imgui.Vec2{X: (imgui.ContentRegionAvail().X-config.App.Vars.BaseLongButton.X)/2 - config.App.Vars.ItemSpacing.X, Y: 19})
		imgui.SameLine()
		if imgui.ButtonV("Save", config.App.Vars.BaseLongButton) {
			config.App.Pin.IsOpen = false
			config.App.Dockbar.IsPinned = true
			config.App.Dockbar.WindowFlags ^= imgui.WindowFlagsNoMove
			config.App.Trade.IsOpen = true
			config.App.Trade.WindowFlags ^= imgui.WindowFlagsNoMove
		}
		imgui.End()
	}
}
