package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
)

func main() {
  mux := http.NewServeMux()

  mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "pong\n")
  })

  go func() {
    err := http.ListenAndServe(":8080", mux)

    if err != nil {
      if !errors.Is(err, http.ErrServerClosed) {
        panic(err)
      }
    }
  }()

  fmt.Print("Server running!\r\n")

  quit := make(chan os.Signal, 1)
  signal.Notify(quit, os.Interrupt)
  <-quit

  fmt.Println("\nServer closed.")
}
