package mime

type MimeType string

func (mt *MimeType) String() string {
	return string(*mt)
}
