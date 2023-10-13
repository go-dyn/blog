CREATE TABLE IF NOT EXISTS comments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL DEFAULT 1,
    post_id INTEGER NOT NULL DEFAULT 1,
    title TEXT NOT NULL,
    body TEXT NOT NULL,
    status INTEGER NOT NULL DEFAULT 1,
    created_at INTEGER NOT NULL,
    updated_at INTEGER NOT NULL,
    deleted_at INTEGER NOT NULL DEFAULT 0
);

CREATE INDEX IF NOT EXISTS idx_comments_on_user_id ON comments (user_id DESC);
CREATE INDEX IF NOT EXISTS idx_comments_on_post_id ON comments (post_id DESC);
CREATE INDEX IF NOT EXISTS idx_comments_on_created_at ON comments (created_at DESC);