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
	JPEG: "jpeg",
	GIF:  "gif",
	PNG:  "png",
}

// Extention returns a file extention  give a mimetype or an error if it's not defined
func Extention(mimeType string) (string, error) {
	if ext, ok := Extentions[mimeType]; ok {
		return ext, nil
	}
	return "", ErrNotFound
}
