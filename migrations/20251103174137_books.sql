-- +goose Up
-- +goose StatementBegin
CREATE TABLE books (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    description TEXT,
    rating FLOAT DEFAULT 0 CHECK (rating >= 0 AND rating <= 5),
    status TEXT DEFAULT 'want' CHECK (status IN ('want', 'reading', 'finished'))
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS books;
-- +goose StatementEnd
