package mime_test

import (
	"testing"

	"gitlab.com/mwarnerdotme/go-mime/mime"
)

func TestUnitMimeTypeString(t *testing.T) {
	testMimeType := mime.APPLICATION_JSON
	testMimeTypeString := "application/json"

	if testMimeType.String() != testMimeTypeString {
		t.Fatal("could not successfully return the mime type as a string")
	}
}

func TestUnitRootTypeString(t *testing.T) {
	testRootType := mime.ROOT_TYPE_APPLICATION
	testRootTypeString := "application"

	if testRootType.String() != testRootTypeString {
		t.Fatal("could not successfully return the root type as a string")
	}
}
