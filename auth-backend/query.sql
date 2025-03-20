-- name: CreateSession :exec
INSERT INTO sessions (
    id, user_id, expires_at
) VALUES (
    $1, $2, $3
);

-- name: GetUserSessionBySessionId :one
SELECT
    sqlc.embed(sessions),
    users.id,
    users.email,
    users.email_verified,
    users.registered_2fa
FROM sessions INNER JOIN users
    ON users.id = sessions.user_id
        WHERE sessions.id = $1;

-- name: DeleteUserSession :exec
DELETE FROM sessions WHERE id = $1;

-- name: UpdateSessionExpireation :exec
UPDATE sessions SET expires_at = $1 WHERE id = $2;

-- name: DeleteAllUserSessions :exec
DELETE FROM sessions WHERE user_id = $1;

-- name: CreateUser :one
INSERT INTO users (email, password_hash)
    VALUES ($1, $2)
    RETURNING id;

-- name: GetPasswordAndId :one
SELECT id, password_hash FROM users WHERE email = $1;

-- name: CreateEmailVerificationRequest :exec
INSERT INTO email_verification_request (
    user_id, email, code, expires_at
) VALUES (
    $1, $2, $3, $4
);

-- name: GetEmailVerificationRequest :one
SELECT evr.id FROM email_verification_request AS evr INNER JOIN sessions AS s
    ON evr.user_id = s.user_id
        WHERE s.id = $1 AND evr.code = $2 AND evr.expires_at > now ();


-- name: DeleteEmailVerificationRequest :exec
DELETE FROM email_verification_request WHERE id = $1;

-- name: UpdateEmailVerification :exec
UPDATE users SET email_verified = $1 WHERE id = $2;
