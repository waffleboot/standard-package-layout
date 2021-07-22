package http

import (
	"net/http"

	"github.com/benbjohnson/myapp"
)

type Handler struct {
	UserService myapp.UserService
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
