package json

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, statusCode int, res any) error {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(statusCode)
  
  if err := json.NewEncoder(w).Encode(res); err != nil {
    return fmt.Errorf("failed to encode json %w", err)
  }

  return nil
}
