-- name: CreateProduct :one
INSERT INTO product (
	account_id,
  title,
  content,
	product_tag
) VALUES (
  $1, $2, $3,$4
) RETURNING *;

-- name: GetProduct :one
SELECT * FROM product
WHERE id = $1 LIMIT 1;

-- name: ListProduct :many
SELECT * FROM product
WHERE account_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateProductTitle :one
UPDATE product
SET title = $2
WHERE id = $1
RETURNING *;

-- name: UpdateProductContent :one
UPDATE product
SET content = $2
WHERE id = $1
RETURNING *;

-- name: UpdateProductTag :one
UPDATE product
SET product_tag = $2
WHERE id = $1
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM product
WHERE id = $1;