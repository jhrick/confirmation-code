package handlers

import (
	"encoding/json"
	"net/http"

	jsonutils "github.com/jhrick/confirmation-code/internal/utils/json"
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
    jsonutils.WriteJSON(w, http.StatusBadRequest, map[string]any{
      "success": false,
      "error": "invalid request body",
    })
    return
  }

  mailSubject := "Your code"
  code := h.CodeService.GenerateCode()

  msg := []byte("To:" + body.Email + "\r\n" +
		"Subject:" + mailSubject + "\r\n" +
		"\r\n" +
		code + "\r\n")

  err := h.MailService.Send([]string{body.Email}, msg)
  if err != nil {
    jsonutils.WriteJSON(w, http.StatusInternalServerError, map[string]any{
      "success": false,
      "error": "internal server error",
    })
    return
  }

  h.CacheManager.Store(code)

  jsonutils.WriteJSON(w, http.StatusCreated, map[string]any{
    "message": "send",
  })
}
