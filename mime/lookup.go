package mime

import (
	"errors"
	"fmt"
	"strings"

	stdMime "mime"
)

// Root of any given MIME type (ie: {RootType}/*).
type RootType string

// Used to scope MIME type lookups.
const (
	ROOT_TYPE_APPLICATION RootType = "application"
	ROOT_TYPE_AUDIO       RootType = "audio"
	ROOT_TYPE_FONT        RootType = "font"
	ROOT_TYPE_IMAGE       RootType = "image"
	ROOT_TYPE_MULTIPART   RootType = "multipart"
	ROOT_TYPE_TEXT        RootType = "text"
	ROOT_TYPE_VIDEO       RootType = "video"
)

/*
Finds the best MIME type based on the provided extension and file type.

File types can be found in the same package (mime.CORE_TYPE_*). They correspond to the roots of MIME types like application, image, text, etc.
*/
func LookupByFileExtension(extension string, rootType RootType) (mime string, err error) {
	// get appropriate slice of MIME types
	var mimeTypeSlice []string
	switch rootType {
	case ROOT_TYPE_APPLICATION:
		mimeTypeSlice = ApplicationMIMETypes
	case ROOT_TYPE_AUDIO:
		mimeTypeSlice = AudioMIMETypes
	case ROOT_TYPE_FONT:
		mimeTypeSlice = FontMIMETypes
	case ROOT_TYPE_IMAGE:
		mimeTypeSlice = ImageMIMETypes
	// case ROOT_TYPE_MULTIPART: // multipart/form does not have a file extension
	// 	mimeTypeSlice = GetMultipartMIMETypes()
	case ROOT_TYPE_TEXT:
		mimeTypeSlice = TextMIMETypes
	case ROOT_TYPE_VIDEO:
		mimeTypeSlice = VideoMIMETypes
	default:
		invalidRootTypeErrorMessage := fmt.Sprintf("could not retrieve list of MIME types for '%s' - please use the mime.ROOT_TYPE_* constants for this lookup", rootType)
		return "", errors.New(invalidRootTypeErrorMessage)
	}

	// remove any period prefixes
	extension = strings.Trim(extension, ".")

	// get the best matching MIME type
	for i := range mimeTypeSlice {
		if strings.Contains(mimeTypeSlice[i], extension) {
			return mimeTypeSlice[i], nil
		}
	}

	// use std 'mime' package as a backup
	backupMIMEType := stdMime.TypeByExtension(fmt.Sprintf(".%s", extension))

	if backupMIMEType != "" {
		return backupMIMEType, nil
	}

	// return
	notSupportedErrorMessage := fmt.Sprintf("provided extension is not supported - no associated MIME type for extension '%s'", extension)
	return "", errors.New(notSupportedErrorMessage)
}

func (rt *RootType) String() string {
	return string(*rt)
}
