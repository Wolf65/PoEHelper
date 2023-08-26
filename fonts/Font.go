package fonts

import (
	"fmt"
	imgui "github.com/AllenDang/cimgui-go"
)

var (
	IconRange       = []imgui.Wchar{IconsFontAwesome6.Min, IconsFontAwesome6.Max, 0}
	IconBrandsRange = []imgui.Wchar{IconsFontAwesome6Brands.Min, IconsFontAwesome6Brands.Max, 0}

	baseFontSize              float32 = 13
	fontAwesomeFontSize       float32 = baseFontSize * 4 / 3
	fontAwesomeBrandsFontSize float32 = baseFontSize * 6 / 3
)

type Font struct {
	// The filenames of the associated TTF files are provided in Filenames,
	// where each entry stores first an abbreviated name for the fot and
	// then the actual filename.
	Filenames [][2]string
	// The range of Unicode code points is given by [Min, Max). The largest
	// 16-bit code point is stored in Max16.
	Min, Max16, Max imgui.Wchar
	// Icons stores the mapping from user-friendly names to code points.
	Icons map[string]string
}

func AppendDefaultFont(io imgui.IO) {
	// Base Font

	baseConfig := imgui.NewFontConfig()
	baseConfig.SetPixelSnapH(true)

	baseRange := imgui.NewGlyphRange()
	baseBuilder := imgui.NewFontGlyphRangesBuilder()
	baseBuilder.AddRanges(io.Fonts().GlyphRangesCyrillic())
	//jetBrainsRangeBuilder.AddRanges(io.Fonts().GlyphRangesKorean())
	//jetBrainsRangeBuilder.AddRanges(io.Fonts().GlyphRangesJapanese())
	//jetBrainsRangeBuilder.AddRanges(io.Fonts().GlyphRangesChineseFull())
	//jetBrainsRangeBuilder.AddRanges(io.Fonts().GlyphRangesGreek())
	//jetBrainsRangeBuilder.AddRanges(io.Fonts().GlyphRangesThai())
	//jetBrainsRangeBuilder.AddRanges(io.Fonts().GlyphRangesVietnamese())
	baseBuilder.BuildRanges(baseRange)

	io.Fonts().AddFontFromFileTTFV("fonts/ttf/JetBrainsMono-Medium.ttf",
		baseFontSize,
		baseConfig,
		baseRange.Data())

	// FontAwesome
	fontAwesomeConfig := imgui.NewFontConfig()
	fontAwesomeConfig.SetMergeMode(true)
	fontAwesomeConfig.SetGlyphMinAdvanceX(fontAwesomeFontSize)
	fontAwesomeConfig.SetGlyphMaxAdvanceX(fontAwesomeFontSize)
	fontAwesomeConfig.SetPixelSnapH(true)
	fontAwesomeConfig.SetGlyphOffset(imgui.Vec2{X: 0, Y: 2})

	io.Fonts().AddFontFromFileTTFV(fmt.Sprintf("fonts/ttf/%s",
		IconsFontAwesome6.Filenames[1][1]),
		fontAwesomeFontSize,
		fontAwesomeConfig,
		imgui.SliceToPtr(IconRange))

	// FontAwesome Brands
	fontAwesomeBrandsConfig := imgui.NewFontConfig()
	fontAwesomeBrandsConfig.SetMergeMode(true)
	fontAwesomeBrandsConfig.SetGlyphMinAdvanceX(fontAwesomeBrandsFontSize)
	fontAwesomeBrandsConfig.SetGlyphMaxAdvanceX(fontAwesomeBrandsFontSize)
	fontAwesomeBrandsConfig.SetPixelSnapH(true)
	fontAwesomeBrandsConfig.SetGlyphOffset(imgui.Vec2{X: 1, Y: 6})

	io.Fonts().AddFontFromFileTTFV(fmt.Sprintf("fonts/ttf/%s",
		IconsFontAwesome6Brands.Filenames[0][1]),
		fontAwesomeBrandsFontSize,
		fontAwesomeBrandsConfig,
		imgui.SliceToPtr(IconBrandsRange))

	io.Fonts().Build()
}
