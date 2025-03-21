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
	"github.com/DomesticMoth/go3a"
	"image/color"
)

func DefaultPalette() []color.Color {
	palette := []color.Color{
		color.RGBA{0x00, 0x00, 0x00, 0xff}, //Black
		color.RGBA{0xff, 0x00, 0x00, 0xff}, //Light red
		color.RGBA{0x00, 0xff, 0x00, 0xff}, //Light green
		color.RGBA{0xff, 0xff, 0x00, 0xff}, //Yellow
		color.RGBA{0x00, 0x00, 0xff, 0xff}, //Light blue
		color.RGBA{0xff, 0x00, 0xff, 0xff}, //Light magenta
		color.RGBA{0x00, 0xff, 0xff, 0xff}, //Light cyan
		color.RGBA{0xff, 0xff, 0xff, 0xff}, //Light white
		color.RGBA{0x80, 0x80, 0x80, 0xff}, //Gray
		color.RGBA{0x80, 0x00, 0x00, 0xff}, //Red
		color.RGBA{0x00, 0x80, 0x00, 0xff}, //Green
		color.RGBA{0x80, 0x80, 0x00, 0xff}, //Brown
		color.RGBA{0x00, 0x00, 0x80, 0xff}, //Blue
		color.RGBA{0x80, 0x00, 0x80, 0xff}, //Magenta
		color.RGBA{0x00, 0x80, 0x80, 0xff}, //Cyan
		color.RGBA{0xC0, 0xC0, 0xC0, 0xff}, //White
		color.RGBA{0x00, 0x00, 0x00, 0xff}, //Default bg
		color.RGBA{0xff, 0xff, 0xff, 0xff}, //Default fg
	}
	return palette
}

func GeneratePalette(opts []*string) []color.Color {
	palette := DefaultPalette()
	for i, c := range opts {
		if *c == "" {
			continue
		}
		color, err := ParseHexColor(*c)
		if err == nil {
			palette[i] = color
		} else {
			panic(err)
		}
	}
	return palette
}

func Color3atoRGB(color go3a.Color) int {
	switch color {
	case go3a.ColorBlack:
		return 0
	case go3a.ColorBlue:
		return 12
	case go3a.ColorGreen:
		return 10
	case go3a.ColorCyan:
		return 14
	case go3a.ColorRed:
		return 9
	case go3a.ColorMagenta:
		return 13
	case go3a.ColorYellow:
		return 11
	case go3a.ColorWhite:
		return 15
	case go3a.ColorGray:
		return 8
	case go3a.ColorBrightBlue:
		return 4
	case go3a.ColorBrightGreen:
		return 2
	case go3a.ColorBrightCyan:
		return 6
	case go3a.ColorBrightRed:
		return 1
	case go3a.ColorBrightMagenta:
		return 5
	case go3a.ColorBrightYellow:
		return 3
	case go3a.ColorBrightWhite:
		return 7
	}
	return 0
}

func ParseHexColor(s string) (c color.RGBA, err error) {
	c.A = 0xff
	switch len(s) {
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("invalid color %s length, must be 7 or 4\n", s)

	}
	return
}
