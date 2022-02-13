-- name: CreateAccount :one
INSERT INTO account (
  owner,
  currency
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetAccount :one
SELECT * FROM account
WHERE id = $1 LIMIT 1;

-- name: ListAccount :many
SELECT * FROM account
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateAccount :one
UPDATE account SET currency =$2
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM account
WHERE id = $1;