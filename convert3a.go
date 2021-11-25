package main

import  (
	"github.com/DomesticMoth/go3a"
	"io"
	"os"
	"fmt"
	"flag"
	"math"
	"strings"
	"path/filepath"
	"image"
	"image/color"
	"image/png"
	"image/gif"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/font/gofont/goitalic"
	"golang.org/x/image/font/gofont/gomono"
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

func Color3atoRGB(color go3a.Color) int{
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

func PrintHelp(){}

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
	for i, c := range opts.Colors{
		if *c == "" {
			continue
		}
		color, err := ParseHexColor(*c)
		if err == nil {
			palette[i] = color
		}else{
			panic(err)
		}
	}
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
