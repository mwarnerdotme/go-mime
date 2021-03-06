package mime

/*
Font MIME types
*/

const (
	// otf: opentype font
	FONT_OTF MIMEType = "font/otf"
	// ttf: truetype font
	FONT_TTF MIMEType = "font/ttf"
	// woff: web open font
	FONT_WOFF MIMEType = "font/woff"
	// woff2: web open font (version 2)
	FONT_WOFF2 MIMEType = "font/woff2"
)

// Aliases/alternatives

const ()

/*
Font content types (HTTP only MIME types) - no file extension
*/

const ()

// Slice containing all available font MIME types.
var FontMIMETypes = []MIMEType{
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

// Slice containing all available font file extensions.
var FontFileExtensions = []FileExtension{
	EXTENSION_OTF,
	EXTENSION_TTF,
	EXTENSION_WOFF,
	EXTENSION_WOFF2,
}

/*
Map of file extensions to MIME types
*/

var FontMIMETypeMap = MIMETypeMap{
	EXTENSION_OTF:   FONT_OTF,
	EXTENSION_TTF:   FONT_TTF,
	EXTENSION_WOFF:  FONT_WOFF,
	EXTENSION_WOFF2: FONT_WOFF2,
}
