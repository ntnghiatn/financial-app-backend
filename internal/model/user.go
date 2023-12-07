package model

import (
	"time"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserId string

type User struct {
	ID           UserId     `json:"id,omitempty" db:"user_id"`
	Email        *string    `json:"email" db:"email"`
	PasswordHash *[]byte    `json:"-" db:"password_hash"`
	CreatedAt    *time.Time `json:"-" db:"created_at"`
	UpdatedAt    *time.Time `json:"-" db:"updated_at"`
	DeletedAt    *time.Time `json:"-" db:"deleted_at"`
}

// Verify all requests fields  before create or uptdate
func (u *User) Verify() error {
	if u.Email == nil || (u.Email != nil && len(*u.Email) == 0) {
		return errors.New("Email is required")
	}
	return nil
}

// set password updates a user's password
func (u *User) SetPassword(password string) error {
	hash, err := HashPassword(password)
	if err != nil {
		return err
	}
	u.PasswordHash = &hash

	return nil
}

// check password
func (u *User) CheckPassword(password string) error {
	if u.PasswordHash != nil && len(*u.PasswordHash) == 0 {
		return errors.New("password not set")
	}

	return bcrypt.CompareHashAndPassword(*u.PasswordHash, []byte(password))
}

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
