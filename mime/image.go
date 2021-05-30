package mime

/*
Image MIME types
*/

const (
	// .bmp: windows bitmap graphics
	IMAGE_BMP = "image/bmp"
	// .gif: graphics interchange
	IMAGE_GIF = "image/gif"
	// .ico: icon
	IMAGE_ICO = "image/vnd.microsoft.icon"
	// .jpeg: jpeg images
	IMAGE_JPEG = "image/jpeg"
	// .png: portable network graphics
	IMAGE_PNG = "image/png"
	// .svg: scalable vector graphics
	IMAGE_SVG = "image/svg+xml"
	// .tiff: tagged image file
	IMAGE_TIFF = "image/tiff"
	// .webp: webp image
	IMAGE_WEBP = "image/webp"
)

// aliases/alternatives

const (
	IMAGE_TIF = IMAGE_TIFF
	IMAGE_JPG = IMAGE_JPEG
)

/*
Image content types (HTTP only MIME types) - no file extension
*/

const ()

// slice containing all available image MIME types
var ImageMIMETypes = []string{
	IMAGE_BMP,
	IMAGE_GIF,
	IMAGE_ICO,
	IMAGE_JPEG,
	IMAGE_PNG,
	IMAGE_SVG,
	IMAGE_TIFF,
	IMAGE_WEBP,
}

/*
Image file extensions
*/

const (
	EXTENSION_BMP  FileExtension = ".bmp"
	EXTENSION_GIF  FileExtension = ".gif"
	EXTENSION_ICO  FileExtension = ".ico"
	EXTENSION_JPEG FileExtension = ".jpeg"
	EXTENSION_JPG  FileExtension = ".jpg"
	EXTENSION_PNG  FileExtension = ".png"
	EXTENSION_SVG  FileExtension = ".svg"
	EXTENSION_TIFF FileExtension = ".tiff"
	EXTENSION_WEBP FileExtension = ".webp"
)

// slice containing all available image file extensions
var ImageFileExtensions = []FileExtension{
	EXTENSION_BMP,
	EXTENSION_GIF,
	EXTENSION_ICO,
	EXTENSION_JPEG,
	EXTENSION_JPG,
	EXTENSION_PNG,
	EXTENSION_SVG,
	EXTENSION_TIFF,
	EXTENSION_WEBP,
}
