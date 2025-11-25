package converter

import (
	"github.com/goNiki/ReviewService/internal/domain"
	revV1 "github.com/goNiki/ReviewService/shared/pkg/openapi/reviewerservice/v1"
)

func UserToDTO(user *domain.User) revV1.OptUser {

	resp := revV1.User{
		UserID:   user.UserId,
		Username: user.Username,
		TeamName: user.TeamName,
		IsActive: user.IsActive,
	}

	return revV1.NewOptUser(resp)
}
