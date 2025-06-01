package repository

import (
	"context"
	"database/sql"
)

type ReserveRepository interface {
	Create(mailAddress string) (int64, error)
}

func NewReserveRepository(ctx context.Context, db *sql.DB) *reserveRepositoryImpl {
	return &reserveRepositoryImpl{ctx: ctx, db: db}
}

type reserveRepositoryImpl struct {
	ctx context.Context
	db  *sql.DB
}

func (r *reserveRepositoryImpl) Create(mailAddress string) (int64, error) {
	query := "INSERT INTO reserve VALUES ($1) RETURNING id"
	res, err := r.db.ExecContext(r.ctx, query, mailAddress)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
