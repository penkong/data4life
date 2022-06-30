-- name: WriteToken :exec
INSERT INTO token (name) VALUES ($1) 
ON CONFLICT (name) 
DO UPDATE SET occur = token.occur + 1;