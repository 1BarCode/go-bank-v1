-- name: CreateUser :one
INSERT INTO user (
    username,
    email,
    hashed_password,
    first_name,
    last_name
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetUser :one
SELECT * FROM user
WHERE username = $1 LIMIT 1;
