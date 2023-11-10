-- name: GetCreditRequestByID :one
SELECT * FROM credit_request
WHERE reference_number = $1 LIMIT 1;

-- name: GetCreditRequestByUser :many
SELECT * FROM credit_request
WHERE user_id = $1; 

-- name: GetCreditRequestByProg :many
SELECT * FROM credit_request
WHERE program = $1; 

-- name: GetCreditRequestByPromo :many
SELECT * FROM credit_request
WHERE program = $1 AND promo_used=$2 ; 

-- name: ListCreditRequestByStatus :many
SELECT * FROM credit_request
WHERE transaction_status = $1
ORDER BY program;

-- name: ListCreditRequest :many
SELECT * FROM credit_request
ORDER BY program;

-- name: CreateCreditRequest :one
INSERT INTO credit_request(
  user_id,promo_used,program,member_id,transaction_time,
  credit_used,reward_should_receive,transaction_status
) VALUES (
  $1, $2, $3, $4, $5, $6,$7,$8
)
RETURNING *;

-- name: UpdateTransactionStatusByID :exec
UPDATE credit_request
SET transaction_status = $1
WHERE reference_number = $2 ;

-- name: UpdateCreditRequest :exec
UPDATE credit_request 
SET user_id = COALESCE($1,user_id),
program = COALESCE($2,program),
member_id = COALESCE($3,member_id),
transaction_time = COALESCE($4,transaction_time),
credit_used = COALESCE($5,credit_used),
reward_should_receive = COALESCE($6,reward_should_receive),
transaction_status = COALESCE($7,transaction_status),
promo_used = COALESCE($8,promo_used)
WHERE reference_number = $9;

-- name: DeleteCreditRequest :exec
DELETE FROM credit_request
WHERE reference_number = $1;
