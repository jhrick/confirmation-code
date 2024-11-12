package handlers

import (
	"fmt"
	"net/http"
)

func (h *Handlers) BindRoutes() {
  h.Router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "pong")
  })

  h.Router.HandleFunc("/smtp", h.handleSendMail)
}
