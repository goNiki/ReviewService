package pullrequests

import (
	"context"
	"errors"

	"github.com/goNiki/ReviewService/internal/converter"
	"github.com/goNiki/ReviewService/models/errorapp"
	revV1 "github.com/goNiki/ReviewService/shared/pkg/openapi/reviewerservice/v1"
)

func (a *Api) PullRequestCreatePost(ctx context.Context, req *revV1.PullRequestCreatePostReq) (r revV1.PullRequestCreatePostRes, _ error) {
	const op = "PullRequestCreatePost"
	prWithReviewers, err := a.prService.CreatePullRequest(ctx, req.PullRequestID, req.PullRequestName, req.AuthorID)

	if err != nil {
		a.log.Error(ctx, op, err)
		switch {
		case errors.Is(err, errorapp.ErrUserNotFound):
			return &revV1.NotFoundError{
				Error: revV1.NotFoundErrorError{
					Code:    revV1.NotFoundErrorErrorCodeNOTFOUND,
					Message: "author not found",
				},
			}, nil
		case errors.Is(err, errorapp.ErrPRExists):
			return &revV1.ConflictError{
				Error: revV1.ConflictErrorError{
					Code:    revV1.ConflictErrorErrorCodePREXISTS,
					Message: "pr already exists",
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

	resp := converter.PrWithReviewersToDTO(prWithReviewers)

	return &revV1.PullRequestCreatePostCreated{
		Pr: revV1.NewOptPullRequest(resp),
	}, nil
}
