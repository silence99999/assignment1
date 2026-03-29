-- +goose Up
CREATE TABLE appointments (
                              id TEXT PRIMARY KEY,
                              title TEXT NOT NULL,
                              description TEXT,
                              doctor_id TEXT NOT NULL,
                              status TEXT NOT NULL,
                              created_at TIMESTAMP NOT NULL,
                              updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE appointments;