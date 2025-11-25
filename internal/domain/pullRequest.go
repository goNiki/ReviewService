package domain

import "time"

const Merged = "MERGED"
const Open = "OPEN"

const MaxReviewers = 2

type PullRequest struct {
	Id        string
	Name      string
	AuthorID  string
	TeamId    int64
	Status    string
	CreatedAt time.Time
	MergedAt  *time.Time
}

type PullRequestWithReviewers struct {
	PR        *PullRequest
	Reviewers []string
}

type PRReviewer struct {
	PRID       string
	ReviewerID string
	AssignedAt time.Time
}

type ReviewerWithPR struct {
	UserID       string
	PullRequests []*PullRequest
}
