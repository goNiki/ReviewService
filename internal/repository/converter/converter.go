package repoconverter

import (
	"github.com/goNiki/ReviewService/internal/domain"
	repomodels "github.com/goNiki/ReviewService/internal/repository/models"
)

func UserRepoToModel(user repomodels.User) *domain.User {
	return &domain.User{
		UserId:   user.UserID,
		Username: user.UserName,
		TeamID:   user.TeamID,
		IsActive: user.IsActive,
	}
}

func UserWithTeamNameToModel(user repomodels.UserWithTeamName) *domain.User {
	return &domain.User{
		UserId:   user.UserID,
		Username: user.UserName,
		TeamName: user.TeamName,
		IsActive: user.IsActive,
	}

}

func UsersRepoToModel(users []repomodels.User) []*domain.User {
	resp := make([]*domain.User, 0, len(users))

	for _, user := range users {
		resp = append(resp, UserRepoToModel(user))
	}
	return resp
}

func TeamRepoTomModels(team repomodels.Team) *domain.Team {
	return &domain.Team{
		ID:   team.Id,
		Name: team.Name,
	}
}
