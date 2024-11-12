package handlers

import (
	"net/http"

	"github.com/jhrick/confirmation-code/internal/mail"
)

type Handlers struct {
  Router      *http.ServeMux
  MailService mail.Mail
}
