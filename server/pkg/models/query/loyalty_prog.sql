
-- name: GetLoyaltyByID :one
SELECT * FROM loyalty_program
WHERE id = $1 LIMIT 1;

-- name: GetLoyaltyByName :one
SELECT * FROM loyalty_program
WHERE name = $1 LIMIT 1;

-- name: ListLoyalty :many
SELECT * FROM loyalty_program
ORDER BY name;

-- name: CreateLoyalty :one
INSERT INTO loyalty_program(
  name, currency_name,processing_time,description,enrollment_link,
  terms_condition_link,format_regex,partner_code,initial_earn_rate
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

-- name: GetRegEx :one
SELECT format_regex FROM loyalty_program
WHere id = $1;

-- name: UpdateLoyalty :one
UPDATE loyalty_program 
SET name = COALESCE($1,name),
currency_name = COALESCE($2,currency_name),
processing_time = COALESCE($3,processing_time),
description = COALESCE($4,description),
enrollment_link = COALESCE($5,enrollment_link),
terms_condition_link = COALESCE($6,terms_condition_link),
format_regex = COALESCE($7,format_regex),
partner_code = COALESCE($8,partner_code),
initial_earn_rate = COALESCE($9,initial_earn_rate)
WHERE id = $10
RETURNING *;


-- name: DeleteLoyalty :exec
DELETE FROM loyalty_program
WHERE id = $1;
