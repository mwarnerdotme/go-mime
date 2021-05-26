package mime

const (
	// http only

	// multipart form data
	MULTIPART_FORM = "multipart/form-data"
)

// slice containing all available multipart MIME types
func GetMultipartMIMETypes() []string {
	return []string{
		MULTIPART_FORM,
	}
}
