
-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByUserName :one
SELECT * FROM users
WHERE user_name = $1 LIMIT 1;

-- name: GetUserByUserCardTier :many
SELECT * FROM users
WHERE card_tier = $1;

-- name: GetUserByUserEmail :many
SELECT * FROM users
WHERE email = $1; 

-- name: GetUserByUserContact :many
SELECT * FROM users
WHERE contact = $1;


-- name: ListUsers :many
SELECT * FROM users
ORDER BY user_name;

-- name: CreateUser :one
INSERT INTO users(
  full_name,credit_balance,email,contact,
  password,user_name,card_tier
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: DecreBalance :exec
UPDATE users 
SET credit_balance = credit_balance-$1 
WHERE id=$2 and credit_balance>=$1;

-- name: IncrBalance :exec
UPDATE users 
SET credit_balance = credit_balance+$1 
WHERE id=$2;



-- name: UpdateUser :exec
UPDATE users 
SET full_name = COALESCE($1,full_name),
credit_balance = COALESCE($2,credit_balance),
email = COALESCE($3,email),
contact = COALESCE($4,contact),
password = COALESCE($5,password),
user_name = COALESCE($6,user_name),
card_tier = COALESCE($7,card_tier)
WHERE id = $8;


-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
