package window

import (
	"poehelper/config"

	imgui "github.com/AllenDang/cimgui-go"
)

func init() {

}

func TradeWindow(isOpen bool) {
	if isOpen {
		if config.App.Trade.IsPinned {
			imgui.SetNextWindowSize(config.App.Trade.WindowSize)
			imgui.BeginV(config.App.Trade.Title, &config.App.Trade.IsOpen, config.App.Trade.WindowFlags)

			tradeCard()

			imgui.End()
		} else {
			imgui.SetNextWindowSize(config.App.Trade.WindowSize)
			imgui.BeginV(config.App.Trade.Title, &config.App.Trade.IsOpen, config.App.Trade.WindowFlags)

			titleSize := imgui.CalcTextSize(config.App.Trade.Title)
			imgui.Dummy(imgui.Vec2{X: (imgui.ContentRegionAvail().X-titleSize.X)/2 - config.App.Vars.ItemSpacing.X, Y: 19})
			imgui.SameLine()
			imgui.Text(config.App.Trade.Title)

			imgui.End()
		}
	}
}

func tradeCard() {
	// imgui.PushItemWidth(imgui.ContentRegionAvail().X)
	imgui.BeginChildStrV("TradeCard", imgui.Vec2{X: imgui.ContentRegionAvail().X, Y: 65}, true, 0)

	imgui.Button("Hideout")
	imgui.SameLine()
	imgui.Button("Chat")
	imgui.SameLine()
	imgui.Button("Cancel")

	imgui.EndChild()
}
