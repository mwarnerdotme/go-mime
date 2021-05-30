package mime_test

import (
	"testing"

	"gitlab.com/mwarnerdotme/go-mime/mime"
)

/*
Lookup by file extension
*/

func TestUnitLookupByFileExtension(t *testing.T) {
	for _, rootType := range mime.RootTypes {
		var fe []mime.FileExtension
		switch rootType {
		case mime.ROOT_TYPE_APPLICATION:
			fe = mime.ApplicationFileExtensions
		case mime.ROOT_TYPE_AUDIO:
			fe = mime.AudioFileExtensions
		case mime.ROOT_TYPE_FONT:
			fe = mime.FontFileExtensions
		case mime.ROOT_TYPE_IMAGE:
			fe = mime.ImageFileExtensions
		case mime.ROOT_TYPE_MULTIPART:
			fe = mime.MultipartFileExtensions
		case mime.ROOT_TYPE_TEXT:
			fe = mime.TextFileExtensions
		case mime.ROOT_TYPE_VIDEO:
			fe = mime.VideoFileExtensions
		}

		for _, extension := range fe {
			result, err := mime.LookupByFileExtension(extension, rootType)

			if err != nil {
				t.Fatal(err)
			}

			if result == "" {
				t.Fatalf("no MIME type was returned for extension '%s'", extension)
			}
		}
	}
}

func TestUnitLookupByFileExtensionInvalidRootType(t *testing.T) {
	invalidRootType := mime.RootType("test")
	_, err := mime.LookupByFileExtension(mime.EXTENSION_3G2, invalidRootType)

	if err == nil {
		t.Fatal("test did not recognize invalid root type")
	}
}

func TestUnitLookupByFileExtensionInvalidFileExtension(t *testing.T) {
	invalidFileExtension := mime.FileExtension("test")
	_, err := mime.LookupByFileExtension(invalidFileExtension, mime.ROOT_TYPE_APPLICATION)

	if err == nil {
		t.Fatal("test did not recognize invalid file extension")
	}
}

/*
Lookup by file extension - simple
*/

func TestUnitLookupByFileExtensionSimple(t *testing.T) {
	for _, extension := range mime.FileExtensions {
		result, err := mime.LookupByFileExtensionSimple(extension)

		if err != nil {
			t.Fatal(err)
		}

		if result == "" {
			t.Fatalf("no MIME type was returned for extension '%s'", extension)
		}
	}
}

func TestUnitLookupByFileExtensionSimpleInvalidFileExtension(t *testing.T) {
	invalidFileExtension := mime.FileExtension("test")
	_, err := mime.LookupByFileExtensionSimple(invalidFileExtension)

	if err == nil {
		t.Fatal("test did not recognize invalid file extension")
	}
}
