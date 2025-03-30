-- +goose Up
CREATE TABLE article_terms (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT NOT NULL,
    content TEXT,
    uuid_term UUID NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW() NOT NULL,
    CONSTRAINT fk_term FOREIGN KEY (uuid_term) REFERENCES terms(uuid) ON UPDATE CASCADE ON DELETE CASCADE
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS article_terms;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
