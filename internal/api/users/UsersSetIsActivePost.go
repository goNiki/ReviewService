package users

import (
	"context"
	"errors"

	"github.com/goNiki/ReviewService/internal/converter"
	"github.com/goNiki/ReviewService/models/errorapp"
	revV1 "github.com/goNiki/ReviewService/shared/pkg/openapi/reviewerservice/v1"
)

func (a *Api) UsersSetIsActivePost(ctx context.Context, req *revV1.UsersSetIsActivePostReq) (r revV1.UsersSetIsActivePostRes, _ error) {
	const op = "UsersSetIsActivePost"

	user, err := a.userService.SetIsActive(ctx, req.UserID, req.IsActive)

	if err != nil {
		a.log.Error(ctx, op, err)
		switch {
		case errors.Is(err, errorapp.ErrUserNotFound):
			return &revV1.NotFoundError{
				Error: revV1.NotFoundErrorError{
					Code:    revV1.NotFoundErrorErrorCodeNOTFOUND,
					Message: "user not found",
				},
			}, nil
		default:
			return &revV1.InternalServerError{
				Error: revV1.InternalServerErrorError{
					Code:    revV1.InternalServerErrorErrorCodeINTERNALERROR,
					Message: "internal server error",
				},
			}, nil
		}
	}

	resp := &revV1.UsersSetIsActivePostOK{
		User: converter.UserToDTO(user),
	}

	return resp, nil
}
