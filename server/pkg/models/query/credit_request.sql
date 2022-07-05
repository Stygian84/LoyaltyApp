-- name: GetCreditRequestByID :one
SELECT * FROM credit_request
WHERE reference_number = $1 LIMIT 1;

-- name: GetCreditRequestByUser :many
SELECT * FROM credit_request
WHERE user_id = $1; 

-- name: GetCreditRequestByProg :many
SELECT * FROM credit_request
WHERE program = $1; 

-- name: ListCreditRequest :many
SELECT * FROM credit_request
ORDER BY program;

-- name: CreateCreditRequest :one
INSERT INTO credit_request(
  user_id,program,member_id,transaction_time,
  amount,transaction_status
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: UpdateCreditRequest :exec
UPDATE credit_request 
SET user_id = COALESCE($1,user_id),
program = COALESCE($2,program),
member_id = COALESCE($3,member_id),
transaction_time = COALESCE($4,transaction_time),
amount = COALESCE($5,amount),
transaction_status = COALESCE($6,transaction_status)
WHERE reference_number = $7;

-- name: DeleteCreditRequest :exec
DELETE FROM credit_request
WHERE reference_number = $1;
