package types

import "errors"

// supported file mime types
const (
	JPEG = "image/jpeg"
	GIF  = "image/gif"
	PNG  = "image/png"
)

// ErrNotFound is returned when the given mimeType is not supported
var ErrNotFound = errors.New("mimeType not supported")

// Extentions maps file extenstions to mime types
var Extentions = map[string]string{
	"jpeg": JPEG,
	"gif":  GIF,
	"png":  PNG,
}

// Extention returns a file extention  give a mimetype or an error if it's not defined
func Extention(mimeType string) (string, error) {
	for key, value := range Extentions {
		if value == mimeType {
			return key, nil
		}
	}
	return "", ErrNotFound
}

// MimeType returns a mimeType given a file extension  or an error if it's not defined
func MimeType(extention string) (string, error) {
	if mime, ok := Extentions[extention]; ok {
		return mime, nil
	}
	return "", ErrNotFound
}
