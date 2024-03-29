-- name: CreateFundraise :one
INSERT INTO fundraise (
	product_id,
  target_amount,
  progress_amount
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetProductFundraise :one
SELECT * FROM fundraise
WHERE product_id = $1 LIMIT 1;

-- name: UpdateFundraiseProgressAmount :one
UPDATE fundraise
SET progress_amount = $2
WHERE product_id = $1
RETURNING *;

-- name: AddFundraiseProgressAmount :one
UPDATE fundraise
SET progress_amount = progress_amount + sqlc.arg(amount)
WHERE product_id = sqlc.arg(id)
RETURNING *;

-- name: ExitFundraise :one
UPDATE fundraise
SET success = $2, end_date = now()
WHERE product_id = $1
RETURNING *;