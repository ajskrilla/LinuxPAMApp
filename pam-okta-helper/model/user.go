package model

import "time"

type User struct {
	ID           string    `db:"id"`
	Username     string    `db:"username"`
	PasswordHash string    `db:"password_hash"`
	SessionToken string    `db:"session_token"`
	IsSudoer     bool      `db:"is_sudoer"`
	LastLogin    time.Time `db:"last_login"`
}

