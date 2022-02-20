-- name: CreateTransfer :one
INSERT INTO transfers (
  from_account_id,
  to_fundraise_id,
  amount
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers
WHERE 
from_account_id = $1 OR to_fundraise_id = $2 
LIMIT 1;

-- name: ListTransfers :many
SELECT * FROM transfers
WHERE 
    from_account_id = $1 OR
    to_fundraise_id = $2
ORDER BY id
LIMIT $3
OFFSET $4;

-- name: TransferSuccess :one
UPDATE transfers SET success = $2
WHERE id = $1
RETURNING *;
