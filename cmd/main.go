package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/goNiki/ReviewService/app/container"
)

func main() {

	c, err := container.NewContainer()
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()

	r.Mount("/", c.ApiServer)

	c.Server.Handler = r
	fmt.Println("Сервер запускается ")
	go func() {
		if err := c.Server.ListenAndServe(); err != nil {
			c.Log.Log.Error("Ошибка запуска сервера", "error", err)
		}
	}()

	c.Log.Log.Info("Сервер запущен на порту", "port", c.Config.ServerConfig.Port)

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
