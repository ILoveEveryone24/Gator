-- name: GetNextFeedToFetch :one
SELECT id, url FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST LIMIT 1;
