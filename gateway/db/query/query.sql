-- name: GetTodo :one
SELECT * FROM todo
WHERE id = $1 LIMIT 1;

