package window

import (
	"poehelper/config"

	imgui "github.com/AllenDang/cimgui-go"
)

func DockbarWindow(isOpen bool) {
	if isOpen {
		if config.App.Dockbar.IsPinned {
			imgui.SetNextWindowSize(config.App.Dockbar.WindowSize)
			imgui.BeginV(config.App.Dockbar.Title, &config.App.Dockbar.IsOpen, config.App.Dockbar.WindowFlags)

			imgui.ButtonV("Home", config.App.Vars.DockbarButton)
			imgui.SameLine()
			if imgui.ButtonV("Met", config.App.Vars.DockbarButton) {
				config.App.Vars.IsOpenMetrics = true
			}
			imgui.SameLine()
			if imgui.ButtonV("Demo", config.App.Vars.DockbarButton) {
				config.App.Vars.IsOpenDemo = true
			}
			imgui.SameLine()
			if imgui.ButtonV("Pin", config.App.Vars.DockbarButton) {
				config.App.Pin.IsOpen = true
				config.App.Dockbar.IsPinned = false
				config.App.Dockbar.WindowFlags ^= imgui.WindowFlagsNoMove
			}
			imgui.SameLine()
			if imgui.ButtonV("Lab", config.App.Vars.DockbarButton) {
				config.App.LabMap.IsOpen = true
			}
			imgui.SameLine()
			if imgui.ButtonV("Set", config.App.Vars.DockbarButton) {
				config.App.Setting.IsOpen = true
			}
			imgui.SameLine()
			if imgui.ButtonV("ESC", config.App.Vars.DockbarButton) {
				config.App.Vars.RootPlatform.GetWindow().SetShouldClose(true)
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
