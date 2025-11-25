package pullrequests

import (
	"context"
	"errors"

	"github.com/goNiki/ReviewService/internal/converter"
	"github.com/goNiki/ReviewService/models/errorapp"
	revV1 "github.com/goNiki/ReviewService/shared/pkg/openapi/reviewerservice/v1"
)

func (a *Api) PullRequestMergePost(ctx context.Context, req *revV1.PullRequestMergePostReq) (r revV1.PullRequestMergePostRes, _ error) {
	const op = "PullRequestMergePost"

	prWithReviewers, err := a.prService.SetMergeInPR(ctx, req.PullRequestID)

	if err != nil {
		a.log.Error(ctx, op, err)

		switch {
		case errors.Is(err, errorapp.ErrPullRequestNotFound):
			return &revV1.NotFoundError{
				Error: revV1.NotFoundErrorError{
					Code:    revV1.NotFoundErrorErrorCodeNOTFOUND,
					Message: "pull reques not found",
				},
			}, nil
		default:
			return &revV1.InternalServerError{
				Error: revV1.InternalServerErrorError{
					Code:    revV1.InternalServerErrorErrorCodeINTERNALERROR,
					Message: "internal error",
				},
			}, nil
		}
	}

	resp := converter.PrWithReviewersToDTO(prWithReviewers)

	return &revV1.PullRequestMergePostOK{
		Pr: revV1.NewOptPullRequest(resp),
	}, nil
}
