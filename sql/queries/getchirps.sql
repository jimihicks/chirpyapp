-- name: GetAllChirps :many

SELECT * FROM chirps
ORDER BY created_at;

-- name: GetChirp :one

SELECT * FROM chirps
WHERE id = $1;