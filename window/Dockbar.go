package window

import (
	"poehelper/config"
	"poehelper/fonts"

	imgui "github.com/AllenDang/cimgui-go"
)

func DockbarWindow(isOpen *bool) {
	style := imgui.CurrentIO().Ctx().Style()
	winSize := imgui.Vec2{
		X: style.WindowPadding().X*2 + config.App.Vars.SquareButton.X*7 + style.ItemSpacing().X*6,
		Y: style.WindowPadding().Y*2 + config.App.Vars.SquareButton.Y,
	}
	if *isOpen {
		if config.App.Dockbar.IsPinned {
			imgui.SetNextWindowSize(winSize)
			imgui.SetNextWindowPosV(imgui.Vec2{X: 3440/2 - winSize.X/2, Y: 1200}, imgui.CondFirstUseEver, imgui.Vec2{X: 0, Y: 0})
			imgui.BeginV(config.App.Dockbar.Title, &config.App.Dockbar.IsOpen, config.App.Dockbar.WindowFlags)

			imgui.ButtonV(fonts.IconsFontAwesome6.Icons["House"], config.App.Vars.SquareButton)
			imgui.SameLine()
			if imgui.ButtonV(fonts.IconsFontAwesome6.Icons["Info"], config.App.Vars.SquareButton) {
				config.App.Vars.IsOpenMetrics = true
			}
			imgui.SameLine()
			if imgui.ButtonV("Demo", config.App.Vars.SquareButton) {
				config.App.Vars.IsOpenDemo = true
			}
			imgui.SameLine()
			if imgui.ButtonV(fonts.IconsFontAwesome6.Icons["Thumbtack"], config.App.Vars.SquareButton) {
				config.App.Pin.IsOpen = true
				config.App.Dockbar.IsPinned = false
				config.App.Dockbar.WindowFlags ^= imgui.WindowFlagsNoMove
				config.App.Trade.IsPinned = false
				config.App.Trade.WindowFlags ^= imgui.WindowFlagsNoMove
			}
			imgui.SameLine()
			if imgui.ButtonV("Lab", config.App.Vars.SquareButton) {
				config.App.LabMap.IsOpen = true
			}
			imgui.SameLine()
			if imgui.ButtonV(fonts.IconsFontAwesome6.Icons["Gear"], config.App.Vars.SquareButton) {
				config.App.Setting.IsOpen = true
			}
			imgui.SameLine()
			if imgui.ButtonV(fonts.IconsFontAwesome6.Icons["ArrowRightFromBracket"], config.App.Vars.SquareButton) {
				imgui.GetBackend().SetShouldClose(true)
			}

			imgui.End()
		} else {
			imgui.SetNextWindowSize(winSize)
			imgui.BeginV(config.App.Dockbar.Title, &config.App.Dockbar.IsOpen, config.App.Dockbar.WindowFlags)

			titleSize := imgui.CalcTextSize(config.App.Dockbar.Title)
			imgui.Dummy(imgui.Vec2{
				X: (imgui.ContentRegionAvail().X-titleSize.X)/2 - style.ItemSpacing().X,
				Y: 0,
			})
			imgui.SameLine()
			imgui.Text(config.App.Dockbar.Title)

			imgui.End()
		}
	}
}
