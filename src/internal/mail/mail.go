package mail

import "net/smtp"

type Mail struct {
  Auth smtp.Auth
  From string
  Host string
  Port string
}

func Init(identity, username, password, host, port string) Mail {
  return Mail{
    Auth: smtp.PlainAuth(identity, username, password, host),
    From: username,
    Host: host,
    Port: port,
  }
}
