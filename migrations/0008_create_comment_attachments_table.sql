CREATE TABLE IF NOT EXISTS comment_attachments (
    comment_id INTEGER NOT NULL,
    attachment_id INTEGER NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_comment_attachments_on_comment_id ON comment_attachments (comment_id DESC);
CREATE INDEX IF NOT EXISTS idx_comment_attachments_on_attachment_id ON comment_attachments (attachment_id DESC);