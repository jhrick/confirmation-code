package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jhrick/confirmation-code/internal/cache"
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

  signal := h.CacheManager.Check(body.Code)
  if signal != cache.Ok {
    var codeStatus string
    switch signal {
      case cache.NotFound:
        codeStatus = "incorrect"
      case cache.Expired:
        codeStatus = "expired"
    }

      http.Error(w, codeStatus + " code", http.StatusBadRequest)
      return
  }

  w.WriteHeader(http.StatusOK)
  fmt.Fprintln(w, "success")
}
