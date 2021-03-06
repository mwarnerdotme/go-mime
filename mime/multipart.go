package mime

/*
Multipart MIME types
*/

const ()

// Aliases/alternatives

const ()

/*
Multipart content types (HTTP only MIME types) - no file extension
*/

const (
	// multipart form data
	MULTIPART_FORM MIMEType = "multipart/form-data"
)

// Slice containing all available multipart MIME types.
var MultipartMIMETypes = []MIMEType{
	MULTIPART_FORM,
}

/*
Multipart file extensions
*/

const ()

// Slice containing all available multipart file extensions.
var MultipartFileExtensions = []FileExtension{}

/*
Map of file extensions to MIME types
*/

var MultipartMIMETypeMap = MIMETypeMap{}
