package mime

// A file extension (ex: .json).
type FileExtension string

// Slice containing all available file extensions.
var FileExtensions = func() []FileExtension {
	var fe []FileExtension

	fe = append(fe, ApplicationFileExtensions...)
	fe = append(fe, AudioFileExtensions...)
	fe = append(fe, FontFileExtensions...)
	fe = append(fe, ImageFileExtensions...)
	fe = append(fe, MultipartFileExtensions...)
	fe = append(fe, TextFileExtensions...)
	fe = append(fe, VideoFileExtensions...)

	return fe
}()

// Returns the FileExtension as a string.
func (fe *FileExtension) String() string {
	return string(*fe)
}
