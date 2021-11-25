package main

import (
	"github.com/DomesticMoth/go3a"
	"os"
	"fmt"
	"strings"
	"path/filepath"
	"github.com/llgcode/draw2d"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/font/gofont/goitalic"
	"golang.org/x/image/font/gofont/gomono"
	"github.com/golang/freetype/truetype"
)

func main() {
	opts := GetOpts()
	if opts.Inp == "" {
		fmt.Println("There is no input file")
		return
	}
	if opts.Out == "" {
		inp := filepath.Base(opts.Inp)
		if strings.HasSuffix(inp, ".3a") {
			inp = strings.TrimSuffix(inp, ".3a")
		}
		if opts.Dyn {
			opts.Out = inp + ".gif"
		}else {
			opts.Out = inp + ".png"
		}
	}
	var font_size float64 = 72
	if opts.FontSizeSetted{
		font_size = float64(opts.FontSize)
	}
	DPI := 72
	if opts.DPISetted{
		DPI = opts.DPI
	}
	// Set font
	fontCache := MyFontCache{}
	TTFs := map[string]([]byte){
		"goregular": goregular.TTF,
		"gobold": gobold.TTF,
		"goitalic": goitalic.TTF,
		"gomono": gomono.TTF,
	}

	for fontName, TTF := range TTFs {
		font, err := truetype.Parse(TTF)
		if err != nil {
			panic(err)
		}
		fontCache.Store(draw2d.FontData{Name: fontName}, font)
	}
	draw2d.SetFontCache(fontCache)
	font_data := draw2d.FontData{
		Name:   "gomono",
		Family: draw2d.FontFamilyMono,
		Style:  draw2d.FontStyleBold,
	}
	// Set color palette
	palette := GeneratePalette(opts.Colors)
	art, err := go3a.LoadFile(opts.Inp)
	if err != nil {
		panic(err)
	}
	preview := int(art.Header.Preview)
	if opts.PreviewSetted {
		preview = opts.Preview
	}
	font_info := FontInfo{font_data, font_size, DPI}
	f, ferr := os.OpenFile(opts.Out, os.O_WRONLY|os.O_CREATE, 0600)
	if ferr != nil {
		panic(ferr)
	}
	defer f.Close()
	if !opts.Dyn{
		RenderToPng(f, float64(art.Header.Width), float64(art.Header.Height), art.Body[preview], font_info, palette)
	}else{
		RenderToGif(f, int(art.Header.Width), int(art.Header.Height), art.Body, int(art.Header.Delay), palette, font_info)
	}
}
