package misc

import (
	"fmt"
	imgui "github.com/AllenDang/cimgui-go"
	"poehelper/config"
)

func HelpMarker(text string) {
	imgui.TextDisabled("(?)")
	if imgui.BeginItemTooltip() {
		imgui.PushTextWrapPosV(imgui.FontSize() * 16)
		imgui.TextUnformatted(text)
		imgui.PopTextWrapPos()
		imgui.EndTooltip()
	}
}

func SelectableSettingItem(label string, contentPage func()) {
	imgui.Begin("Settings")
	imgui.BeginChildStr("IDSettingMenu")
	if label != "About" {
		imgui.BeginChildStr("add")
	} else {
		imgui.BeginChildStr("last")
	}
	//---
	isSelect := config.SelectMenuItem == label
	if isSelect {
		imgui.SetItemDefaultFocus()

		imgui.Begin("Settings")
		imgui.BeginChildStr("IDSettingContent")

		contentPage()

		imgui.EndChild()
		imgui.End()
	}
	if imgui.SelectableBoolPtr(fmt.Sprintf(" %s", label), &isSelect) {
		config.SelectMenuItem = label
	}
	imgui.Spacing()
	//---
	imgui.EndChild()
	imgui.EndChild()
	imgui.End()
}
