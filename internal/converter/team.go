package converter

import (
	"github.com/goNiki/ReviewService/internal/domain"
	revV1 "github.com/goNiki/ReviewService/shared/pkg/openapi/reviewerservice/v1"
)

func TeamDtoToModel(req revV1.Team) *domain.TeamWithMembers {
	members := make([]*domain.User, 0, len(req.Members))
	for _, m := range req.Members {
		members = append(members, &domain.User{
			UserId:   m.UserID,
			Username: m.Username,
			IsActive: m.IsActive,
		})
	}

	return &domain.TeamWithMembers{
		Team: &domain.Team{
			Name: req.TeamName,
		},
		Members: members,
	}
}

func TeamModelToDto(teamWithMembers *domain.TeamWithMembers) revV1.Team {

	dto := revV1.Team{
		TeamName: teamWithMembers.Team.Name,
		Members:  make([]revV1.TeamMember, 0, len(teamWithMembers.Members)),
	}

	for _, m := range teamWithMembers.Members {
		dto.Members = append(dto.Members, revV1.TeamMember{
			UserID:   m.UserId,
			Username: m.Username,
			IsActive: m.IsActive,
		})
	}

	return dto
}

func ReviewerWithPRToDTO(dates *domain.ReviewerWithPR) *revV1.UsersGetReviewGetOK {
	return &revV1.UsersGetReviewGetOK{
		UserID:       dates.UserID,
		PullRequests: PullRequestsToDTO(dates.PullRequests),
	}

}

func PullRequestsToDTO(prs []*domain.PullRequest) []revV1.PullRequestShort {
	resp := make([]revV1.PullRequestShort, 0, len(prs))

	for _, pr := range prs {
		if pr != nil {
			resp = append(resp, PullRequestToDTO(pr))
		}
	}

	return resp
}

func PullRequestToDTO(pr *domain.PullRequest) revV1.PullRequestShort {
	if pr == nil {
		return revV1.PullRequestShort{}
	}

	return revV1.PullRequestShort{
		PullRequestID:   pr.Id,
		PullRequestName: pr.Name,
		AuthorID:        pr.AuthorID,
		Status:          StatusPRToDTO(pr.Status),
	}
}

func StatusPRToDTO(status string) revV1.PullRequestShortStatus {
	switch status {
	case domain.Open:
		return revV1.PullRequestShortStatusOPEN
	case domain.Merged:
		return revV1.PullRequestShortStatusMERGED
	default:
		return revV1.PullRequestShortStatusOPEN
	}
}
