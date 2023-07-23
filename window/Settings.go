package window

import (
	"fmt"
	"poehelper/config"
	"poehelper/fonts"
	"poehelper/misc"

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
	misc.Log().Debug("init settings")
	registrationSettingsPages(settingPage{"General", "pageGeneral", pageGeneral})
	registrationSettingsPages(settingPage{"About", "pageAbout", pageAbout})
	misc.Log().Debug("register setting pages")
}

func SettingsWindow(isOpen bool) {
	if isOpen {
		imgui.SetNextWindowSize(config.App.Setting.WindowSize)
		imgui.BeginV(config.App.Setting.Title, &config.App.Setting.IsOpen, config.App.Setting.WindowFlags)
		heightChild := imgui.ContentRegionAvail().Y - config.App.Vars.BaseButton.Y - config.App.Vars.ItemSpacing.Y
		imgui.BeginChildStrV("settingMenu", imgui.NewVec2(150, heightChild), true, 0)

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
		imgui.ButtonV(fonts.IconsFontAwesome6.Icons["Rotate"], imgui.Vec2{X: config.App.Vars.BaseButton.Y, Y: config.App.Vars.BaseButton.Y})
		imgui.SameLine()
		imgui.Dummy(imgui.Vec2{X: imgui.ContentRegionAvail().X - config.App.Vars.BaseButton.X*2 - config.App.Vars.ItemSpacing.X*2, Y: 0})
		imgui.SameLine()
		imgui.ButtonV("Save", config.App.Vars.BaseButton)
		imgui.SameLine()
		if imgui.ButtonV("Cancel", config.App.Vars.BaseButton) {
			config.App.Setting.IsOpen = false
		}

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
	imgui.Text(config.App.Info.ProjectName)
	imgui.Spacing()
	imgui.Text("Used:")
	imgui.BulletText(fmt.Sprintf("Cimgui-go (Dear ImGui %s)", imgui.Version()))
	imgui.BulletText("Icons FontAwesome6")
	imgui.Spacing()
	imgui.AlignTextToFramePadding()
	imgui.Text("History: ")
	imgui.SameLine()
	if imgui.ButtonV("Changelog", config.App.Vars.BaseLongButton) {
		config.App.Changelog.IsOpen = true
	}
	imgui.Spacing()
	imgui.Text("Social links:")
	imgui.ButtonV(fonts.IconsFontAwesome6Brands.Icons["Github"], imgui.NewVec2(32, 32))
	imgui.SameLine()
	imgui.ButtonV(fonts.IconsFontAwesome6Brands.Icons["Telegram"], imgui.NewVec2(32, 32))
}
