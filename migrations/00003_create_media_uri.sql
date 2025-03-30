-- +goose Up
CREATE TABLE media_uris (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    path TEXT NOT NULL,
    media_type TEXT NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW() NOT NULL
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS media_uris;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
