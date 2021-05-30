package mime

import (
	"errors"
	"fmt"
	"strings"

	stdMIME "mime"
)

/*
Finds the best MIME type based on the provided extension and file type.

File types can be found in the same package (mime.ROOT_TYPE_*). They correspond to the roots of MIME types like application, image, text, etc.
*/
func LookupByFileExtension(extension FileExtension, rootType RootType) (mime string, err error) {
	// get appropriate slice of MIME types
	var mimeTypeSlice []MIMEType
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

	// get the best matching MIME type
	for i := range mimeTypeSlice {
		if strings.Contains(mimeTypeSlice[i].String(), strings.Trim(extension.String(), ".")) {
			return mimeTypeSlice[i].String(), nil
		}
	}

	// use std 'mime' package as a backup
	backupMIMEType := stdMIME.TypeByExtension(fmt.Sprintf(".%s", extension))

	if backupMIMEType != "" {
		return backupMIMEType, nil
	}

	// return
	notSupportedErrorMessage := fmt.Sprintf("provided extension is not supported - no associated MIME type for extension '%s'", extension)
	return "", errors.New(notSupportedErrorMessage)
}

/*
Finds the best MIME type based on the provided extension.

May return a mime type from an inappropriate root type. To fix this, use mime.LookupByFileExtension to apply a root type scope.
*/
func LookupByFileExtensionSimple(extension FileExtension) (mime string, err error) {
	// get the best matching MIME type
	for i := range MIMETypes {
		if strings.Contains(MIMETypes[i].String(), strings.Trim(extension.String(), ".")) {
			return MIMETypes[i].String(), nil
		}
	}

	// use std 'mime' package as a backup
	backupMIMEType := stdMIME.TypeByExtension(fmt.Sprintf(".%s", extension))

	if backupMIMEType != "" {
		return backupMIMEType, nil
	}

	// return
	notSupportedErrorMessage := fmt.Sprintf("provided extension is not supported - no associated MIME type for extension '%s'", extension)
	return "", errors.New(notSupportedErrorMessage)
}
