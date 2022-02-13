-- name: CreateFundraise :one
INSERT INTO fundraise (
	product_id,
  target_amount,
  progress_amount
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetFundraise :one
SELECT * FROM fundraise
WHERE product_id = $1 LIMIT 1;

-- name: ExitFundraise :one
UPDATE fundraise
SET success = $2, end_date = $3
WHERE product_id = $1
RETURNING *;