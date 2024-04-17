-- name: GetStore :one
SELECT
    *
FROM
    stores
WHERE
    id = ?
LIMIT
    1;

-- name: ListStores :many
SELECT
    *
FROM
    stores
ORDER By
    id;

-- name: ListStoresByUserId :many
SELECT
    *
FROM
    stores
WHERE
    userId = ?
ORDER By
    id;

-- name: CreateStore :execresult
INSERT INTO
    stores (title, description, userId)
VALUES
    (?, ?, ?);

-- name: UpdateStore :exec
UPDATE stores
SET
    title = ?,
    description = ?,
    userId = ?
WHERE
    id = ?;

-- name: DeleteStore :exec
DELETE FROM stores
WHERE
    id = ?;