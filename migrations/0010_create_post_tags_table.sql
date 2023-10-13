CREATE TABLE IF NOT EXISTS post_tags (
    post_id INTEGER NOT NULL,
    tag_id INTEGER NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_post_tags_on_post_id ON post_tags (post_id DESC);
CREATE INDEX IF NOT EXISTS idx_post_tags_on_tag_id ON post_tags (tag_id DESC);