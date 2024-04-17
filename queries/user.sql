-- name: GetUser :one
SELECT
    *
FROM
    users
WHERE
    id = ?
LIMIT
    1;

-- name: ListUsers :many
SELECT
    *
FROM
    users
ORDER By
    id;

-- name: CreateUser :execresult
INSERT INTO
    users (email, password)
VALUES
    (?, ?);

-- name: UpdateUser :exec
UPDATE users
SET
    email = ?,
    password = ?
WHERE
    id = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE
    id = ?;