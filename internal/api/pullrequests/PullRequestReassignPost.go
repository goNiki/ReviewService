package pullrequests

import (
	"context"
	"errors"

	"github.com/goNiki/ReviewService/internal/converter"
	"github.com/goNiki/ReviewService/models/errorapp"
	revV1 "github.com/goNiki/ReviewService/shared/pkg/openapi/reviewerservice/v1"
)

func (a *Api) PullRequestReassignPost(ctx context.Context, req *revV1.PullRequestReassignPostReq) (r revV1.PullRequestReassignPostRes, _ error) {
	const op = "PullRequestReassignPost"

	prWithReviewers, newReviewer, err := a.prService.ReassignReviewer(ctx, req.PullRequestID, req.OldUserID)

	if err != nil {
		a.log.Error(ctx, op, err)

		switch {
		case errors.Is(err, errorapp.ErrPullRequestNotFound):
			return &revV1.NotFoundError{
				Error: revV1.NotFoundErrorError{
					Code:    revV1.NotFoundErrorErrorCodeNOTFOUND,
					Message: "pull request not found",
				},
			}, nil
		case errors.Is(err, errorapp.ErrPRAlreadyMerged):
			return &revV1.ConflictError{
				Error: revV1.ConflictErrorError{
					Code:    revV1.ConflictErrorErrorCodePRMERGED,
					Message: "pull request already merged",
				},
			}, nil
		case errors.Is(err, errorapp.ErrInvalidReviewer):
			return &revV1.ConflictError{
				Error: revV1.ConflictErrorError{
					Code:    revV1.ConflictErrorErrorCodeNOTASSIGNED,
					Message: "reviewer is not assigned to this PR",
				},
			}, nil
		case errors.Is(err, errorapp.ErrUserNotFound):
			return &revV1.NotFoundError{
				Error: revV1.NotFoundErrorError{
					Code:    revV1.NotFoundErrorErrorCodeNOTFOUND,
					Message: "user not found",
				},
			}, nil
		case errors.Is(err, errorapp.ErrNoCandidate):
			return &revV1.ConflictError{
				Error: revV1.ConflictErrorError{
					Code:    revV1.ConflictErrorErrorCodeNOCANDIDATE,
					Message: "no active replacement candidate in team",
				},
			}, nil
		default:
			return &revV1.InternalServerError{
				Error: revV1.InternalServerErrorError{
					Code:    revV1.InternalServerErrorErrorCodeINTERNALERROR,
					Message: "internal service error",
				},
			}, nil
		}
	}

	Pr := converter.PrWithReviewersToDTO(prWithReviewers)

	return &revV1.PullRequestReassignPostOK{
		Pr:         Pr,
		ReplacedBy: newReviewer,
	}, nil
}
