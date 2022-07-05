-- name: GetLoyaltyMembershipByID :one
SELECT * FROM loyalty_membership
WHERE id = $1 LIMIT 1;

-- name: GetLoyaltyMembershipByProgram :many
SELECT * FROM loyalty_membership
WHERE program = $1; 

-- name: ListLoyaltyMembership :many
SELECT * FROM loyalty_membership
ORDER BY program;

-- name: CreateLoyaltyMembership :one
INSERT INTO loyalty_membership(
  program,name
) VALUES (
  $1, $2 
)
RETURNING *;

-- name: UpdateLoyaltyMembership :exec
UPDATE loyalty_membership 
SET name = COALESCE($1,name),
program = COALESCE($3,program)
WHERE id = $2;

-- name: DeleteLoyaltyMembership :exec
DELETE FROM loyalty_membership
WHERE id = $1;
