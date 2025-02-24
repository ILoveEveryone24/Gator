-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = CURRENT_TIMESTAMP,
updated_at = CURRENT_TIMESTAMP
WHERE id = $1;
