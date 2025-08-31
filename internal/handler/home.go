package handler

import (
	"io"
	"net/http"
)

type homeHandler struct {
}

func (h homeHandler) handleIndex(w http.ResponseWriter, r *http.Request) error {
	_, err := io.WriteString(w, "Hello World")
	if err != nil {
		return err
	}
	return nil
}
