/*
This file is part of convert3a.

convert3a is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

convert3a is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with convert3a.  If not, see <https://www.gnu.org/licenses/>.
*/
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/DomesticMoth/go3a"
	"github.com/golang/freetype/truetype"
	"github.com/llgcode/draw2d"
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/font/gofont/goitalic"
	"golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/font/gofont/goregular"
)

func main() {
	opts := GetOpts()
	if opts.Inp == "" {
		fmt.Println("There is no input file")
		return
	}
	if opts.Out == "" {
		inp := filepath.Base(opts.Inp)
		inp = strings.TrimSuffix(inp, ".3a")
		if opts.Dyn {
			opts.Out = inp + ".gif"
		} else {
			opts.Out = inp + ".png"
		}
	}
	var font_size float64 = 72
	if opts.FontSizeSetted {
		font_size = float64(opts.FontSize)
	}
	DPI := 72
	if opts.DPISetted {
		DPI = opts.DPI
	}
	// Set font
	fontCache := MyFontCache{}
	TTFs := map[string]([]byte){
		"goregular": goregular.TTF,
		"gobold":    gobold.TTF,
		"goitalic":  goitalic.TTF,
		"gomono":    gomono.TTF,
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
	if !opts.Dyn {
		RenderToPng(f, float64(art.Header.Width), float64(art.Header.Height), art.Body[preview], font_info, palette)
	} else {
		RenderToGif(f, int(art.Header.Width), int(art.Header.Height), art.Body, int(art.Header.Delay), palette, font_info)
	}
}
