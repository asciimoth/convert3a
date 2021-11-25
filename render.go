package main

import (
	"math"
	"image"
	"image/png"
	"image/gif"
	"github.com/llgcode/draw2d/draw2dimg"
	"io"
	"github.com/llgcode/draw2d"
	"github.com/DomesticMoth/go3a"
	"image/color"
)

type FontInfo struct{
	 font_data draw2d.FontData
	 font_size float64
	 DPI int
}

func getBoundsOfFont(font_data draw2d.FontData, font_size float64, DPI int) (float64, float64){
	dest := image.NewRGBA(image.Rect(0, 0, 100, 100))
	gc := draw2dimg.NewGraphicContext(dest)
	gc.SetFontData(font_data)
	gc.SetFontSize(font_size)
	gc.SetDPI(DPI)
	width := 0.0
	height := 0.0
	letters := []string{
		"A", "B", "C", "D", "E", "F", 
		"G", "H", "I", "J", "K", "L", 
		"M", "N", "O", "P", "Q", "R", 
		"S", "T", "U", "V", "W", "X", 
		"Y", "Z", "a", "b", "c", "d", 
		"e", "f", "g", "h", "i", "j", 
		"k", "l", "m", "n", "o", "p", 
		"q", "r", "s", "t", "u", "v", 
		"w", "x", "y", "x", "1", "2", 
		"3", "4", "5", "6", "7", "8", 
		"9", "0", "`", "'", "\"", ".", 
		",", "/", "\\", "|", ":", ";", 
		"*", "&", "!", "@", "#", "$", 
		"%", "~", "*", "(", ")", "[", 
		"]", "{", "}",
		}
	for _, letter := range letters{
		left, top, right, bottom := gc.GetStringBounds(letter)
		h := math.Abs(bottom-top)
		w := math.Abs(left-right)
		width = math.Max(width, w)
		height = math.Max(height, h)
	}
	return width, height
}

func DrawFullblock(gc draw2d.GraphicContext, left_x, right_x, up_y, down_y, middle_x, middle_y float64){
	gc.MoveTo(left_x, up_y)
	gc.BeginPath()
	gc.LineTo(right_x, up_y)
	gc.LineTo(right_x, down_y)
	gc.LineTo(left_x, down_y)
	gc.LineTo(left_x, up_y)
	gc.Close()
	gc.FillStroke()
}

func DrawUpperHalfBlock(gc draw2d.GraphicContext, left_x, right_x, up_y, down_y, middle_x, middle_y float64){
	gc.MoveTo(left_x, up_y)
	gc.BeginPath()
	gc.LineTo(right_x, up_y)
	gc.LineTo(right_x, middle_y)
	gc.LineTo(left_x, middle_y)
	gc.LineTo(left_x, up_y)
	gc.Close()
	gc.FillStroke()
}

func DrawLowerHalfBlock(gc draw2d.GraphicContext, left_x, right_x, up_y, down_y, middle_x, middle_y float64){
	gc.MoveTo(left_x, down_y)
	gc.BeginPath()
	gc.LineTo(right_x, down_y)
	gc.LineTo(right_x, middle_y)
	gc.LineTo(left_x, middle_y)
	gc.LineTo(left_x, down_y)
	gc.Close()
	gc.FillStroke()
}

func DrawQuadrantUpperLeftAndLowerLeftAndLowerRight(gc draw2d.GraphicContext, left_x, right_x, up_y, down_y, middle_x, middle_y float64){
	gc.MoveTo(left_x, down_y)
	gc.BeginPath()
	gc.LineTo(right_x, down_y)
	gc.LineTo(right_x, middle_y)
	gc.LineTo(middle_x, middle_y)
	gc.LineTo(middle_x, up_y)
	gc.LineTo(left_x, up_y)
	gc.LineTo(left_x, down_y)
	gc.Close()
	gc.FillStroke()
}

func DrawQuadrantUpperLeftAndUpperRightAndLowerLeft(gc draw2d.GraphicContext, left_x, right_x, up_y, down_y, middle_x, middle_y float64){
	gc.MoveTo(left_x, down_y)
	gc.BeginPath()
	gc.LineTo(middle_x, down_y)
	gc.LineTo(middle_x, middle_y)
	gc.LineTo(right_x, middle_y)
	gc.LineTo(right_x, up_y)
	gc.LineTo(left_x, up_y)
	gc.LineTo(left_x, down_y)
	gc.Close()
	gc.FillStroke()
}

func DrawQuadrantUpperLeftAndUpperRightAndLowerRight(gc draw2d.GraphicContext, left_x, right_x, up_y, down_y, middle_x, middle_y float64){
	gc.MoveTo(left_x, middle_y)
	gc.BeginPath()
	gc.LineTo(middle_x, middle_y)
	gc.LineTo(middle_x, down_y)
	gc.LineTo(right_x, down_y)
	gc.LineTo(right_x, up_y)
	gc.LineTo(left_x, up_y)
	gc.LineTo(left_x, middle_y)
	gc.Close()
	gc.FillStroke()
}

func DrawQuadrantUpperRightAndLowerLeftAndLowerRight(gc draw2d.GraphicContext, left_x, right_x, up_y, down_y, middle_x, middle_y float64){
	gc.MoveTo(left_x, down_y)
	gc.BeginPath()
	gc.LineTo(right_x, down_y)
	gc.LineTo(right_x, up_y)
	gc.LineTo(middle_x, up_y)
	gc.LineTo(middle_x, middle_y)
	gc.LineTo(left_x, middle_y)
	gc.LineTo(left_x, down_y)
	gc.Close()
	gc.FillStroke()
}

func DrawQuadrantLowerLeft(gc draw2d.GraphicContext, left_x, right_x, up_y, down_y, middle_x, middle_y float64){
	gc.MoveTo(left_x, down_y)
	gc.BeginPath()
	gc.LineTo(middle_x, down_y)
	gc.LineTo(middle_x, middle_y)
	gc.LineTo(left_x, middle_y)
	gc.LineTo(left_x, down_y)
	gc.Close()
	gc.FillStroke()
}

func DrawQuadrantLowerRight(gc draw2d.GraphicContext, left_x, right_x, up_y, down_y, middle_x, middle_y float64){
	gc.MoveTo(middle_x, down_y)
	gc.BeginPath()
	gc.LineTo(right_x, down_y)
	gc.LineTo(right_x, middle_y)
	gc.LineTo(middle_x, middle_y)
	gc.LineTo(middle_x, down_y)
	gc.Close()
	gc.FillStroke()
}

func DrawQuadrantUpperLeft(gc draw2d.GraphicContext, left_x, right_x, up_y, down_y, middle_x, middle_y float64){
	gc.MoveTo(left_x, middle_y)
	gc.BeginPath()
	gc.LineTo(middle_x, middle_y)
	gc.LineTo(middle_x, up_y)
	gc.LineTo(left_x, up_y)
	gc.LineTo(left_x, middle_y)
	gc.Close()
	gc.FillStroke()
}

func DrawQuadrantUpperRight(gc draw2d.GraphicContext, left_x, right_x, up_y, down_y, middle_x, middle_y float64){
	gc.MoveTo(middle_x, up_y)
	gc.BeginPath()
	gc.LineTo(right_x, up_y)
	gc.LineTo(right_x, middle_y)
	gc.LineTo(middle_x, middle_y)
	gc.LineTo(middle_x, up_y)
	gc.Close()
	gc.FillStroke()
}

func renderFrame(rect image.Rectangle, frame go3a.Frame, font_info FontInfo, palette []color.Color, char_width, char_height float64) *image.RGBA {
	img := image.NewRGBA(rect)
	gc := draw2dimg.NewGraphicContext(img)
	gc.SetFontData(font_info.font_data)
	gc.SetFontSize(font_info.font_size)
	gc.SetDPI(font_info.DPI)
	gc.SetLineWidth(1)
	y := char_height
	for _, row := range frame {
		x := 0.0
		for _, row_fragment := range row {
			color_fg := palette[17]
			color_bg := palette[16]
			if row_fragment.FgColor != go3a.NoColor {
				color_fg = palette[Color3atoRGB(row_fragment.FgColor)]
			}
			if row_fragment.BgColor != go3a.NoColor {
				color_bg = palette[Color3atoRGB(row_fragment.BgColor)]
			}
			for _, c := range row_fragment.Text {
				char := string(c)
				gc.SetFillColor(color_bg)
				gc.SetStrokeColor(color_bg)
				gc.MoveTo(x, y)
				gc.BeginPath()
				gc.LineTo(x+char_width, y)
				gc.LineTo(x+char_width, y-char_height)
				gc.LineTo(x, y-char_height)
				gc.LineTo(x, y)
				gc.Close()
				gc.FillStroke()
				gc.SetFillColor(color_fg)
				gc.SetStrokeColor(color_fg)
				d := 0.1
				left_x := math.Max(x-char_width+(char_width*(1-d)), 0)
				right_x := math.Min(x+char_width*(1+d), float64(rect.Dx()))
				down_y := math.Min(y+char_height-(char_height*(1-d)), float64(rect.Dy()))
				up_y := math.Max(y-char_height*(1+d), 0)
				middle_x := x+char_width/2
				middle_y := y-char_height/2
				if char == "█"{
					DrawFullblock(gc, left_x, right_x, up_y, down_y, middle_x, middle_y)
				}else if char == "▀"{
					DrawUpperHalfBlock(gc, left_x, right_x, up_y, down_y, middle_x, middle_y)
				}else if char == "▄"{
					DrawLowerHalfBlock(gc, left_x, right_x, up_y, down_y, middle_x, middle_y)
				}else if char == "▙"{
					DrawQuadrantUpperLeftAndLowerLeftAndLowerRight(gc, left_x, right_x, up_y, down_y, middle_x, middle_y)
				}else if char == "▛"{
					DrawQuadrantUpperLeftAndUpperRightAndLowerLeft(gc, left_x, right_x, up_y, down_y, middle_x, middle_y)
				}else if char == "▜"{
					DrawQuadrantUpperLeftAndUpperRightAndLowerRight(gc, left_x, right_x, up_y, down_y, middle_x, middle_y)
				}else if char == "▟"{
					DrawQuadrantUpperRightAndLowerLeftAndLowerRight(gc, left_x, right_x, up_y, down_y, middle_x, middle_y)
				}else if char == "▚"{
					DrawQuadrantUpperLeft(gc, left_x, right_x, up_y, down_y, middle_x, middle_y)
					DrawQuadrantLowerRight(gc, left_x, right_x, up_y, down_y, middle_x, middle_y)
				}else if char == "▞"{
					DrawQuadrantUpperRight(gc, left_x, right_x, up_y, down_y, middle_x, middle_y)
					DrawQuadrantLowerLeft(gc, left_x, right_x, up_y, down_y, middle_x, middle_y)
				}else if char == "▖"{
					DrawQuadrantLowerLeft(gc, left_x, right_x, up_y, down_y, middle_x, middle_y)
				}else if char == "▗"{
					DrawQuadrantLowerRight(gc, left_x, right_x, up_y, down_y, middle_x, middle_y)
				}else if char == "▘"{
					DrawQuadrantUpperLeft(gc, left_x, right_x, up_y, down_y, middle_x, middle_y)
				}else if char == "▝"{
					DrawQuadrantUpperRight(gc, left_x, right_x, up_y, down_y, middle_x, middle_y)
				}else{
					gc.FillStringAt(char, x, y)
				}
				x += char_width
			}
		}
		y += char_height
	}
	return img
}

func RenderToPng(out io.Writer, chars_x, chars_y float64, frame go3a.Frame, font_info FontInfo, palette []color.Color){
	char_width, char_height := getBoundsOfFont(font_info.font_data, font_info.font_size, font_info.DPI)
	rect := image.Rect(0, 0, int(char_width*chars_x), int(char_height*chars_y))
	img := renderFrame(rect, frame, font_info, palette, char_width, char_height)
	png.Encode(out, img)
}

func RGBtoPaletted(rgb *image.RGBA, plt *image.Paletted){
	for x := 0 ; x < plt.Rect.Dx() ; x++ {
		for y := 0 ; y < plt.Rect.Dy() ; y++ {
			color_rgba := color.RGBA{
				rgb.Pix[(y-rgb.Rect.Min.Y)*rgb.Stride + (x-rgb.Rect.Min.X)*4],
				rgb.Pix[(y-rgb.Rect.Min.Y)*rgb.Stride + (x-rgb.Rect.Min.X)*4+1],
				rgb.Pix[(y-rgb.Rect.Min.Y)*rgb.Stride + (x-rgb.Rect.Min.X)*4+2],
				rgb.Pix[(y-rgb.Rect.Min.Y)*rgb.Stride + (x-rgb.Rect.Min.X)*4+3],
			}
			for i, color := range plt.Palette {
				if color == color_rgba{
					plt.Pix[(y-plt.Rect.Min.Y)*plt.Stride + (x-plt.Rect.Min.X)*1] = uint8(i)
					break
				}
			}
		}
	}
}

func RenderToGif(out io.Writer, chars_x, chars_y int, frames go3a.Body, delay int, palette []color.Color, font_info FontInfo){
	char_width, char_height := getBoundsOfFont(font_info.font_data, font_info.font_size, font_info.DPI)
	rect := image.Rect(0, 0, int(char_width*float64(chars_x)), int(char_height*float64(chars_y)))
	var images []*image.Paletted
	var delays []int
	for _, frame := range frames {
		delays = append(delays, delay/10)
		img_rgb := renderFrame(rect, frame, font_info, palette, char_width, char_height)
		pm := image.NewPaletted(img_rgb.Bounds(), palette)
		RGBtoPaletted(img_rgb, pm)
		images = append(images, pm)
	}
	gif.EncodeAll(out, &gif.GIF{
        Image: images,
        Delay: delays,
    })
}
