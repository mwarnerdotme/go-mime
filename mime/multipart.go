package mime

/*
Multipart MIME types
*/

const ()

// aliases/alternatives

const ()

/*
Multipart content types (HTTP only MIME types) - no file extension
*/

const (
	// multipart form data
	MULTIPART_FORM MimeType = "multipart/form-data"
)

// slice containing all available multipart MIME types
var MultipartMIMETypes = []MimeType{
	MULTIPART_FORM,
}

/* Multipart file extensions */

const ()

// slice containing all available multipart file extensions
var MultipartFileExtensions = []FileExtension{}
