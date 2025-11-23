-- +goose Up
-- +goose StatementBegin
CREATE TABLE pr_reviewers (
    id SERIAL PRIMARY KEY,
    pr_id VARCHAR(50) NOT NULL REFERENCES pull_requests(pr_id) ON DELETE CASCADE,
    reviewer_id VARCHAR(50) NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    assigned_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(pr_id, reviewer_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS pr_reviewers CASCADE;
-- +goose StatementEnd
