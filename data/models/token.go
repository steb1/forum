package models

import "time"

const tokenExpiration = time.Hour

type Token struct {
    UserID string `json:"user_id"`
	Username string `json:"username"`
    ExpiresAt int64 `json:"exp"`

}