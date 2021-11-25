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
	"flag"
	"os"
)

type Opts struct{
	Dyn bool
	Preview int
	PreviewSetted bool
	FontSize int
	FontSizeSetted bool
	DPI int
	DPISetted bool
	Inp string
	Out string
	Colors []*string
}

func PrintHelp(){
	fmt.Println("convert3a by DomesticMoth")
	fmt.Println("Tool for convert 3a animations to media formats\n")
	fmt.Println("Args:")
	fmt.Println("	--in [file]			Path to 3a file (required)")
	fmt.Println("	--out [file]			Path to output media file")
	fmt.Println("	--dyn					Render animation instead of static image (flag)")
	fmt.Println("	--preview [nomber]")
	fmt.Println("	--fontsize [nomber]\n")
	fmt.Println("You can change the palette by specifying values for the following parameters in hex format")
	fmt.Println("	--ColorBlack")
	fmt.Println("	--ColorLightRed")
	fmt.Println("	--ColorLightGreen")
	fmt.Println("	--ColorYellow")
	fmt.Println("	--ColorLightBlue")
	fmt.Println("	--ColorLightMagenta")
	fmt.Println("	--ColorLightCyan")
	fmt.Println("	--ColorLightWhite")
	fmt.Println("	--ColorGray")
	fmt.Println("	--ColorRed")
	fmt.Println("	--ColorGreen")
	fmt.Println("	--ColorBrown")
	fmt.Println("	--ColorBlue")
	fmt.Println("	--ColorMagenta")
	fmt.Println("	--ColorCyan")
	fmt.Println("	--ColorWhite")
	fmt.Println("	--ColorDefaultBg")
	fmt.Println("	--ColorDefaultFg")
	fmt.Println("Example '--ColorRed #ff0000'")
}

func GetOpts() Opts {
	CallHelp := flag.Bool("help", false, "")
	Dyn := flag.Bool("dyn", false, "")
	Preview := flag.Int("preview", 2147483647, "")
	FontSize := flag.Int("fontsize", 2147483647, "")
	DPI := flag.Int("dpi", 2147483647, "")
	Inp := flag.String("in", "", "")
	Out := flag.String("out", "", "")
	colors := []*string{
		flag.String("ColorBlack", "", ""),
		flag.String("ColorLightRed", "", ""),
		flag.String("ColorLightGreen", "", ""),
		flag.String("ColorYellow", "", ""),
		flag.String("ColorLightBlue", "", ""),
		flag.String("ColorLightMagenta", "", ""),
		flag.String("ColorLightCyan", "", ""),
		flag.String("ColorLightWhite", "", ""),
		flag.String("ColorGray", "", ""),
		flag.String("ColorRed", "", ""),
		flag.String("ColorGreen", "", ""),
		flag.String("ColorBrown", "", ""),
		flag.String("ColorBlue", "", ""),
		flag.String("ColorMagenta", "", ""),
		flag.String("ColorCyan", "", ""),
		flag.String("ColorWhite", "", ""),
		flag.String("ColorDefaultBg", "", ""),
		flag.String("ColorDefaultFg", "", ""),
	}
	flag.Parse()
	PreviewSetted := true
	if *Preview >= 2147483647 {
		PreviewSetted = false
	}
	FontSizeSetted := true
	if *FontSize >= 2147483647 {
		FontSizeSetted = false
	}
	DPISetted := true
	if *DPI >= 2147483647 {
		DPISetted = false
	}
	if *CallHelp{
		PrintHelp()
		os.Exit(0)
	}
	return Opts{*Dyn, *Preview, PreviewSetted, *FontSize, FontSizeSetted, *DPI, DPISetted,*Inp, *Out, colors}
}
