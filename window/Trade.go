package window

import (
	"fmt"
	imgui "github.com/AllenDang/cimgui-go"
	"poehelper/config"
	"poehelper/fonts"
	"poehelper/misc"
)

var (
	cards = []Trade{
		{ID: 1, Nick: "Морозилочка", Item: "Mirror", Transaction: Sale, Price: 11.11, Currency: "Chaos Orb", IsOpen: true},
		{ID: 2, Nick: "일병강호남_시련", Item: "Orb", Transaction: Bay, Price: 12.23, Currency: "Divine Orb", IsOpen: true},
		{ID: 3, Nick: "Nick", Transaction: Bay, Item: "Vaal Orb", Price: 42.23, Currency: "Chaos Orb", IsOpen: true},
		{ID: 4, Nick: "<HKoL> Snifferrrrr", Item: "Chromatic Orb", Transaction: Sale, Price: 1, Currency: "Mirror of Kalandra", IsOpen: true},
	}
)

func init() {

}

func TradeWindow(isOpen *bool) {
	if config.App.Setting.IsOpen {
		misc.SelectableSettingItem("Trade", pageTrade)
	}
	style := imgui.CurrentIO().Ctx().Style()
	if *isOpen {
		if config.App.Trade.IsPinned {
			imgui.SetNextWindowSize(config.App.Trade.WindowSize)
			imgui.BeginV(fmt.Sprintf("%s (%v trade offers)###Trade", config.App.Trade.Title, len(cards)), &config.App.Trade.IsOpen, config.App.Trade.WindowFlags)

			for _, t := range cards {
				tradeCard(&t)
			}

			imgui.End()
		} else {
			imgui.SetNextWindowSize(config.App.Trade.WindowSize)
			imgui.BeginV(config.App.Trade.Title, &config.App.Trade.IsOpen, config.App.Trade.WindowFlags)

			titleSize := imgui.CalcTextSize(config.App.Trade.Title)
			imgui.Dummy(imgui.Vec2{X: (imgui.ContentRegionAvail().X-titleSize.X)/2 - style.ItemSpacing().X, Y: 19})
			imgui.SameLine()
			imgui.Text(config.App.Trade.Title)
			imgui.Text(fmt.Sprintf("%v", config.App.Trade.IsPinned))

			imgui.End()
		}
	}
}

type Trade struct {
	ID uint

	Nick        string
	Item        string
	Transaction Transaction
	Price       float32
	Currency    string

	IsOpen bool
	//AdditionTime time.Time
}

type Transaction bool

const (
	Sale Transaction = true
	Bay  Transaction = false
)

func TransactionArrow(t Transaction) string {
	if t {
		return fonts.IconsFontAwesome6.Icons["AnglesLeft"]
	} else {
		return fonts.IconsFontAwesome6.Icons["AnglesRight"]
	}
}

func init() {
	misc.Log().Debug("init trade")
	maps()
	misc.Log().Debug("register trade pages")
}

func tradeCard(trd *Trade) {
	if imgui.CollapsingHeaderBoolPtrV(fmt.Sprintf("%s  %s  %v %s##%v",
		trd.Nick,
		TransactionArrow(trd.Transaction),
		trd.Price,
		trd.Currency,
		trd.ID),
		&trd.IsOpen, 0) {
		imgui.BeginChildStrV(fmt.Sprintf("%v", trd.ID), imgui.Vec2{X: 0, Y: 58}, true, 0)

		//line 1
		imgui.Text(trd.Item)
		imgui.Spacing()

		//line 2
		if trd.Transaction == Sale {
			imgui.Button("1m")
			imgui.SameLine()
			imgui.Button("In {locate name}")
			imgui.SameLine()
			imgui.Button("Sold")
		} else {
			imgui.Button("Thanks")
		}

		imgui.EndChild()
	}
}

var isTFTVBlacklist bool

func pageTrade() {
	imgui.SeparatorText("TRADE")
	// https://raw.githubusercontent.com/The-Forbidden-Trove/character_name_blacklist/main/blacklist.txt
	imgui.Checkbox("Use TFT blacklist", &isTFTVBlacklist)
	if imgui.TreeNodeStr("Button bay") {
		imgui.Text("Thanks")
		imgui.TreePop()
	}
	if imgui.TreeNodeStr("Button sell") {
		imgui.Text("1m")
		imgui.Text("In {locate name}")
		imgui.Text("Sold")
		imgui.TreePop()
	}
}
