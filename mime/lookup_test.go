package mime_test

import (
	"testing"

	"gitlab.com/mwarnerdotme/go-mime/mime"
)

func TestUnitLookupByFileExtensionApplication(t *testing.T) {
	result, err := mime.LookupByFileExtension(".json", mime.ROOT_TYPE_APPLICATION)

	if err != nil {
		t.Fatal(err)
	}

	expectedResult := "application/json"
	if result != expectedResult {
		t.Fatalf("an incorrect MIME type was returned: expected '%s' and received '%s'", expectedResult, result)
	}
}

func TestUnitLookupByFileExtensionAudio(t *testing.T) {
	result, err := mime.LookupByFileExtension(".mp3", mime.ROOT_TYPE_AUDIO)

	if err != nil {
		t.Fatal(err)
	}

	expectedResult := "audio/mpeg"
	if result != expectedResult {
		t.Fatalf("an incorrect MIME type was returned: expected '%s' and received '%s'", expectedResult, result)
	}
}

func TestUnitLookupByFileExtensionFont(t *testing.T) {
	result, err := mime.LookupByFileExtension(".ttf", mime.ROOT_TYPE_FONT)

	if err != nil {
		t.Fatal(err)
	}

	expectedResult := "font/ttf"
	if result != expectedResult {
		t.Fatalf("an incorrect MIME type was returned: expected '%s' and received '%s'", expectedResult, result)
	}
}

func TestUnitLookupByFileExtensionImage(t *testing.T) {
	result, err := mime.LookupByFileExtension(".jpeg", mime.ROOT_TYPE_IMAGE)

	if err != nil {
		t.Fatal(err)
	}

	expectedResult := "image/jpeg"
	if result != expectedResult {
		t.Fatalf("an incorrect MIME type was returned: expected '%s' and received '%s'", expectedResult, result)
	}
}

func TestUnitLookupByFileExtensionText(t *testing.T) {
	result, err := mime.LookupByFileExtension(".csv", mime.ROOT_TYPE_TEXT)

	if err != nil {
		t.Fatal(err)
	}

	expectedResult := "text/csv"
	if result != expectedResult {
		t.Fatalf("an incorrect MIME type was returned: expected '%s' and received '%s'", expectedResult, result)
	}
}

func TestUnitLookupByFileExtensionVideo(t *testing.T) {
	result, err := mime.LookupByFileExtension(".mp4", mime.ROOT_TYPE_VIDEO)

	if err != nil {
		t.Fatal(err)
	}

	expectedResult := "video/mp4"
	if result != expectedResult {
		t.Fatalf("an incorrect MIME type was returned: expected '%s' and received '%s'", expectedResult, result)
	}
}

func TestUnitLookupByFileExtensionInvalidExtension(t *testing.T) {
	_, err := mime.LookupByFileExtension(".test", mime.ROOT_TYPE_TEXT)

	if err == nil {
		t.Fatal("test did not recognize the invalid file extension")
	}
}

func TestUnitLookupByFileExtensionInvalidRootType(t *testing.T) {
	_, err := mime.LookupByFileExtension(".csv", "test")

	if err == nil {
		t.Fatal("test did not recognize the invalid root MIME type")
	}
}

func TestUnitRootTypeString(t *testing.T) {
	testRootType := mime.ROOT_TYPE_APPLICATION
	testRootTypeString := "application"

	if testRootType.String() != testRootTypeString {
		t.Fatal("could not successfully return the root type as a string")
	}
}
