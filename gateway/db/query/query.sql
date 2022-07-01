-- name: WriteToken :exec
INSERT INTO token (name) VALUES ($1) 
ON CONFLICT (name) 
DO UPDATE SET occur = token.occur + 1;

-- name: BashInsert :exec
INSERT INTO token(name, occur)(select * from unnest($1::int[], $2::int[]));