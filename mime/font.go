package mime

const (
	// otf: opentype font
	FONT_OTF = "font/otf"
	// ttf: truetype font
	FONT_TTF = "font/ttf"
	// woff: web open font
	FONT_WOFF = "font/woff"
	// woff2: web open font (version 2)
	FONT_WOFF2 = "font/woff2"
)

// slice containing all available font MIME types
func GetFontMIMETypes() []string {
	return []string{
		FONT_OTF,
		FONT_TTF,
		FONT_WOFF,
		FONT_WOFF2,
	}
}
