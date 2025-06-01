package repository

import (
	"auth-server/util"
	"context"
	"database/sql"
	"errors"
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

type AccountRepository interface {
	Create(mailAddress string, hashedPassword string) (string, error)
	GetId(mailAddress string, hashedPassword string) (string, error)
	UpdateById(id int64, column string, value string) error
	UpdateByMailAddress(mailAddress string, column string, value string) error
}

type accountRepositoryImpl struct {
	ctx context.Context
	db  *sql.DB
}

func NewAccountRepository(ctx context.Context, db *sql.DB) *accountRepositoryImpl {
	return &accountRepositoryImpl{ctx: ctx, db: db}
}

func (r *accountRepositoryImpl) Create(mailAddress string, hashedPassword string) (string, error) {
	query := "INSERT INTO accounts VALUES ($1, $2, $3)"
	id := generateId()
	if _, err := r.db.ExecContext(r.ctx, query, id, mailAddress, hashedPassword); err != nil {
		return "", err
	}
	return id, nil
}

func (r *accountRepositoryImpl) GetId(mailAddress string, hashedPassword string) (string, error) {
	query := "SELECT id FROM accounts WHERE mail_address = $1 AND hashed_password = $2 AND is_deleted=false"
	row := r.db.QueryRowContext(r.ctx, query, mailAddress, hashedPassword)
	var id string
	if err := row.Scan(&id); err != nil {
		return "", err
	}
	if id == "" {
		return "", errors.New(util.NotFoundMessage)
	}
	return id, nil
}

func (r *accountRepositoryImpl) UpdateById(id int64, column string, value string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	query := "UPDATE accounts SET $1 = $2 WHERE id = $3 AND is_deleted=false"
	if _, err := tx.ExecContext(r.ctx, query, column, value, id); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (r *accountRepositoryImpl) UpdateByMailAddress(mailAddress string, column string, value string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	query := "UPDATE accounts SET $1 = $2 WHERE id = $3 AND is_deleted=false"
	if _, err := tx.ExecContext(r.ctx, query, column, value, mailAddress); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func generateId() string {
	en := rand.New(rand.NewSource(time.Now().UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(time.Now()), en)
	return id.String()
}
