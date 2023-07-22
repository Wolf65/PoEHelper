package fonts

import (
	"fmt"
	"unsafe"

	imgui "github.com/AllenDang/cimgui-go"
)

var (
	RusRanges       = []uint16{0x0020, 0xA69F, 0}
	IconRange       = []uint16{uint16(IconsFontAwesome6.Min), uint16(IconsFontAwesome6.Max16), 0}
	IconBrandsRange = []uint16{uint16(IconsFontAwesome6Brands.Min), uint16(IconsFontAwesome6Brands.Max16), 0}

	baseFontSize float32 = 13
	iconFontSize float32 = 13
)

type Font struct {
	// The filenames of the associated TTF files are provided in Filenames,
	// where each entry stores first an abbreviated name for the fot and
	// then the actual filename.
	Filenames [][2]string
	// The range of Unicode code points is given by [Min, Max). The largest
	// 16-bit code point is stored in Max16.
	Min, Max16, Max int
	// Icons stores the mapping from user-friendly names to code points.
	Icons map[string]string
}

func AppendDefaultFont(io imgui.IO) {
	ioF := imgui.CurrentIO()

	fontFa := imgui.NewFontConfig()
	fontFa.SetMergeMode(true)
	fontFa.SetPixelSnapH(true)

	ioF.Fonts().AddFontFromFileTTFV("fonts/ttf/JetBrainsMono-Medium.ttf", baseFontSize, 0, (*imgui.Wchar)(unsafe.Pointer(&RusRanges[0])))
	// ioF.Fonts().AddFontFromFileTTFV(fmt.Sprintf("fonts/ttf/%s", IconsFontAwesome6.Filenames[0][1]), iconFontSize, fontFa, (*imgui.Wchar)(unsafe.Pointer(&IconRange[0])))
	ioF.Fonts().AddFontFromFileTTFV(fmt.Sprintf("fonts/ttf/%s", IconsFontAwesome6.Filenames[1][1]), iconFontSize, fontFa, (*imgui.Wchar)(unsafe.Pointer(&IconRange[0])))
	ioF.Fonts().AddFontFromFileTTFV(fmt.Sprintf("fonts/ttf/%s", IconsFontAwesome6Brands.Filenames[0][1]), iconFontSize, fontFa, (*imgui.Wchar)(unsafe.Pointer(&IconBrandsRange[0])))

	ioF.Fonts().Build()
}
