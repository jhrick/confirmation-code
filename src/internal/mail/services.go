package mail

import "net/smtp"

func (m *Mail) Send(to []string, msg []byte) error {
  addr := m.Host + ":" + m.Port

  err := smtp.SendMail(addr, m.Auth, m.From, to, msg)
  if err != nil {
    return err
  }

  return nil
}
