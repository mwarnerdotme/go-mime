package mime_test

import (
	"testing"

	"gitlab.com/mwarnerdotme/go-mime/mime"
)

/*
Lookup by file extension
*/
func TestUnitLookupByFileExtensionApplication(t *testing.T) {
	result, err := mime.LookupByFileExtension(mime.EXTENSION_JSON, mime.ROOT_TYPE_APPLICATION)

	if err != nil {
		t.Fatal(err)
	}

	expectedResult := "application/json"
	if result != expectedResult {
		t.Fatalf("an incorrect MIME type was returned: expected '%s' and received '%s'", expectedResult, result)
	}
}

func TestUnitLookupByFileExtensionAudio(t *testing.T) {
	result, err := mime.LookupByFileExtension(mime.EXTENSION_AAC, mime.ROOT_TYPE_AUDIO)

	if err != nil {
		t.Fatal(err)
	}

	expectedResult := "audio/aac"
	if result != expectedResult {
		t.Fatalf("an incorrect MIME type was returned: expected '%s' and received '%s'", expectedResult, result)
	}
}

func TestUnitLookupByFileExtensionFont(t *testing.T) {
	result, err := mime.LookupByFileExtension(mime.EXTENSION_TTF, mime.ROOT_TYPE_FONT)

	if err != nil {
		t.Fatal(err)
	}

	expectedResult := "font/ttf"
	if result != expectedResult {
		t.Fatalf("an incorrect MIME type was returned: expected '%s' and received '%s'", expectedResult, result)
	}
}

func TestUnitLookupByFileExtensionImage(t *testing.T) {
	result, err := mime.LookupByFileExtension(mime.EXTENSION_JPEG, mime.ROOT_TYPE_IMAGE)

	if err != nil {
		t.Fatal(err)
	}

	expectedResult := "image/jpeg"
	if result != expectedResult {
		t.Fatalf("an incorrect MIME type was returned: expected '%s' and received '%s'", expectedResult, result)
	}
}

func TestUnitLookupByFileExtensionText(t *testing.T) {
	result, err := mime.LookupByFileExtension(mime.EXTENSION_CSV, mime.ROOT_TYPE_TEXT)

	if err != nil {
		t.Fatal(err)
	}

	expectedResult := "text/csv"
	if result != expectedResult {
		t.Fatalf("an incorrect MIME type was returned: expected '%s' and received '%s'", expectedResult, result)
	}
}

func TestUnitLookupByFileExtensionVideo(t *testing.T) {
	result, err := mime.LookupByFileExtension(mime.EXTENSION_MP4, mime.ROOT_TYPE_VIDEO)

	if err != nil {
		t.Fatal(err)
	}

	expectedResult := "video/mp4"
	if result != expectedResult {
		t.Fatalf("an incorrect MIME type was returned: expected '%s' and received '%s'", expectedResult, result)
	}
}

func TestUnitLookupByFileExtensionInvalidExtension(t *testing.T) {
	_, err := mime.LookupByFileExtension(mime.FileExtension(".test"), mime.ROOT_TYPE_TEXT)

	if err == nil {
		t.Fatal("test did not recognize the invalid file extension")
	}
}

func TestUnitLookupByFileExtensionInvalidRootType(t *testing.T) {
	_, err := mime.LookupByFileExtension(mime.EXTENSION_CSV, mime.RootType("test"))

	if err == nil {
		t.Fatal("test did not recognize the invalid root MIME type")
	}
}

/*
Lookup by file extension - simple
*/

func TestUnitLookupByFileExtensionSimpleApplication(t *testing.T) {
	result, err := mime.LookupByFileExtensionSimple(mime.EXTENSION_JSON)

	if err != nil {
		t.Fatal(err)
	}

	expectedResult := "application/json"
	if result != expectedResult {
		t.Fatalf("an incorrect MIME type was returned: expected '%s' and received '%s'", expectedResult, result)
	}
}
