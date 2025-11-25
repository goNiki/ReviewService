package teams

import (
	"context"
	"fmt"

	"github.com/goNiki/ReviewService/internal/domain"
	"github.com/goNiki/ReviewService/internal/infrastructure/database"
	"github.com/goNiki/ReviewService/models/errorapp"
)

func (s *Service) CreateTeamWithMembers(ctx context.Context, teamWithMembers *domain.TeamWithMembers) (*domain.TeamWithMembers, error) {

	result := &domain.TeamWithMembers{}
	//создание транзакции для обеспечения атомарности (невозможно создать пользователей без создания команды и команду без пользователей)
	err := s.txManager.WithTransaction(ctx, func(ctx context.Context, tx database.Tx) error {
		team, err := s.checkAndcreateTeam(ctx, tx, teamWithMembers.Team.Name)
		if err != nil {
			return err
		}

		member, err := s.handleNewUsers(ctx, tx, teamWithMembers.Members, team.ID)

		if err != nil {
			return err
		}
		result.Team = team
		result.Members = member
		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Service) checkAndcreateTeam(ctx context.Context, tx database.Tx, teamName string) (*domain.Team, error) {
	//получаем, информацию о том, если команда с таким названием
	existingTeam, err := s.teamRepo.GetTeamByName(ctx, tx, teamName)
	if err != nil {
		return nil, fmt.Errorf("%w: failed to check team existence: %v", errorapp.ErrDatabaseQuery, err)
	}

	var team *domain.Team

	if existingTeam != nil {
		return nil, errorapp.ErrTeamAlreadyExists
	} else {
		team, err = s.teamRepo.CreateTeam(ctx, tx, teamName)
		if err != nil {
			return nil, fmt.Errorf("%w: failed to create team: %v", errorapp.ErrDatabaseQuery, err)
		}
	}

	return team, nil
}

func (s *Service) handleNewUsers(ctx context.Context, tx database.Tx, newUsers []*domain.User, teamId int64) ([]*domain.User, error) {
	res := make([]*domain.User, 0, len(newUsers))
	for _, user := range newUsers {
		create, err := s.handleNewUser(ctx, tx, user, teamId)
		if err != nil {
			return nil, err
		}
		res = append(res, create)
	}
	return res, nil
}

func (s *Service) handleNewUser(ctx context.Context, tx database.Tx, newUsers *domain.User, teamId int64) (*domain.User, error) {
	existingUser, err := s.userRepo.GetUserByID(ctx, tx, newUsers.UserId)
	if err != nil {
		return nil, fmt.Errorf("%w: failed to get user: %v", errorapp.ErrDatabaseQuery, err)
	}

	if existingUser != nil {
		newUsers.TeamID = teamId
		updatedUser, err := s.userRepo.UpdateUser(ctx, tx, newUsers)
		if err != nil {
			return nil, fmt.Errorf("%w: failed to update user: %v", errorapp.ErrDatabaseQuery, err)
		}
		return updatedUser, nil
	} else {
		newUsers.TeamID = teamId
		createUser, err := s.userRepo.CreateUser(ctx, tx, newUsers)
		if err != nil {
			return nil, fmt.Errorf("%w: failed to create user: %v", errorapp.ErrDatabaseQuery, err)
		}
		return createUser, nil
	}
}
