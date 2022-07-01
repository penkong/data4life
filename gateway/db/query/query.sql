-- name: WriteToken :exec
INSERT INTO token (name) VALUES ($1) 
ON CONFLICT (name) 
DO UPDATE SET occur = token.occur + 1;

-- -- name: BashInsert :exec
-- INSERT INTO token(name, occur)(select * from unnest($1::int[], $2::int[]));


-- name: ReadNonUnique :many
SELECT name, occur FROM token WHERE occur > 1 ORDER BY occur DESC LIMIT $1 OFFSET $2;

-- name: CountNonUnique :one
SELECT COUNT(*) FROM token WHERE occur > 1;