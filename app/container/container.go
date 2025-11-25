package container

import (
	"fmt"
	"net"
	"net/http"

	"github.com/goNiki/ReviewService/internal/api"
	pullRequestsApi "github.com/goNiki/ReviewService/internal/api/pullrequests"
	teamsApi "github.com/goNiki/ReviewService/internal/api/teams"
	usersApi "github.com/goNiki/ReviewService/internal/api/users"
	"github.com/goNiki/ReviewService/internal/infrastructure/config"
	"github.com/goNiki/ReviewService/internal/infrastructure/database"
	txmanager "github.com/goNiki/ReviewService/internal/infrastructure/database/transactionManager"
	"github.com/goNiki/ReviewService/internal/infrastructure/logger"
	"github.com/goNiki/ReviewService/internal/infrastructure/migrator"
	pullRequestRepo "github.com/goNiki/ReviewService/internal/repository/pullrequests"
	teamsRepo "github.com/goNiki/ReviewService/internal/repository/teams"
	usersRepo "github.com/goNiki/ReviewService/internal/repository/users"
	pullRequestService "github.com/goNiki/ReviewService/internal/services/pullrequests"
	teamsService "github.com/goNiki/ReviewService/internal/services/teams"
	usersService "github.com/goNiki/ReviewService/internal/services/users"
	"github.com/goNiki/ReviewService/models/errorapp"
	"github.com/goNiki/ReviewService/shared/pkg/openapi/reviewerservice/v1"
)

type Container struct {
	Config             *config.Config
	Log                *logger.Log
	Db                 *database.Db
	Migration          *migrator.Migrator
	TransactionManager *txmanager.TxManager
	TeamsRepo          *teamsRepo.Repository
	UserRepo           *usersRepo.Repository
	PullRequestRepo    *pullRequestRepo.Repository
	TeamsService       *teamsService.Service
	UserService        *usersService.Service
	PullRequestService *pullRequestService.Service
	PullRequestsApiApi *pullRequestsApi.Api
	TeamsApi           *teamsApi.Api
	UsersApi           *usersApi.Api
	CompositeApi       *api.CompositeApi
	ApiServer          *reviewerservice.Server
	Server             *http.Server
}

var (
	migrationPath = "./migrations"
)

func NewContainer() (*Container, error) {

	c := &Container{}
	var err error

	c.Config, err = config.InitConfig()
	if err != nil {
		return &Container{}, err
	}

	c.Log = logger.InitLogger(c.Config.ServerConfig.Env)

	c.Db, err = database.InitDatabase(c.Config.DBConfig)
	if err != nil {
		return &Container{}, err
	}

	c.Migration = migrator.NewMigrator(c.Db.Pool, migrationPath)
	c.Migration.Up()

	c.TransactionManager = txmanager.NewTransactionManager(c.Db)

	c.TeamsRepo = teamsRepo.NewTeamsRepository(c.Db)

	c.UserRepo = usersRepo.NewUsersRepository(c.Db)

	c.PullRequestRepo = pullRequestRepo.NewPullRequestRepository(c.Db)

	c.TeamsService = teamsService.NewTeamsService(c.TransactionManager, c.TeamsRepo, c.UserRepo)

	c.UserService = usersService.NewUserService(c.UserRepo, c.PullRequestRepo)

	c.PullRequestService = pullRequestService.NewPullRequestService(c.TransactionManager, c.UserRepo, c.PullRequestRepo)

	c.PullRequestsApiApi = pullRequestsApi.NewPullRequestsApi(c.Log, c.PullRequestService)

	c.TeamsApi = teamsApi.NewTeamsApi(c.Log, c.TeamsService)

	c.UsersApi = usersApi.NewUsersApi(c.Log, c.UserService)

	c.CompositeApi = api.NewCompositeApi(c.TeamsApi, c.UsersApi, c.PullRequestsApiApi)

	c.ApiServer, err = reviewerservice.NewServer(c.CompositeApi)
	if err != nil {
		return &Container{}, fmt.Errorf("%w: %v", errorapp.ErrCreateApiServer, err)
	}

	c.Server = &http.Server{
		Addr:        net.JoinHostPort(c.Config.ServerConfig.Host, c.Config.ServerConfig.Port),
		ReadTimeout: c.Config.ServerConfig.Timeout,
		IdleTimeout: c.Config.ServerConfig.IdleTimeout,
	}

	return c, nil
}
