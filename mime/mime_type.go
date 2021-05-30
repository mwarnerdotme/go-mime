package mime

// Root of a MIME type (ie: {RootType}/{SubType}).
type RootType string

const (
	ROOT_TYPE_APPLICATION RootType = "application"
	ROOT_TYPE_AUDIO       RootType = "audio"
	ROOT_TYPE_FONT        RootType = "font"
	ROOT_TYPE_IMAGE       RootType = "image"
	ROOT_TYPE_MULTIPART   RootType = "multipart"
	ROOT_TYPE_TEXT        RootType = "text"
	ROOT_TYPE_VIDEO       RootType = "video"
)

// Slice containing all available root MIME types.
var RootTypes = []RootType{
	ROOT_TYPE_APPLICATION,
	ROOT_TYPE_AUDIO,
	ROOT_TYPE_FONT,
	ROOT_TYPE_IMAGE,
	ROOT_TYPE_MULTIPART,
	ROOT_TYPE_TEXT,
	ROOT_TYPE_VIDEO,
}

// Returns the RootType as a string.
func (rt *RootType) String() string {
	return string(*rt)
}

// A MIME type (content type in HTTP).
type MIMEType string

// Slice containing all available MIME types.
var MIMETypes = func() []MIMEType {
	var mt []MIMEType

	mt = append(mt, ApplicationMIMETypes...)
	mt = append(mt, AudioMIMETypes...)
	mt = append(mt, FontMIMETypes...)
	mt = append(mt, ImageMIMETypes...)
	mt = append(mt, MultipartMIMETypes...)
	mt = append(mt, TextMIMETypes...)
	mt = append(mt, VideoMIMETypes...)

	return mt
}()

// Returns the MIMEType as a string.
func (mt *MIMEType) String() string {
	return string(*mt)
}

// Map of file extensions to their corresponding MIME types
type MIMETypeMap = map[FileExtension]MIMEType

// Map of all file extensions to their MIME types.
var Map = func() MIMETypeMap {
	var mtm = MIMETypeMap{}

	maps := []MIMETypeMap{
		ApplicationMIMETypeMap,
		AudioMIMETypeMap,
		FontMIMETypeMap,
		ImageMIMETypeMap,
		MultipartMIMETypeMap,
		TextMIMETypeMap,
		VideoMIMETypeMap,
	}

	for i := range maps {
		for k, v := range maps[i] {
			mtm[k] = v
		}
	}

	return mtm
}()
