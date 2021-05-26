package mime

const (
	// .aac: aac audio
	AUDIO_AAC = "audio/aac"
	// .midi: musical insrument digital interface
	AUDIO_MIDI = "audio/midi"
	// .mp3: mp3 audio
	AUDIO_MP3 = "audio/mpeg"
	// .oga: ogg audio
	AUDIO_OGA = "audio/ogg"
	// .opus: opus audio
	AUDIO_OPUS = "audio/opus"
	// .wav: waveform audio
	AUDIO_WAV = "audio/wav"
	// .weba: webm audio
	AUDIO_WEBA = "audio/webm"
	// .midi: musical insrument digital interface
	AUDIO_X_MIDI = "audio/x-midi"
	// .3gp: audio/video container
	AUDIO_3GP = "audio/3gpp"
	// .3g2: audio/video container
	AUDIO_3G2 = "audio/3gpp2"

	// aliases/alternatives

	AUDIO_MID   = AUDIO_MIDI
	AUDIO_OGG   = AUDIO_OGA
	AUDIO_X_MID = AUDIO_X_MIDI
)

// slice containing all available audio MIME types
func GetAudioMIMETypes() []string {
	return []string{
		AUDIO_AAC,
		AUDIO_MIDI,
		AUDIO_MP3,
		AUDIO_OGA,
		AUDIO_OPUS,
		AUDIO_WAV,
		AUDIO_WEBA,
		AUDIO_X_MIDI,
		AUDIO_3GP,
		AUDIO_3G2,
	}
}
