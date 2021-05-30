package mime

/*
Font MIME types
*/

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

// aliases/alternatives

const ()

/*
Font content types (HTTP only MIME types) - no file extension
*/

const ()

// slice containing all available font MIME types
var FontMIMETypes = []MimeType{
	FONT_OTF,
	FONT_TTF,
	FONT_WOFF,
	FONT_WOFF2,
}

/*
Font file extensions
*/

const (
	EXTENSION_OTF   FileExtension = ".otf"
	EXTENSION_TTF   FileExtension = ".ttf"
	EXTENSION_WOFF  FileExtension = ".woff"
	EXTENSION_WOFF2 FileExtension = ".woff2"
)

// slice containing all available font file extensions
var FontFileExtensions = []FileExtension{
	EXTENSION_OTF,
	EXTENSION_TTF,
	EXTENSION_WOFF,
	EXTENSION_WOFF2,
}
