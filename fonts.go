package main

import (
	"github.com/golang/freetype/truetype"
	"github.com/llgcode/draw2d"
	"fmt"
)

type MyFontCache map[string]*truetype.Font

func (fc MyFontCache) Store(fd draw2d.FontData, font *truetype.Font) {
	fc[fd.Name] = font
}

func (fc MyFontCache) Load(fd draw2d.FontData) (*truetype.Font, error) {
	font, stored := fc[fd.Name]
	if !stored {
		return nil, fmt.Errorf("Font %s is not stored in font cache.", fd.Name)
	}
	return font, nil
}
