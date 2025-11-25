package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/goNiki/ReviewService/app/container"
	"github.com/goNiki/ReviewService/internal/infrastructure/swagger"
)

//go:embed openapi-bundled.yaml
var swaggerDoc []byte

func main() {

	c, err := container.NewContainer()
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)

	swagger.RegisterRoutes(r, swaggerDoc)

	r.Mount("/", c.ApiServer)

	c.Server.Handler = r

	serverAddr := fmt.Sprintf("%s:%s", c.Config.ServerConfig.Host, c.Config.ServerConfig.Port)
	fmt.Printf("Сервер запускается на %s\n", serverAddr)

	go func() {
		if err := c.Server.ListenAndServe(); err != nil {
			log.Printf("Ошибка запуска сервера: %v\n", err)
		}
	}()

	fmt.Printf("Сервер запущен на порту %s\n", c.Config.ServerConfig.Port)

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	c.Log.Log.Info("Остановка сервера")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := c.Server.Shutdown(ctx); err != nil {
		c.Log.Log.Error("Ошибка остановки сервера", "error", err)
	}

	c.Log.Log.Info("Сервер остановлен")

}
