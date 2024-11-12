package env

import (
	"os"

	"github.com/joho/godotenv"
)

var (
  MailUsername string
  MailPassword string
  MailHost     string
  MailPort     string
)

func LoadEnv() {
  if err := godotenv.Load("../.env"); err != nil {
    panic(err)
  }

  MailUsername = os.Getenv("MAIL_USERNAME")
  MailPassword = os.Getenv("MAIL_PASSWORD")
  MailHost = os.Getenv("MAIL_HOST")
  MailPort = os.Getenv("MAIL_PORT")

}
