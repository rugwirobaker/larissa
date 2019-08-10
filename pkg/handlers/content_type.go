package handlers

import (
	"net/http"

	"github.com/rugwirobaker/larissa/pkg/types"
)

func ext(content []byte) (string, error) {
	filetype := mime(content)
	return types.Extention(filetype)
}

func mime(content []byte) string {
	return http.DetectContentType(content)
}
