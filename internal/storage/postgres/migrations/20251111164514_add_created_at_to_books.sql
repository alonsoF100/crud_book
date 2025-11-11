-- +goose Up
-- +goose StatementBegin
ALTER TABLE books ADD COLUMN created_at TIMESTAMP default NOW();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE books DROP COLUMN created_at;
-- +goose StatementEnd
