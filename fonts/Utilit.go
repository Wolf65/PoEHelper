package fonts

import (
	"unsafe"

	imgui "github.com/AllenDang/cimgui-go"
)

var (
	W = []imgui.Wchar{0x0020, 0x00FF, 0}
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

func AppendDefaultFont(io imgui.IO) {
	ioF := imgui.CurrentIO()

	fontFa := imgui.NewFontConfig()
	fontFa.SetMergeMode(true)
	fontFa.SetPixelSnapH(true)

	ranges := imgui.NewGlyphRange()
	builder := imgui.NewFontGlyphRangesBuilder()

	builder.AddRanges((*imgui.Wchar)(unsafe.Pointer(&W)))
	builder.BuildRanges(ranges)

	ioF.Fonts().AddFontFromFileTTFV("fonts/JetBrainsMono-Medium.ttf", 13, 0, ranges.Data())

	ioF.Fonts().Build()
}
