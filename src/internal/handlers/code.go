package handlers

import (
	"encoding/json"
	"net/http"


	"github.com/jhrick/confirmation-code/internal/cache"
	jsonutils "github.com/jhrick/confirmation-code/internal/utils/json"
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

      jsonutils.WriteJSON(w, http.StatusBadRequest, map[string]any{
        "success": false,
        "error": codeStatus + " code",
      })
      return
  }

  jsonutils.WriteJSON(w, http.StatusOK, map[string]any{
    "success": true,
  })
}
