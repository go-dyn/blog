CREATE TABLE IF NOT EXISTS post_categories (
    post_id INTEGER NOT NULL,
    category_id INTEGER NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_post_categories_on_post_id ON post_categories (post_id DESC);
CREATE INDEX IF NOT EXISTS idx_post_categories_on_category_id ON post_categories (category_id DESC);