package database

import (
	"context"

	"github.com/lib/pq"
	"github.com/ntnghiatn/financial-app-backend/internal/model"
	"github.com/pkg/errors"
)

var ErrUserExists = errors.New("user with that email exists")

type UsersDB interface {
	CreateUser(ctx context.Context, user *model.User) error
}

var createUserQuery = `
	INSERT INTO users (
		email, password_hash
	)
	VALUE (
		:email, :password_hash
	)
	RETURNING user_id
`

func (db database) CreateUser(ctx context.Context, user *model.User) error {
	rows, err := db.conn.NamedQueryContext(ctx, createUserQuery, user)
	if rows != nil {
		defer rows.Close()
	}

	if err != nil {
		if pqError, ok := err.(*pq.Error); ok {
			if pqError.Code.Name() == UniqueViolation {
				if pqError.Constraint == "user_email" {
					return ErrUserExists
				}
			}
		}
		return errors.Wrap(err, "could not create user")
	}
	rows.Next()
	if err := rows.Scan(&user.ID); err != nil {
		return errors.Wrap(err, "could not get created userID")
	}
	return nil
}
