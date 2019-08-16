package handlers

import (
	"net/http"
)

func mime(content []byte) string {
	return http.DetectContentType(content)
}
