package pullrequests

import (
	"github.com/goNiki/ReviewService/internal/infrastructure/logger"
	"github.com/goNiki/ReviewService/internal/services/pullrequests"
)

type Api struct {
	log       *logger.Log
	prService *pullrequests.Service
}

func NewPullRequestsApi(log *logger.Log, prService *pullrequests.Service) *Api {
	return &Api{
		log:       log,
		prService: prService,
	}
}
