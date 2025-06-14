package models

type Account struct {
	Id             int    `validate:"required"`
	MailAddress    string `validate:"required,email"`
	HashedPassword string `validate:"required,sha256"`
}

type AccountFragment struct {
	MailAddress    string `validate:"required,email"`
	HashedPassword string `validate:"required,sha256"`
}
