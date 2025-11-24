-- +goose Up
-- +goose StatementBegin
CREATE TYPE pr_status AS ENUM ('OPEN', 'MERGED');
CREATE TABLE pull_requests (
    pr_id VARCHAR(50) PRIMARY KEY,
    pr_name VARCHAR(255) NOT NULL,
    author_id VARCHAR(50) NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    team_id INTEGER NOT NULL REFERENCES teams(id) ON DELETE CASCADE,
    status pr_status NOT NULL, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    merged_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE TABLE IF EXISTS pull_requests;
DELETE TYPE IF EXISTS pr_status;
-- +goose StatementEnd
