-- name: CreateProduct :one
INSERT INTO product (
	account_id,
  title,
  content,
	product_tag
) VALUES (
  $1, $2, $3,$4
) RETURNING *;

-- name: GetAccountProduct :one
SELECT * FROM product
WHERE id = $1 LIMIT 1;

-- name: ListMyProduct :many
SELECT * FROM product
WHERE account_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateProductDetail :one
UPDATE product
SET title = $2, content = $3, product_tag = $4, last_update = now()
WHERE id = $1
RETURNING *;

-- name: DeleteAccountProduct :exec
DELETE FROM product
WHERE id = $1;