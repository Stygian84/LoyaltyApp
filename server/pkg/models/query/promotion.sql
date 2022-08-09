
-- name: GetPromotionByID :one
SELECT * FROM promotion
WHERE id = $1 LIMIT 1;

-- name: GetPromotionByProg :many
SELECT * FROM promotion
WHERE program = $1; 

-- name: GetPromotionByDateRange :many
SELECT * FROM promotion
WHERE $1 >= start_date AND $1 <= end_date AND program=$2;

-- name: ListPromotion :many
SELECT * FROM promotion LEFT JOIN loyalty_program ON promotion.program= loyalty_program.id;

-- name: CreatePromotion :one
INSERT INTO promotion(
  program,promo_type,start_date,end_date,
  earn_rate_type,constant, card_tier, loyalty_membership
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;


-- name: UpdatePromotion :exec
UPDATE promotion 
SET program = COALESCE($1,program),
promo_type = COALESCE($2,promo_type),
start_date = COALESCE($3,start_date),
end_date = COALESCE($4,end_date),
earn_rate_type = COALESCE($5,earn_rate_type),
constant = COALESCE($6,constant),
card_tier = COALESCE($7,card_tier),
loyalty_membership = COALESCE($8,loyalty_membership)
WHERE id = $9;


-- name: DeletePromotion :exec
DELETE FROM promotion
WHERE id = $1;
