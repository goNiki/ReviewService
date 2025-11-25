package converter

import (
	"github.com/goNiki/ReviewService/internal/domain"
)

func SliseUserToMapa(members []*domain.User) map[string]*domain.User {
	resp := make(map[string]*domain.User, len(members))

	for _, member := range members {
		resp[member.UserId] = member
	}
	return resp
}
