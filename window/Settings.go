package window

import (
	"fmt"
	"poehelper/config"

	"golang.org/x/exp/slices"

	imgui "github.com/AllenDang/cimgui-go"
)

var (
	inputPath string

	settingPages []settingPage
)

type settingPage struct {
	label   string
	alias   string
	content func()
}

func registrationSettingsPages(sP settingPage) {
	switch sP.alias {
	case "pageGeneral":
		settingPages = slices.Insert(settingPages, 0, sP)
	case "pageAbout":
		settingPages = slices.Insert(settingPages, len(settingPages), sP)
	default:
		if len(settingPages) == 0 {
			settingPages = append(settingPages, sP)
		} else {
			settingPages = slices.Insert(settingPages, len(settingPages)-1, sP)
		}
	}
}

func init() {
	registrationSettingsPages(settingPage{"General", "pageGeneral", pageGeneral})
	registrationSettingsPages(settingPage{"About", "pageAbout", pageAbout})
}

func SettingsWindow(isOpen bool) {
	if isOpen {
		imgui.SetNextWindowSize(config.App.Setting.WindowSize)
		imgui.BeginV(config.App.Setting.Title, &config.App.Setting.IsOpen, config.App.Setting.WindowFlags)
		heightChild := imgui.ContentRegionAvail().Y - config.App.Vars.BaseButton.Y - config.App.Vars.ItemSpacing.Y
		imgui.BeginChildStrV("settingMenu", imgui.NewVec2(150, heightChild), true, 0)
		imgui.Text(config.App.Setting.SelectPage)

		for _, page := range settingPages {
			if imgui.ButtonV(page.label, imgui.NewVec2(imgui.ContentRegionAvail().X, 0)) {
				config.App.Setting.SelectPage = page.alias
			}
		}

		imgui.EndChild()
		imgui.SameLine()
		imgui.BeginChildStrV("settingContent", imgui.NewVec2(0, heightChild), true, 0)

		for i, page := range settingPages {
			if page.alias == config.App.Setting.SelectPage {
				settingPages[i].content()
			}
		}

		imgui.EndChild()

		imgui.AlignTextToFramePadding()
		imgui.Text(fmt.Sprintf("Ver: %s", config.App.Info.ProjectVersion))
		imgui.SameLine()
		// informate new update
		imgui.ButtonV("UP", imgui.Vec2{X: config.App.Vars.BaseButton.Y, Y: config.App.Vars.BaseButton.Y})
		imgui.SameLine()
		imgui.Dummy(imgui.Vec2{X: imgui.ContentRegionAvail().X - config.App.Vars.BaseButton.X*2 - config.App.Vars.ItemSpacing.X*2, Y: 0})
		imgui.SameLine()
		imgui.ButtonV("Save", config.App.Vars.BaseButton)
		imgui.SameLine()
		imgui.ButtonV("Cancel", config.App.Vars.BaseButton)

		imgui.End()
	}
}

func pageGeneral() {
	imgui.AlignTextToFramePadding()
	imgui.Text("Path to PoE logs")
	imgui.SameLine()
	imgui.PushItemWidth(imgui.ContentRegionAvail().X)
	imgui.InputTextWithHint("##inputTextLogPath", "c:/..", &inputPath, imgui.InputTextFlagsAllowTabInput, nil)
}

func pageAbout() {
	imgui.Text("PoE Helper")
	imgui.Spacing()
	imgui.Text(fmt.Sprintf("Used cingui-go (Dear ImGui %s)", imgui.Version()))
	imgui.AlignTextToFramePadding()
	imgui.Text("History: ")
	imgui.SameLine()
	if imgui.ButtonV("Changelog", config.App.Vars.BaseLongButton) {
		config.App.Changelog.IsOpen = true
	}
	imgui.Text("Social links:")
	imgui.ButtonV("GitHub", imgui.NewVec2(32, 32))
}
