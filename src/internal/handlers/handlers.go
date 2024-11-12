package handlers

import (
	"net/http"

	"github.com/jhrick/confirmation-code/internal/mail"
	"github.com/jhrick/confirmation-code/internal/services"
)

type Handlers struct {
  Router      *http.ServeMux
  MailService mail.Mail
  CodeService services.CodeService
}
