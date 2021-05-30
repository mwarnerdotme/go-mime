package mime

import (
	"errors"
	"fmt"

	stdMIME "mime"
)

/*
Finds the best MIME type based on the provided extension and file type.

File types can be found in the same package (mime.ROOT_TYPE_*). They correspond to the roots of MIME types like application, image, text, etc.
*/
func LookupByFileExtension(extension FileExtension, rootType RootType) (string, error) {
	// get appropriate MIMETypeMap
	var mtm MIMETypeMap
	switch rootType {
	case ROOT_TYPE_APPLICATION:
		mtm = ApplicationMIMETypeMap
	case ROOT_TYPE_AUDIO:
		mtm = AudioMIMETypeMap
	case ROOT_TYPE_FONT:
		mtm = FontMIMETypeMap
	case ROOT_TYPE_IMAGE:
		mtm = ImageMIMETypeMap
	case ROOT_TYPE_MULTIPART:
		mtm = MultipartMIMETypeMap
	case ROOT_TYPE_TEXT:
		mtm = TextMIMETypeMap
	case ROOT_TYPE_VIDEO:
		mtm = VideoMIMETypeMap
	default:
		invalidRootTypeErrorMessage := fmt.Sprintf("could not retrieve list of MIME types for '%s' - please use the mime.ROOT_TYPE_* constants for this lookup", rootType)
		return "", errors.New(invalidRootTypeErrorMessage)
	}

	// get the MIME type from the map
	if mtm[extension] != "" {
		mt := mtm[extension]
		return mt.String(), nil
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
func LookupByFileExtensionSimple(extension FileExtension) (string, error) {
	// get the MIME type from the aggregate map
	if Map[extension] != "" {
		mt := Map[extension]
		return mt.String(), nil
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
