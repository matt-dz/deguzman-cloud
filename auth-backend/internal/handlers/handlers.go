package handlers

import (
	"context"
	"crypto/subtle"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"net/mail"
	"time"

	"deguzman-auth/internal/argon2id"
	"deguzman-auth/internal/cors"
	"deguzman-auth/internal/database"
	"deguzman-auth/internal/email"
	"deguzman-auth/internal/logger"
	pswd "deguzman-auth/internal/password"
	"deguzman-auth/internal/session"
	"deguzman-auth/internal/sqlc"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/argon2"
)

const verificationCodeLength = 6

var db = database.GetDb()
var sqlcDb = sqlc.New(db)
var log = logger.GetLogger()

var pgError *pgconn.PgError

var expiredSessionError = errors.New("session expired")

func HandleHeartbeat(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func HandleCors(w http.ResponseWriter, r *http.Request) {
	cors.AddCors(w, r)
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	/* Parse body */
	var body HandleLoginBody
	if err := decodeJson(&body, r); err != nil {
		log.Error("Failed to decode body", slog.Any("error", err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	/* Retrieve password from db */
	log.DebugContext(r.Context(), "Retrieving password from db")
	passwordAndId, err := sqlcDb.GetPasswordAndId(ctx, body.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.ErrorContext(r.Context(), "Email not found", slog.String("email", body.Email))
			http.Error(w, "Incorrect email or password", http.StatusUnauthorized)
		} else {
			log.ErrorContext(r.Context(), "Failed to get password", slog.Any("error", err))
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	/* Compare password hashes */
	log.DebugContext(r.Context(), "Decoding argon2 paramters")
	argonParams, salt, truePasswordHash, err := argon2id.DecodeHash(passwordAndId.PasswordHash)
	if err != nil {
		log.ErrorContext(r.Context(), "Failed to decode password hash", slog.Any("error", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	log.DebugContext(r.Context(), "Hashing given password")
	givenPasswordHash := argon2.IDKey([]byte(body.Password), salt, argonParams.Iterations, argonParams.Memory, argonParams.Parallelism, argonParams.KeyLength)

	log.DebugContext(r.Context(), "Comparing passwords")
	if subtle.ConstantTimeCompare(truePasswordHash, []byte(givenPasswordHash)) == 0 {
		log.ErrorContext(r.Context(), "Incorrect password", slog.String("email", body.Email))
		http.Error(w, "Incorrect email or password", http.StatusUnauthorized)
		return
	}

	/* Create session */
	log.DebugContext(r.Context(), "Creating session")
	token := session.GenerateSessionToken()
	userSession, err := session.CreateSession(sqlcDb, token, passwordAndId.ID)
	if err != nil {
		log.ErrorContext(r.Context(), "Failed to create session", slog.Any("error", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	log.DebugContext(r.Context(), "Setting session token cookie")
	session.SetSessionTokenCookie(w, token, userSession.ExpiresAt.Time)

	/* Redirect */
	json.NewEncoder(w).Encode(HandleLoginResponse{
		Redirect: sanitizeRedirect(r.URL.Query().Get("redirect")),
	})
}

func HandleSignup(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	/* Parse body */
	var body HandleSignupBody
	if err := decodeJson(&body, r); err != nil {
		log.Error("Failed to decode body", slog.Any("error", err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	/* Validate email */
	log.DebugContext(r.Context(), "Validating email")
	if _, err := mail.ParseAddress(body.Email); err != nil {
		log.ErrorContext(r.Context(), "Invalid email", slog.String("email", body.Email))
		http.Error(w, "Invalid email", http.StatusBadRequest)
		return
	}

	log.DebugContext(r.Context(), "Validating password")
	if err := pswd.ValidatePassword(body.Password); err != nil {
		log.ErrorContext(r.Context(), "Password does not meet requirements", slog.Any("error", err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.DebugContext(r.Context(), "Hashing Password")
	passwordHash, err := argon2id.EncodeHash(body.Password, argon2id.DefaultParams)
	if err != nil {
		log.ErrorContext(r.Context(), "Failed to hash password", slog.Any("error", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	/* Create user */
	log.DebugContext(r.Context(), "Creating user")
	userId, err := sqlcDb.CreateUser(ctx, sqlc.CreateUserParams{
		Email:        body.Email,
		PasswordHash: passwordHash,
	})
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" {
				log.ErrorContext(r.Context(), "Email already exists", slog.String("email", body.Email))
				http.Error(w, "Email already exists", http.StatusBadRequest)
				return
			}
		}
		log.ErrorContext(r.Context(), "Failed to create user", slog.Any("error", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	/* Create session */
	log.DebugContext(r.Context(), "Creating session")
	token := session.GenerateSessionToken()
	userSession, err := session.CreateSession(sqlcDb, token, userId)
	if err != nil {
		log.ErrorContext(r.Context(), "Failed to create session", slog.Any("error", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	log.DebugContext(r.Context(), "Setting session token cookie")
	session.SetSessionTokenCookie(w, token, userSession.ExpiresAt.Time)
}

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	log.DebugContext(r.Context(), "Invalidating session")
	session.DeleteSessionTokenCookie(w)
}

func HandleSessionValidation(w http.ResponseWriter, r *http.Request) {
	if verified, ok := r.Context().Value("emailVerified").(bool); ok {
		if !verified {
			log.ErrorContext(r.Context(), "Email not verified")
			http.Error(w, "Email not verified", http.StatusUnauthorized)
		}
		return
	}
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func HandleEmailVerification(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	/* Parse form */
	var body HandleEmailVerificationBody
	if err := decodeJson(&body, r); err != nil {
		log.Error("Failed to decode body", slog.Any("error", err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	/* Retrieve email from session */
	log.DebugContext(r.Context(), "Retrieving email from session token")
	user, err := session.RetriveUserFromSessionToken(r)
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			log.ErrorContext(r.Context(), "Session cookie not found", slog.Any("error", err))
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		} else if errors.Is(err, pgx.ErrNoRows) {
			log.ErrorContext(r.Context(), "Session not found", slog.Any("error", err))
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		} else {
			log.ErrorContext(r.Context(), "Failed to extract email from session token", slog.Any("error", err))
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}

	/* Retrieve verification code from db */
	log.DebugContext(r.Context(), "Retrieving verification code from db")
	sessionId, err := session.ExtractSessionId(r)
	if err != nil {
		log.ErrorContext(r.Context(), "Failed to extract session id", slog.Any("error", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	verificationRequestId, err := sqlcDb.GetEmailVerificationRequest(ctx, sqlc.GetEmailVerificationRequestParams{
		ID:   sessionId,
		Code: body.Code,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.ErrorContext(r.Context(), "Verification code not found", slog.String("code", body.Code))
			http.Error(w, "Invalid validation code", http.StatusUnauthorized)
		} else {
			log.ErrorContext(r.Context(), "Failed to get verification code", slog.Any("error", err))
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	log.DebugContext(r.Context(), "Verification code found!")

	/* Begin transaction */
	tx, err := db.Begin(ctx)
	if err != nil {
		log.ErrorContext(r.Context(), "Failed to begin transaction", slog.Any("error", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer tx.Rollback(ctx)
	qtx := sqlcDb.WithTx(tx)

	/* Delete verification request */
	log.DebugContext(r.Context(), "Deleting verification request")
	if err := qtx.DeleteEmailVerificationRequest(ctx, verificationRequestId); err != nil {
		log.ErrorContext(r.Context(), "Failed to delete verification request", slog.Any("error", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	/* Update email */
	log.DebugContext(r.Context(), "Updating email")
	err = qtx.UpdateEmailVerification(ctx, sqlc.UpdateEmailVerificationParams{
		EmailVerified: true,
		ID:            user.ID,
	})
	if err != nil {
		log.ErrorContext(r.Context(), "Failed to update email", slog.Any("error", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	/* End transaction */
	if err := tx.Commit(ctx); err != nil {
		log.ErrorContext(r.Context(), "Failed to commit transaction", slog.Any("error", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

// TODO: Check if code has been verified in last 10 minutes
func HandleEmailVerificationSend(w http.ResponseWriter, r *http.Request) {
	/* Retrieve recipient email */
	log.DebugContext(r.Context(), "Retrieving email from session token")
	user, err := session.RetriveUserFromSessionToken(r)
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			log.ErrorContext(r.Context(), "Session cookie not found", slog.Any("error", err))
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		} else if errors.Is(err, pgx.ErrNoRows) {
			log.ErrorContext(r.Context(), "Session not found", slog.Any("error", err))
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		} else {
			log.ErrorContext(r.Context(), "Failed to extract email from session token", slog.Any("error", err))
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}

	if user.EmailVerified {
		log.ErrorContext(r.Context(), "Email already verified", slog.String("email", user.Email))
		http.Error(w, "Email already verified", http.StatusBadRequest)
		return
	}

	/* Generate verification code */
	log.DebugContext(r.Context(), "Generating verification code")
	code, err := email.GenerateVerificationCode(verificationCodeLength)
	if err != nil {
		log.ErrorContext(r.Context(), "Failed to generate verification code", slog.Any("error", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	/* Insert verification code in DB */
	log.DebugContext(context.Background(), "Inserting verification code in DB")
	err = sqlcDb.CreateEmailVerificationRequest(context.Background(), sqlc.CreateEmailVerificationRequestParams{
		UserID: user.ID,
		Email:  user.Email,
		Code:   code,
		ExpiresAt: pgtype.Timestamptz{
			Time:  time.Now().Add(time.Minute * 10),
			Valid: true,
		},
	})
	if err != nil {
		log.Error("Failed to insert verification code in DB", slog.Any("error", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	/* Send verification code */
	log.DebugContext(r.Context(), "Sending verification email")
	if err := email.SendVerficationCode(user.Email, code); err != nil {
		if errors.Is(err, email.ErrCreateMail) {
			log.ErrorContext(r.Context(), "Failed to create email", slog.Any("error", err))
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		} else {
			log.ErrorContext(r.Context(), "Failed to send email", slog.Any("error", err))
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}
