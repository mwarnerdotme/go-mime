package mime

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

	// aliases/alternatives

	IMAGE_TIF = IMAGE_TIFF
	IMAGE_JPG = IMAGE_JPEG
)

// slice containing all available image MIME types
func GetImageMIMETypes() []string {
	return []string{
		IMAGE_BMP,
		IMAGE_GIF,
		IMAGE_ICO,
		IMAGE_JPEG,
		IMAGE_PNG,
		IMAGE_SVG,
		IMAGE_TIFF,
		IMAGE_WEBP,
	}
}
