package mime

const (
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
	// .3gp: audio/video container
	VIDEO_3GP = "video/3gpp"
	// .3g2: audio/video container
	VIDEO_3G2 = "video/3gpp2"

	// aliases/alternatives

	VIDEO_OGG = VIDEO_OGV
)

// slice containing all available video MIME types
func GetVideoMIMETypes() []string {
	return []string{
		VIDEO_AVI,
		VIDEO_MP4,
		VIDEO_MPEG,
		VIDEO_OGV,
		VIDEO_TS,
		VIDEO_WEBM,
		VIDEO_3GP,
		VIDEO_3G2,
	}
}
