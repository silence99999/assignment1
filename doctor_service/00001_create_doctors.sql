-- +goose Up
CREATE TABLE doctors (
                         id TEXT PRIMARY KEY,
                         full_name TEXT NOT NULL,
                         specialization TEXT,
                         email TEXT NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE doctors;
