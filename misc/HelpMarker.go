package misc

import imgui "github.com/AllenDang/cimgui-go"

func HelpMarker(text string) {
	imgui.TextDisabled("(?)")
	if imgui.BeginItemTooltip() {
		imgui.PushTextWrapPosV(imgui.FontSize() * 16)
		imgui.TextUnformatted(text)
		imgui.PopTextWrapPos()
		imgui.EndTooltip()
	}
}
