-- +goose Up
CREATE TABLE news (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT NOT NULL,
    content TEXT,
    published_at TIMESTAMP,
    uuid_media_image UUID,
    uuid_media_link UUID,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW() NOT NULL,
    CONSTRAINT fk_media_image FOREIGN KEY (uuid_media_image) REFERENCES media_uris(uuid) ON UPDATE SET NULL ON DELETE SET NULL,
    CONSTRAINT fk_media_link FOREIGN KEY (uuid_media_link) REFERENCES media_uris(uuid) ON UPDATE SET NULL ON DELETE SET NULL
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS news;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
