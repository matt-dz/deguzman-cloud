package session

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"deguzman-auth/internal/database"
	"deguzman-auth/internal/logger"
	"deguzman-auth/internal/sqlc"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

var db = database.GetDb()
var sqlcDb = sqlc.New(db)
var log = logger.GetLogger()

const sessionCookieName = "session"

var (
	ErrUnauthorized     = errors.New("unauthorized")
	ErrEmailNotVerified = errors.New("email not verified")
	ErrInvalidSession   = errors.New("invalid session")
	ErrExpiredSession   = errors.New("session expired")
)

type Session struct {
	id        uuid.UUID
	token     string
	expiresAt time.Time
}

type User struct {
	id uuid.UUID
}

func GenerateSessionId(token string) (string, error) {
	hashEncoder := sha256.New()
	_, err := hashEncoder.Write([]byte(token))
	if err != nil {
		return "", nil
	}
	return hex.EncodeToString(hashEncoder.Sum(nil)), nil
}

func within15Days(date time.Time) bool {
	return time.Now().Before(date.Add(time.Hour * 24 * 15))
}

func GenerateSessionToken() string {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	return base64.StdEncoding.EncodeToString(bytes)
}

func CreateSession(db *sqlc.Queries, token string, userId pgtype.UUID) (*sqlc.CreateSessionParams, error) {
	sessionId, err := GenerateSessionId(token)
	if err != nil {
		return nil, err
	}

	session := sqlc.CreateSessionParams{
		ID:     sessionId,
		UserID: userId,
		ExpiresAt: pgtype.Timestamptz{
			Time:  time.Now().Add(time.Hour * 24 * 30),
			Valid: true,
		},
	}
	if err := db.CreateSession(context.Background(), session); err != nil {
		return nil, err
	}

	return &session, nil
}

func ExtractSessionId(r *http.Request) (string, error) {
	sessionCookie, err := r.Cookie(sessionCookieName)
	if err != nil {
		return "", err
	}

	sessionId, err := GenerateSessionId(sessionCookie.Value)
	if err != nil {
		return "", err
	}

	return sessionId, nil

}

func RetriveUserFromSessionToken(r *http.Request) (*sqlc.GetUserSessionBySessionIdRow, error) {
	sessionCookie, err := r.Cookie(sessionCookieName)
	if err != nil {
		return nil, err
	}
	sessionId, err := GenerateSessionId(sessionCookie.Value)
	if err != nil {
		return nil, err
	}

	session, err := sqlcDb.GetUserSessionBySessionId(context.Background(), sessionId)
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func ValidateSessionToken(token string) (*sqlc.GetUserSessionBySessionIdRow, error) {
	sessionId, err := GenerateSessionId(token)
	if err != nil {
		return nil, err
	}

	session, err := sqlcDb.GetUserSessionBySessionId(context.Background(), sessionId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.Join(ErrInvalidSession, err)
		}
		return nil, err
	}

	if time.Now().After(session.Session.ExpiresAt.Time) {
		sqlcDb.DeleteUserSession(context.Background(), sessionId)
		return nil, ErrExpiredSession
	}

	if within15Days(session.Session.ExpiresAt.Time) {
		session.Session.ExpiresAt.Time = time.Now().Add(time.Hour * 24 * 30)
		err = sqlcDb.UpdateSessionExpireation(context.Background(), sqlc.UpdateSessionExpireationParams{
			ExpiresAt: session.Session.ExpiresAt,
			ID:        session.Session.ID,
		})
		if err != nil {
			log.Warn("Unable to extend session lifetime", slog.Any("error", err))
		}
	}

	return &session, nil
}

func InvalidateSessionToken(sessionId string) error {
	return sqlcDb.DeleteUserSession(context.Background(), sessionId)
}

func InvalidateAllSessions(userId pgtype.UUID) error {
	return sqlcDb.DeleteAllUserSessions(context.Background(), userId)
}

func SetSessionTokenCookie(w http.ResponseWriter, token string, expiresAt time.Time) {
	if os.Getenv("ENV") == "PROD" {
		http.SetCookie(w, &http.Cookie{
			Name:     sessionCookieName,
			Value:    token,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
			Expires:  expiresAt,
			Path:     "/",
			Secure:   true,
		})
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     sessionCookieName,
			Value:    token,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
			Expires:  expiresAt,
			Path:     "/",
		})

	}
}

func DeleteSessionTokenCookie(w http.ResponseWriter) {
	if os.Getenv("ENV") == "PROD" {
		http.SetCookie(w, &http.Cookie{
			Name:     sessionCookieName,
			Value:    "",
			SameSite: http.SameSiteLaxMode,
			MaxAge:   0,
			Path:     "/",
			Secure:   true,
		})
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     sessionCookieName,
			Value:    "",
			SameSite: http.SameSiteLaxMode,
			MaxAge:   0,
			Path:     "/",
		})
	}
}

func AuthorizeSession(w http.ResponseWriter, r *http.Request) (*sqlc.GetUserSessionBySessionIdRow, error) {
	/* Validate session cookie */
	log.Debug("Cookies")
	for _, cookie := range r.Cookies() {
		log.Debug("Cookie", slog.String("name", cookie.Name), slog.String("value", cookie.Value))
	}
	token, err := r.Cookie(sessionCookieName)
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			return nil, errors.Join(ErrUnauthorized, err)
		}
		return nil, err
	}

	sessionAndUser, err := ValidateSessionToken(token.Value)
	if err != nil {
		if errors.Is(err, ErrInvalidSession) || errors.Is(err, ErrExpiredSession) {
			return nil, errors.Join(ErrUnauthorized, err)
		}
		return nil, err
	}
	SetSessionTokenCookie(w, token.Value, sessionAndUser.Session.ExpiresAt.Time)

	return sessionAndUser, nil
}
