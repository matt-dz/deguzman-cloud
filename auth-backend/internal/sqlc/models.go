// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package sqlc

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

func (e *Role) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Role(s)
	case string:
		*e = Role(s)
	default:
		return fmt.Errorf("unsupported scan type for Role: %T", src)
	}
	return nil
}

type NullRole struct {
	Role  Role
	Valid bool // Valid is true if Role is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullRole) Scan(value interface{}) error {
	if value == nil {
		ns.Role, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Role.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Role), nil
}

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
	FirstName     string
	LastName      string
	Email         string
	PasswordHash  string
	EmailVerified bool
	CreatedAt     pgtype.Timestamptz
	Roles         []Role
}
