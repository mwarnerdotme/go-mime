package mime

type FileExtension string

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
}

func (fe *FileExtension) String() string {
	return string(*fe)
}
