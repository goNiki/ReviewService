package teams

import (
	"github.com/goNiki/ReviewService/internal/infrastructure/logger"
	"github.com/goNiki/ReviewService/internal/services/teams"
)

type Api struct {
	log         *logger.Log
	teamService *teams.Service
}

func NewTeamsApi(log *logger.Log, teamService *teams.Service) *Api {
	return &Api{
		log:         log,
		teamService: teamService,
	}
}

