package types

import (
	"github.com/rugwirobaker/larissa/pkg/errors"
)

// supported file mime types
const (
	JPEG = "image/jpeg"
	GIF  = "image/gif"
	PNG  = "image/png"
)

// Extentions maps file extenstions to mime types
var Extentions = map[string]string{
	"jpeg": JPEG,
	"gif":  GIF,
	"png":  PNG,
}

// Extention returns a file extention  give a mimetype or an error if it's not defined
func Extention(mimeType string) (string, error) {
	const op errors.Op = "types.Extention"

	for key, value := range Extentions {
		if value == mimeType {
			return key, nil
		}
	}
	return "", errors.E(op, "content type not supported")
}

// MimeType returns a mimeType given a file extension  or an error if it's not defined
func MimeType(extention string) (string, error) {
	const op errors.Op = "types.MimeType"

	if mime, ok := Extentions[extention]; ok {
		return mime, nil
	}
	return "", errors.E(op, "content type not supported")
}
