package mediauri

import (
	"fmt"
	"hestia/internal/logger"
	"slices"
	"strings"
)

var (
	MP3 = ".mp3"
	OGG = ".ogg"
	FLAC = ".flac"
	MP4 = ".mp4"
	M4A = ".m4a"
	WAV = ".wav"

	PDF = ".pdf"
	JPEG = ".jpeg"
	PNG = ".png"
	GIF = ".gif"
	JPG = ".jpg"

	IMAGE_EXT = []string{JPEG, PNG, GIF, JPG}
	AUDIO_EXT = []string{MP3, OGG, FLAC, MP4, M4A, WAV}
	PDF_EXT = []string{PDF}
)

/*
definition du type de fichier pour enregistrement du media

params:
	- extension : string not null
return:
	- string
	- error
	- type de fichier (audio, pdf, images)
*/
func DefineTypeFileMedia(extension string) (string, error) {
	extension = strings.ToLower(extension)

	switch {
	case slices.Contains(AUDIO_EXT, extension):
		return "audio", nil
	case slices.Contains(PDF_EXT, extension):
		return "pdf", nil
	case slices.Contains(IMAGE_EXT, extension):
		return "images", nil
	default:
		logger.Error("Type de fichier non supporté")
		return "", fmt.Errorf("type de fichier non supporté")
	}
}