package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handlers) handleCheckCode(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
    return
  }

  type _body struct {
    Code string `json:"code"`
  }

  var body _body
  
  if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
    http.Error(w, "invalid json", http.StatusBadRequest)
    return
  }

  correct := h.CacheManager.Check(body.Code)
  if !correct {
    http.Error(w, "incorrect code", http.StatusBadRequest)
    return
  }

  w.WriteHeader(http.StatusOK)
  fmt.Fprintln(w, "success")
}
