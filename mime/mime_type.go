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

// Returns the root type as a string.
func (rt *RootType) String() string {
	return string(*rt)
}

// A MIME type (content type in HTTP).
type MimeType string

// Returns the mime type as a string.
func (mt *MimeType) String() string {
	return string(*mt)
}
