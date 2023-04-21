// Filename: internal/models/user_info.go
package models

import (
	"context"
	"database/sql"
	"time"
)

type User struct {
	UserID      int64
	Fname       string
	Lname       string
	Email       string
	Address     string
	PhoneNumber string
	Password    string
}

type UserModel struct {
	DB *sql.DB
}

// Code to access the database
func (m *UserModel) Get() (*User, error) {
	var u User

	statement := `
				SELECT id, fname, lname, email, address, phone_number, password
				FROM users_info
				LIMIT 1
				`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement).Scan(&u.UserID, &u.Fname, &u.Lname, &u.Address, &u.PhoneNumber, &u.Password)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
