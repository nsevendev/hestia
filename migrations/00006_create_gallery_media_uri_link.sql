-- +goose Up
CREATE TABLE gallery_media_uri_links (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    uuid_gallery UUID NOT NULL,
    uuid_media_uri UUID NOT NULL UNIQUE,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW() NOT NULL,
    CONSTRAINT fk_gallery FOREIGN KEY (uuid_gallery) REFERENCES galleries(uuid) ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT fk_media_uri FOREIGN KEY (uuid_media_uri) REFERENCES media_uris(uuid) ON UPDATE CASCADE ON DELETE CASCADE
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS gallery_media_uri_links;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
