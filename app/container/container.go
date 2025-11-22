package container

import (
	"fmt"
	"net"
	"net/http"

	"github.com/goNiki/ReviewService/internal/infrastructure/config"
	"github.com/goNiki/ReviewService/internal/infrastructure/database"
	"github.com/goNiki/ReviewService/internal/infrastructure/logger"
	"github.com/goNiki/ReviewService/models/errorapp"
	"github.com/goNiki/ReviewService/shared/pkg/openapi/reviewerservice/v1"
)

type Container struct {
	Config    *config.Config
	Log       *logger.Log
	Db        *database.Db
	ApiServer *reviewerservice.Server
	Server    *http.Server
}

var configPath = "./configs/configs.yaml"

func NewContainer() (*Container, error) {

	c := &Container{}
	var err error

	c.Config, err = config.InitConfig(configPath)
	if err != nil {
		return &Container{}, err
	}

	c.Log = logger.InitLogger(c.Config.ServerConfig.Env)

	c.Db, err = database.InitDatabase(c.Config.DBConfig)
	if err != nil {
		return &Container{}, err
	}

	c.ApiServer, err = reviewerservice.NewServer(reviewerservice.UnimplementedHandler{})
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
