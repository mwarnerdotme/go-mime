package mime

/*
Video MIME types
*/

const (
	// .3g2: audio/video container
	VIDEO_3G2 MIMEType = "video/3gpp2"
	// .3gp: audio/video container
	VIDEO_3GP MIMEType = "video/3gpp"
	// .avi: audio video interleave
	VIDEO_AVI MIMEType = "video/x-msvideo"
	// .mp4: mp4 video
	VIDEO_MP4 MIMEType = "video/mp4"
	// .mpeg: mpeg video
	VIDEO_MPEG MIMEType = "video/mpeg"
	// .ogv: ogg video
	VIDEO_OGV MIMEType = "video/ogg"
	// .ts: mpeg transport stream
	VIDEO_TS MIMEType = "video/mp2t"
	// .webm: webm video
	VIDEO_WEBM MIMEType = "video/webm"
)

// Aliases/alternatives

const (
	VIDEO_OGG = VIDEO_OGV
)

/*
Video content types (HTTP only MIME types) - no file extension
*/

const ()

// Slice containing all available video MIME types.
var VideoMIMETypes = []MIMEType{
	VIDEO_3G2,
	VIDEO_3GP,
	VIDEO_AVI,
	VIDEO_MP4,
	VIDEO_MPEG,
	VIDEO_OGV,
	VIDEO_TS,
	VIDEO_WEBM,
}

/*
Video file extensions
*/

const (
	// EXTENSION_3G2  FileExtension = ".3g2" // already exists in audio
	// EXTENSION_3GP  FileExtension = ".3gp" // already exists in audio
	EXTENSION_AVI  FileExtension = ".avi"
	EXTENSION_MP4  FileExtension = ".mp4"
	EXTENSION_MPEG FileExtension = ".mpeg"
	EXTENSION_MPG  FileExtension = ".mpg"
	EXTENSION_OGV  FileExtension = ".ogv"
	EXTENSION_TS   FileExtension = ".ts"
	EXTENSION_WEBM FileExtension = ".webm"
)

// Slice containing all available video file extensions.
var VideoFileExtensions = []FileExtension{
	EXTENSION_3G2,
	EXTENSION_3GP,
	EXTENSION_AVI,
	EXTENSION_MP4,
	EXTENSION_MPEG,
	EXTENSION_MPG,
	EXTENSION_OGV,
	EXTENSION_TS,
	EXTENSION_WEBM,
}

/*
Map of file extensions to MIME types
*/

var VideoMIMETypeMap = MIMETypeMap{
	EXTENSION_3G2:  VIDEO_3G2,
	EXTENSION_3GP:  VIDEO_3GP,
	EXTENSION_AVI:  VIDEO_AVI,
	EXTENSION_MP4:  VIDEO_MP4,
	EXTENSION_MPEG: VIDEO_MPEG,
	EXTENSION_MPG:  VIDEO_MPEG,
	EXTENSION_OGV:  VIDEO_OGV,
	EXTENSION_TS:   VIDEO_TS,
	EXTENSION_WEBM: VIDEO_WEBM,
}
