-- +goose Up
INSERT INTO galleries (uuid, title, created_at, updated_at)
SELECT gen_random_uuid(), 'Galerie principale', now(), now()
WHERE NOT EXISTS (SELECT 1 FROM galleries);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DELETE FROM galleries WHERE title = 'Galerie principale';
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
