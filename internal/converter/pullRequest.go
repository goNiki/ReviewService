package converter

import (
	"github.com/goNiki/ReviewService/internal/domain"
	revV1 "github.com/goNiki/ReviewService/shared/pkg/openapi/reviewerservice/v1"
)

func PrWithReviewersToDTO(prWithReviewers *domain.PullRequestWithReviewers) revV1.PullRequest {
	var mergedAt revV1.OptNilDateTime
	if prWithReviewers.PR.MergedAt != nil {
		mergedAt = revV1.NewOptNilDateTime(*prWithReviewers.PR.MergedAt)
	}

	return revV1.PullRequest{
		PullRequestID:     prWithReviewers.PR.Id,
		PullRequestName:   prWithReviewers.PR.Name,
		AuthorID:          prWithReviewers.PR.AuthorID,
		Status:            (PullRequestStatusToDTO(prWithReviewers.PR.Status)),
		AssignedReviewers: prWithReviewers.Reviewers,
		CreatedAt:         revV1.NewOptNilDateTime(prWithReviewers.PR.CreatedAt),
		MergedAt:          mergedAt,
	}

}

func PullRequestStatusToDTO(status string) revV1.PullRequestStatus {
	switch status {
	case domain.Open:
		return revV1.PullRequestStatusOPEN
	case domain.Merged:
		return revV1.PullRequestStatusMERGED
	default:
		return revV1.PullRequestStatusOPEN
	}
}
