CREATE TABLE IF NOT EXISTS post_attachments (
    post_id INTEGER NOT NULL,
    attachment_id INTEGER NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_post_attachments_on_post_id ON post_attachments (post_id DESC);
CREATE INDEX IF NOT EXISTS idx_post_attachments_on_attachment_id ON post_attachments (attachment_id DESC);