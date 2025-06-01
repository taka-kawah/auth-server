package tables

import (
	mailinterface "auth-server/interfaces/mail_interface"
	"auth-server/interfaces/repository"
)

type AccountService interface {
	SendEmail(mailAddress string, id int64) error
	Create(mailAddress string, hashedPassword string) (string, error)
	Login(mailAddress string, hashedPassword string) (string, error)
	UpdateMailAddress(mailAddress string, newMailAddress string) error
	UpdateHashedPassword(mailAddress string, newHashedPassword string) error
}

type accountServiceImpl struct {
	m mailinterface.MailSender
	r repository.AccountRepository
}

func NewAccountService(m mailinterface.MailSender, r repository.AccountRepository) *accountServiceImpl {
	return &accountServiceImpl{m: m, r: r}
}

func (a *accountServiceImpl) SendEmail(mailAddress string, id int64) error {
	return a.m.SendPasswordFromEmail(mailAddress, id)
}

func (a *accountServiceImpl) Create(mailAddress string, hashedPassword string) (string, error) {
	return a.r.Create(mailAddress, hashedPassword)
}

func (a *accountServiceImpl) Login(mailAddress string, hashedPassword string) (string, error) {
	return a.r.GetId(mailAddress, hashedPassword)
}

func (a *accountServiceImpl) UpdateMailAddress(mailAddress string, newMailAddress string) error {
	return a.r.UpdateByMailAddress(mailAddress, "mail_address", newMailAddress)
}

func (a *accountServiceImpl) UpdateHashedPassword(mailAddress string, newHashedPassword string) error {
	return a.r.UpdateByMailAddress(mailAddress, "hashed_password", newHashedPassword)
}
