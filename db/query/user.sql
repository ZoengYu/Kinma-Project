-- name: CreateUser :one
INSERT INTO users (
  username,
  hashed_password,
  email
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users 
SET username = $2, hashed_password = $3, email = $4
WHERE username = $1
RETURNING *;

-- name: GetUserByMail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;