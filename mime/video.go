package mime

/*
Video MIME types
*/

const (
	// .3g2: audio/video container
	VIDEO_3G2 = "video/3gpp2"
	// .3gp: audio/video container
	VIDEO_3GP = "video/3gpp"
	// .avi: audio video interleave
	VIDEO_AVI = "video/x-msvideo"
	// .mp4: mp4 video
	VIDEO_MP4 = "video/mp4"
	// .mpeg: mpeg video
	VIDEO_MPEG = "video/mpeg"
	// .ogv: ogg video
	VIDEO_OGV = "video/ogg"
	// .ts: mpeg transport stream
	VIDEO_TS = "video/mp2t"
	// .webm: webm video
	VIDEO_WEBM = "video/webm"
)

// aliases/alternatives

const (
	VIDEO_OGG = VIDEO_OGV
)

/*
Video content types (HTTP only MIME types) - no file extension
*/

const ()

// slice containing all available video MIME types
var VideoMIMETypes = []MimeType{
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
	EXTENSION_MPG  FileExtension = ".mpeg"
	EXTENSION_OGV  FileExtension = ".ogv"
	EXTENSION_TS   FileExtension = ".ts"
	EXTENSION_WEBM FileExtension = ".webm"
)

// slice containing all available video file extensions
var VideoFileExtensions = []FileExtension{
	EXTENSION_3G2,
	EXTENSION_3GP,
	EXTENSION_AVI,
	EXTENSION_MP4,
	EXTENSION_MPG,
	EXTENSION_OGV,
	EXTENSION_TS,
	EXTENSION_WEBM,
}
