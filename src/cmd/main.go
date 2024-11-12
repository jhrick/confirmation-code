package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/jhrick/confirmation-code/internal/handlers"
	"github.com/jhrick/confirmation-code/internal/mail"
	"github.com/jhrick/confirmation-code/internal/utils/env"
)

func main() {
  env.LoadEnv()

  handler := handlers.Handlers{
    Router: http.NewServeMux(),
    MailService: mail.Init("", env.MailUsername, env.MailPassword, env.MailHost, env.MailPort),
  }

  handler.BindRoutes()

  go func() {
    err := http.ListenAndServe(":8080", handler.Router)

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
