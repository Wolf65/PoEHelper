package window

import (
	"fmt"
	"poehelper/config"

	imgui "github.com/AllenDang/cimgui-go"
)

var (
	countTrade uint
	card       = []string{"1", "2", "3"}
)

func init() {

}

func TradeWindow(isOpen bool) {
	if isOpen {
		if config.App.Trade.IsPinned {
			imgui.SetNextWindowSize(config.App.Trade.WindowSize)
			imgui.BeginV(fmt.Sprintf("%s (%v trade offers)###Trade", config.App.Trade.Title, len(card)), &config.App.Trade.IsOpen, config.App.Trade.WindowFlags)

			for _, s := range card {
				tradeCard(s)
			}

			imgui.End()
		} else {
			imgui.SetNextWindowSize(config.App.Trade.WindowSize)
			imgui.BeginV(config.App.Trade.Title, &config.App.Trade.IsOpen, config.App.Trade.WindowFlags)

			titleSize := imgui.CalcTextSize(config.App.Trade.Title)
			imgui.Dummy(imgui.Vec2{X: (imgui.ContentRegionAvail().X-titleSize.X)/2 - config.App.Vars.ItemSpacing.X, Y: 19})
			imgui.SameLine()
			imgui.Text(config.App.Trade.Title)
			imgui.Text(fmt.Sprintf("%v", config.App.Trade.IsPinned))

			imgui.End()
		}
	}
}

func tradeCard(id string) {
	// imgui.PushItemWidth(imgui.ContentRegionAvail().X)
	imgui.BeginChildStrV(fmt.Sprintf("TradeCard##%s", id), imgui.Vec2{X: imgui.ContentRegionAvail().X, Y: 65}, true, 0)

	imgui.Text(fmt.Sprintf("ID: %s", id))

	imgui.Button("Hideout")
	imgui.SameLine()
	imgui.Button("Chat")
	imgui.SameLine()
	if imgui.Button("Cancel") {
		card = append(card[:1], card[2:]...)
	}

	imgui.EndChild()
}
