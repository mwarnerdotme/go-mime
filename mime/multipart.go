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
	MULTIPART_FORM MIMEType = "multipart/form-data"
)

// slice containing all available multipart MIME types
var MultipartMIMETypes = []MIMEType{
	MULTIPART_FORM,
}

/* Multipart file extensions */

const ()

// slice containing all available multipart file extensions
var MultipartFileExtensions = []FileExtension{}
