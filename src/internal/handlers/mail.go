package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (h *Handlers) handleSendMail(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
    return
  }

  type _body struct {
    Email string `json:"email"`
  }

  var body _body

  if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
    http.Error(w, "invalid json", http.StatusBadRequest)
    return
  }

  mailSubject := "Your code"
  mailBody := h.CodeService.GenerateCode()

  msg := []byte("To:" + body.Email + "\r\n" +
		"Subject:" + mailSubject + "\r\n" +
		"\r\n" +
		mailBody + "\r\n")

  err := h.MailService.Send([]string{body.Email}, msg)
  if err != nil {
    log.Println(err)
    http.Error(w, "internal server error", http.StatusInternalServerError)
    return
  }

  w.WriteHeader(http.StatusCreated)
  fmt.Fprintln(w, "send")
}
