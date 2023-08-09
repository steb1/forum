package lib

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gofrs/uuid"
)

var sessions sync.Map

type Session struct {
	UserID   string    `json:"user_id"`
	Username string    `json:"username"`
	ExpireAt time.Time `json:"exp"`
}

func (s Session) isExpired() bool {
	return s.ExpireAt.Before(time.Now())
}

func ValidSession(req *http.Request) bool {
	cookie, err := req.Cookie("auth_session")
	if err == nil {
		if _, ok := sessions.Load(cookie.Value); ok {
			return ok
		}
	}
	return false
}

func NewSessionToken(res http.ResponseWriter, UserID, Username string) {
	sessionToken, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("❌ Failed to generate UUID: %v", err)
	}
	deleteSessionIfExist(Username)
	ExpireAt := time.Now().Add(2 * time.Hour)
	sessions.Store(sessionToken.String(), Session{UserID, Username, ExpireAt})
	http.SetCookie(res, &http.Cookie{
		Name:     "auth_session",
		Value:    sessionToken.String(),
		HttpOnly: true,
		Expires:  ExpireAt,
	})
}

func deleteSessionIfExist(username string) {
	sessions.Range(func(key, value interface{}) bool {
		if username == value.(Session).Username {
			sessions.Delete(key)
		}
		return true
	})
}

func DeleteExpiredSessions() {
	for {
		sessions.Range(func(key, value interface{}) bool {
			if value.(Session).isExpired() {
				sessions.Delete(key)
			}
			return true
		})
		time.Sleep(10 * time.Second)
	}
}
