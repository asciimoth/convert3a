# convert3a
This tool allows you to convert [3a](https://github.com/DomesticMoth/3a) animations to png or gif formats  
Build:  
`$ git clone https://github.com/DomesticMoth/convert3a.git`  
`$ go build -o convert3a *.go`  
Usage:  
```
convert3a by DomesticMoth
Tool for convert 3a animations to media formats

Args:
	--in [file]			Path to 3a file (required)
	--out [file]			Path to output media file
	--dyn					Render animation instead of static image (flag)
	--preview [nomber]
	--fontsize [nomber]

You can change the palette by specifying values for the following parameters in hex format
	--ColorBlack
	--ColorLightRed
	--ColorLightGreen
	--ColorYellow
	--ColorLightBlue
	--ColorLightMagenta
	--ColorLightCyan
	--ColorLightWhite
	--ColorGray
	--ColorRed
	--ColorGreen
	--ColorBrown
	--ColorBlue
	--ColorMagenta
	--ColorCyan
	--ColorWhite
	--ColorDefaultBg
	--ColorDefaultFg
Example '--ColorRed #ff0000'
```
