package handlers

import (
	"github.com/rs/zerolog/log"
	"net/http"
)

func (h *Handlers) HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	_, err := w.Write([]byte("Hello"))

	if err != nil {
		log.Log().Err(err)
	}
}