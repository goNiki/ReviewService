-- +goose Up
-- +goose StatementBegin
CREATE TABLE teams (
    id SERIAL PRIMARY KEY, 
    team_name VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS teams;
-- +goose StatementEnd
