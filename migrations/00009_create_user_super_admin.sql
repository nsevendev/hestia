-- +goose Up
INSERT INTO users (uuid, email, hashed_password, username, created_at, updated_at)
VALUES (
    gen_random_uuid(),
    'superadmin@hestia.local',
    '$2a$12$oR09W2I/7Wpm8XNgDDRAGOrTriRM5MqoKInK5Y.gBqh5EkAT8PgTa',
    'superadmin',
    now(),
    now()
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DELETE FROM users WHERE email = 'superadmin@hestia.local';
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
