package window

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"poehelper/config"
	"poehelper/fonts"
	"poehelper/misc"
	"runtime"

	imgui "github.com/AllenDang/cimgui-go"
)

var (
	content         string
	leaguesName     = []string{"Standard", "Hardcore", "Temp Standard", "Temp Hardcore"}
	currentLeagueID = 0
)

func init() {
	misc.Log().Debug("init settings")
	misc.Log().Debug("register setting pages")
}

func SettingsWindow(isOpen *bool) {
	style := imgui.CurrentIO().Ctx().Style()
	if *isOpen {
		imgui.SetNextWindowSize(config.App.Setting.WindowSize)
		imgui.BeginV(config.App.Setting.Title, &config.App.Setting.IsOpen, config.App.Setting.WindowFlags)
		heightChild := imgui.ContentRegionAvail().Y - config.App.Vars.BaseButton.Y - style.ItemSpacing().Y

		imgui.BeginChildStrV("IDSettingMenu", imgui.NewVec2(150, heightChild), true, 0)
		imgui.PushStyleVarVec2(imgui.StyleVarWindowPadding, imgui.Vec2{X: 0, Y: 0})

		imgui.BeginChildStrV("add", imgui.NewVec2(0, -(imgui.FontSize()+style.ItemSpacing().Y*2)), false, 0)
		imgui.EndChild()
		imgui.BeginChildStrV("last", imgui.NewVec2(0, 0), false, 0)
		imgui.EndChild()

		imgui.PopStyleVar()
		imgui.EndChild()
		imgui.SameLine()

		imgui.BeginChildStrV("IDSettingContent", imgui.Vec2{X: 0, Y: heightChild}, true, 0)
		imgui.EndChild()

		imgui.AlignTextToFramePadding()
		imgui.Text(fmt.Sprintf("Version: %s", config.App.Info.ProjectVersion))
		imgui.SameLine()
		misc.HelpMarker("DEV")
		imgui.SameLine()
		imgui.Dummy(imgui.Vec2{X: imgui.ContentRegionAvail().X - config.App.Vars.BaseButton.X*2 - style.ItemSpacing().X*2, Y: 0})
		imgui.SameLine()
		imgui.ButtonV("Save", config.App.Vars.BaseButton)
		imgui.SameLine()
		if imgui.ButtonV("Cancel", config.App.Vars.BaseButton) {
			config.App.Setting.IsOpen = false
		}

		imgui.End()
	}
	if config.App.Setting.IsOpen {
		misc.SelectableSettingItem("General", pageGeneral)
		misc.SelectableSettingItem("About", pageAbout)
	}
}

func callback(data imgui.InputTextCallbackData) int {
	misc.Log().WithFields(logrus.Fields{
		"data": data.Buf(),
	}).Debug("Input callback")
	return 0
}

func pageGeneral() {
	imgui.SeparatorText("GENERAL")
	imgui.AlignTextToFramePadding()
	imgui.Text("Path to Client.txt")
	imgui.SameLine()
	imgui.PushItemWidth(imgui.ContentRegionAvail().X)
	imgui.InputTextWithHint("##ClientlogPath", "C:\\Program Files (x86)\\Steam\\steamapps\\common\\Path of Exile\\logs\\Client.txt", &content, imgui.InputTextFlagsAllowTabInput, callback)
	//imgui.SameLine()
	//imgui.Button("Change")

	imgui.AlignTextToFramePadding()
	imgui.Text("League")
	imgui.SameLine()
	imgui.PushItemWidth(imgui.ContentRegionAvail().X)
	if imgui.BeginComboV("##LeaguesCombo", leaguesName[currentLeagueID], 0) {
		for i, _ := range leaguesName {
			isSelect := currentLeagueID == i
			if imgui.SelectableBoolPtr(leaguesName[i], &isSelect) {
				currentLeagueID = i
			}
			if isSelect {
				imgui.SetItemDefaultFocus()
			}
		}
		imgui.EndCombo()
	}
	if imgui.Button("File browser") {
		config.App.FileBrowser.IsOpen = true
	}
}

func pageAbout() {
	imgui.SeparatorText("ABOUT")
	imgui.Text(config.App.Info.ProjectName)
	imgui.Spacing()
	imgui.Text("Used:")
	imgui.BulletText(fmt.Sprintf("Golang: %s", runtime.Version()))
	imgui.BulletText(fmt.Sprintf("UI: Cimgui-go (Dear ImGui %s)", imgui.Version()))
	imgui.BulletText("Icons: FontAwesome6")
	imgui.Spacing()
	imgui.AlignTextToFramePadding()
	imgui.Text("History: ")
	imgui.SameLine()
	if imgui.ButtonV("Changelog", config.App.Vars.BaseLongButton) {
		config.App.Changelog.IsOpen = true
	}
	imgui.Spacing()
	imgui.Text("Social links:")
	imgui.ButtonV(fonts.IconsFontAwesome6Brands.Icons["Github"], imgui.NewVec2(45, 45))
	imgui.SameLine()
	imgui.ButtonV(fonts.IconsFontAwesome6Brands.Icons["Telegram"], imgui.NewVec2(45, 45))
}
