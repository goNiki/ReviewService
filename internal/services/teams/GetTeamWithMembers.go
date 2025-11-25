package teams

import (
	"context"
	"fmt"

	"github.com/goNiki/ReviewService/internal/domain"
	"github.com/goNiki/ReviewService/models/errorapp"
)

func (s *Service) GetTeamWithMembers(ctx context.Context, teamName string) (*domain.TeamWithMembers, error) {
	team, err := s.teamRepo.GetTeamByName(ctx, nil, teamName)
	if err != nil {
		return nil, fmt.Errorf("%w: failed to get team: %v", errorapp.ErrDatabaseQuery, err)
	}

	if team == nil {
		return nil, errorapp.ErrTeamNotFound
	}

	members, err := s.userRepo.GetUsersByTeamID(ctx, nil, team.ID)
	if err != nil {
		return nil, fmt.Errorf("%w: failed to get members: %v", errorapp.ErrDatabaseQuery, err)
	}

	return &domain.TeamWithMembers{
		Team:    team,
		Members: members,
	}, nil

}
