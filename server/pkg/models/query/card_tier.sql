-- name: GetCardTierByID :one
SELECT * FROM card_tier
WHERE id = $1 LIMIT 1;

-- name: GetCardTierByName :one
SELECT * FROM card_tier
WHERE name = $1 LIMIT 1;

-- name: ListCardTier :many
SELECT * FROM card_tier
ORDER BY name;

-- name: UpdateCardTier :exec
UPDATE card_tier 
SET name = COALESCE($1,name)
WHERE id = $2;

-- name: CreateCardTier :one
INSERT INTO card_tier(
  name,tier
) VALUES (
  $1, $2 
)
RETURNING *;

-- name: DeleteCardTier :exec
DELETE FROM card_tier
WHERE id = $1;
