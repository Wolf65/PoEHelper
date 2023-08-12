package window

import (
	"poehelper/config"
	"poehelper/fonts"

	imgui "github.com/AllenDang/cimgui-go"
)

func DockbarWindow(isOpen bool) {
	if isOpen {
		imgui.SetNextWindowPos(imgui.Vec2{X: 200, Y: 700})
		if config.App.Dockbar.IsPinned {
			imgui.SetNextWindowSize(config.App.Dockbar.WindowSize)
			imgui.BeginV(config.App.Dockbar.Title, &config.App.Dockbar.IsOpen, config.App.Dockbar.WindowFlags)

			imgui.ButtonV(fonts.IconsFontAwesome6.Icons["House"], config.App.Vars.DockbarButton)
			imgui.SameLine()
			if imgui.ButtonV(fonts.IconsFontAwesome6.Icons["Info"], config.App.Vars.DockbarButton) {
				config.App.Vars.IsOpenMetrics = true
			}
			imgui.SameLine()
			if imgui.ButtonV("Demo", config.App.Vars.DockbarButton) {
				config.App.Vars.IsOpenDemo = true
			}
			imgui.SameLine()
			if imgui.ButtonV(fonts.IconsFontAwesome6.Icons["Thumbtack"], config.App.Vars.DockbarButton) {
				config.App.Pin.IsOpen = true
				config.App.Dockbar.IsPinned = false
				config.App.Dockbar.WindowFlags ^= imgui.WindowFlagsNoMove
				config.App.Trade.IsPinned = false
				config.App.Trade.WindowFlags ^= imgui.WindowFlagsNoMove
			}
			imgui.SameLine()
			if imgui.ButtonV("Lab", config.App.Vars.DockbarButton) {
				config.App.LabMap.IsOpen = true
			}
			imgui.SameLine()
			if imgui.ButtonV(fonts.IconsFontAwesome6.Icons["Gear"], config.App.Vars.DockbarButton) {
				config.App.Setting.IsOpen = true
			}
			imgui.SameLine()
			if imgui.ButtonV(fonts.IconsFontAwesome6.Icons["ArrowRightFromBracket"], config.App.Vars.DockbarButton) {
				imgui.GetBackend().SetShouldClose(true)
			}

			imgui.End()
		} else {
			imgui.SetNextWindowSize(config.App.Dockbar.WindowSize)
			imgui.BeginV(config.App.Dockbar.Title, &config.App.Dockbar.IsOpen, config.App.Dockbar.WindowFlags)

			titleSize := imgui.CalcTextSize(config.App.Dockbar.Title)
			imgui.Dummy(imgui.Vec2{X: (imgui.ContentRegionAvail().X-titleSize.X)/2 - config.App.Vars.ItemSpacing.X, Y: 19})
			imgui.SameLine()
			imgui.Text(config.App.Dockbar.Title)

			imgui.End()
		}
	}
}
