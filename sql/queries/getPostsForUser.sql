-- name: GetPostsForUser :many
SELECT posts.* FROM posts
JOIN feed_follows 
ON feed_follows.feed_id = posts.feed_id
WHERE feed_follows.user_id = $1
ORDER BY published_at DESC
LIMIT $2;
