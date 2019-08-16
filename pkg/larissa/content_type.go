package larissa

import (
	"net/http"

	"github.com/rugwirobaker/larissa/pkg/types"
)

func ext(content []byte) (string, error) {
	filetype := http.DetectContentType(content)
	return types.Extention(filetype)
}

