package fonts

import (
	imgui "github.com/AllenDang/cimgui-go"
)

var (
	W = []imgui.Wchar{0x0020, 0x00FF, 0}

// W1 = []imgui.Wchar{0xf081, 0xf949, 0}
)

// type font struct {
// 	// The filenames of the associated TTF files are provided in Filenames,
// 	// where each entry stores first an abbreviated name for the fot and
// 	// then the actual filename.
// 	Filenames [][2]string
// 	// The range of Unicode code points is given by [Min, Max). The largest
// 	// 16-bit code point is stored in Max16.
// 	Min, Max16, Max int
// 	// Icons stores the mapping from user-friendly names to code points.
// 	Icons map[string]string
// }

func AppendDefaultFont() {
	// io := imgui.CurrentIO()

	// io.Fonts().AddFontFromFileTTF("fonts/JetBrainsMono-Medium.ttf", 13)
	// io.Fonts().Build()
	// f.Fonts().Build()
	// f.Fonts().AddFontFromFileTTFV("fonts/JetBrainsMono-Medium.ttf", 13, 0, f.Fonts().GlyphRangesDefault())
	// f.Fonts().Build()

	// f.Fonts().Build()
	// f.AddFontFromFileTTFV("fonts/JetBrainsMono-Medium.ttf", 13, 0, f.GlyphRangesDefault())

	// f.Build()

	// ioF := imgui.CurrentIO()@

	// fontFa := imgui.NewFontConfig()
	// fontFa.SetMergeMode(true)
	// fontFa.SetPixelSnapH(true)

	// ranges := imgui.NewGlyphRange()
	// builder := imgui.NewFontGlyphRangesBuilder()

	// builder.AddRanges((*imgui.Wchar)(unsafe.Pointer(&W)))
	// builder.BuildRanges(ranges)

	// rangesFa := imgui.NewGlyphRange()
	// builderFa := imgui.NewFontGlyphRangesBuilder()

	// builderFa.AddRanges((*imgui.Wchar)(unsafe.Pointer(&W1[0])))
	// builderFa.BuildRanges(rangesFa)

	// ioF.Fonts().AddFontFromFileTTFV("fonts/JetBrainsMono-Medium.ttf", 13, 0, ioF.Fonts().GlyphRangesDefault())
	// ioF.Fonts().AddFontFromFileTTFV("fonts/JetBrainsMono-Medium.ttf", 13, fontFa, ranges.Data())
	// ioF.Fonts().AddFontFromFileTTFV("fonts/fa-brands-400.ttf", 13, fontFa, rangesFa.Data())

	// ioF.Fonts().Build()
}
