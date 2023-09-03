package window

import (
	"os"
	"poehelper/config"

	imgui "github.com/AllenDang/cimgui-go"
)

var (
	text, _ = os.ReadFile(config.App.Changelog.ChangelogPath)

	changelogText = string(text)
)

func Changelog(isOpen *bool) {
	if *isOpen {
		imgui.SetNextWindowSize(config.App.Changelog.WindowSize)
		imgui.BeginV(config.App.Changelog.Title, &config.App.Changelog.IsOpen, config.App.Changelog.WindowFlags)
		imgui.InputTextMultiline("##changelogMultilineText", &changelogText, imgui.ContentRegionAvail(), imgui.InputTextFlagsReadOnly, nil)
		imgui.End()
	}
}
