// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package sqlc

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type EmailVerificationRequest struct {
	ID        int32
	UserID    pgtype.UUID
	Email     string
	Code      string
	ExpiresAt pgtype.Timestamptz
}

type PasswordResetSession struct {
	ID                string
	UserID            pgtype.UUID
	Email             string
	Code              string
	ExpiresAt         pgtype.Timestamptz
	EmailVerified     bool
	TwoFactorVerified bool
}

type Session struct {
	ID                string
	UserID            pgtype.UUID
	ExpiresAt         pgtype.Timestamptz
	TwoFactorVerified bool
}

type User struct {
	ID            pgtype.UUID
	Email         string
	PasswordHash  string
	EmailVerified bool
	Registered2fa bool
	TotpKey       []byte
	RecoveryCode  []byte
	CreatedAt     pgtype.Timestamptz
}
