package mime_test

import (
	"testing"

	"gitlab.com/mwarnerdotme/go-mime/mime"
)

func TestUnitFileExtensions(t *testing.T) {
	expectedResult := true
	fileExtensionsLength := len(mime.FileExtensions)
	if (fileExtensionsLength > 0) != expectedResult {
		t.Fatalf("file extension slice was not populated: received (%d) '%+v'", fileExtensionsLength, mime.FileExtensions)
	}
}

func TestUnitFileExtensionString(t *testing.T) {
	testFileExtension := mime.EXTENSION_JSON
	testFileExtensionString := ".json"

	if testFileExtension.String() != testFileExtensionString {
		t.Fatal("could not successfully return the file extension as a string")
	}
}
