package mime

/*
Audio MIME types
*/

const (
	// .3g2: audio/video container
	AUDIO_3G2 MIMEType = "audio/3gpp2"
	// .3gp: audio/video container
	AUDIO_3GP MIMEType = "audio/3gpp"
	// .aac: aac audio
	AUDIO_AAC MIMEType = "audio/aac"
	// .midi: musical insrument digital interface
	AUDIO_MIDI MIMEType = "audio/midi"
	// .mp3: mp3 audio
	AUDIO_MP3 MIMEType = "audio/mpeg"
	// .oga: ogg audio
	AUDIO_OGA MIMEType = "audio/ogg"
	// .opus: opus audio
	AUDIO_OPUS MIMEType = "audio/opus"
	// .wav: waveform audio
	AUDIO_WAV MIMEType = "audio/wav"
	// .weba: webm audio
	AUDIO_WEBA MIMEType = "audio/webm"
	// .midi: musical insrument digital interface
	AUDIO_X_MIDI MIMEType = "audio/x-midi"
)

// Aliases/alternatives

const (
	AUDIO_MID   = AUDIO_MIDI
	AUDIO_OGG   = AUDIO_OGA
	AUDIO_X_MID = AUDIO_X_MIDI
)

/*
Audio content types (HTTP only MIME types) - no file extension
*/

const ()

// Slice containing all available audio MIME types.
var AudioMIMETypes = []MIMEType{
	AUDIO_3G2,
	AUDIO_3GP,
	AUDIO_AAC,
	AUDIO_MIDI,
	AUDIO_MP3,
	AUDIO_OGA,
	AUDIO_OPUS,
	AUDIO_WAV,
	AUDIO_WEBA,
	AUDIO_X_MIDI,
}

/*
Audio file extensions
*/

const (
	EXTENSION_3G2  FileExtension = ".3g2"
	EXTENSION_3GP  FileExtension = ".3gp"
	EXTENSION_AAC  FileExtension = ".aac"
	EXTENSION_MIDI FileExtension = ".midi"
	EXTENSION_MID  FileExtension = ".mid"
	EXTENSION_MP3  FileExtension = ".mp3"
	EXTENSION_OGA  FileExtension = ".oga"
	EXTENSION_OPUS FileExtension = ".opus"
	EXTENSION_WAV  FileExtension = ".wav"
	EXTENSION_WEBA FileExtension = ".weba"
)

// Slice containing all available audio file extensions.
var AudioFileExtensions = []FileExtension{
	EXTENSION_3G2,
	EXTENSION_3GP,
	EXTENSION_AAC,
	EXTENSION_MIDI,
	EXTENSION_MID,
	EXTENSION_MP3,
	EXTENSION_OGA,
	EXTENSION_OPUS,
	EXTENSION_WAV,
	EXTENSION_WEBA,
}

/*
Map of file extensions to MIME types
*/

var AudioMIMETypeMap = MIMETypeMap{
	EXTENSION_3G2:  AUDIO_3G2,
	EXTENSION_3GP:  AUDIO_3GP,
	EXTENSION_AAC:  AUDIO_AAC,
	EXTENSION_MIDI: AUDIO_MIDI,
	EXTENSION_MID:  AUDIO_MID,
	EXTENSION_MP3:  AUDIO_MP3,
	EXTENSION_OGA:  AUDIO_OGA,
	EXTENSION_OPUS: AUDIO_OPUS,
	EXTENSION_WAV:  AUDIO_WAV,
	EXTENSION_WEBA: AUDIO_WEBA,
}
