package window

import (
	"poehelper/config"

	imgui "github.com/AllenDang/cimgui-go"
)

func PinWindow(isOpen *bool) {
	style := imgui.CurrentIO().Ctx().Style()
	textPin := "Position the windows as you like"
	textSize := imgui.CalcTextSize(textPin)
	if *isOpen {
		imgui.SetNextWindowSize(config.App.Pin.WindowSize)
		imgui.SetNextWindowPosV(imgui.Vec2{
			X: 3440/2 - config.App.Pin.WindowSize.X/2,
			Y: 1440/2 - config.App.Pin.WindowSize.Y/2},
			imgui.CondAlways, imgui.Vec2{X: 0, Y: 0})
		imgui.BeginV(config.App.Pin.Title, nil, config.App.Pin.WindowFlags)
		imgui.Dummy(imgui.Vec2{
			X: (imgui.ContentRegionAvail().X-textSize.X)/2 - style.ItemSpacing().X,
			Y: imgui.ContentRegionAvail().Y - style.ItemSpacing().Y - config.App.Vars.BaseLongButton.Y,
		})
		imgui.SameLine()
		imgui.Text(textPin)
		imgui.Dummy(imgui.Vec2{X: (imgui.ContentRegionAvail().X-config.App.Vars.BaseLongButton.X)/2 - style.ItemSpacing().X, Y: 0})
		imgui.SameLine()
		if imgui.ButtonV("Save", config.App.Vars.BaseLongButton) {
			config.App.Pin.IsOpen = false
			config.App.Dockbar.IsPinned = true
			config.App.Dockbar.WindowFlags ^= imgui.WindowFlagsNoMove
			config.App.Trade.IsPinned = true
			config.App.Trade.WindowFlags ^= imgui.WindowFlagsNoMove
		}
		imgui.End()
	}
}
