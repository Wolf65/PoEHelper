package misc

import (
	imgui "github.com/AllenDang/cimgui-go"
	"poehelper/config"
)

var buf string

func FileBrowserWindow(isOpen *bool) {
	style := imgui.CurrentIO().Ctx().Style()
	if *isOpen {
		imgui.SetNextWindowSize(config.App.FileBrowser.WindowSize)
		imgui.BeginV(config.App.FileBrowser.Title, nil, config.App.FileBrowser.WindowFlags)

		imgui.SetNextItemWidth(40)
		if imgui.BeginComboV("##combo", "C:", 0) {
			imgui.EndCombo()
		}
		imgui.SameLine()
		imgui.SetNextItemWidth(imgui.ContentRegionAvail().X)
		imgui.InputTextWithHint("##pathFileBrowserl", "C:\\", &buf, 0, nil)

		heightChild := imgui.ContentRegionAvail().Y - config.App.Vars.BaseButton.Y - style.ItemSpacing().Y
		imgui.BeginChildStrV("fileBrowser", imgui.Vec2{X: 0, Y: heightChild}, true, 0)
		imgui.EndChild()

		if imgui.Button("Close") {
			config.App.FileBrowser.IsOpen = false
		}
		imgui.SameLine()
		imgui.Button("Select")

		imgui.End()
	}
}
