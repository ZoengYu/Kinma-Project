-- name: CreateUser :one
INSERT INTO users (
  username,
  hashed_password,
  email,
  phone
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users 
SET username = $2, hashed_password = $3, email = $4, phone = $5
WHERE username = $1
RETURNING *;

-- name: GetUserByMail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;