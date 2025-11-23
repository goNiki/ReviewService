package main

import (
	"bufio"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/goNiki/ReviewService/internal/infrastructure/config"
	"github.com/goNiki/ReviewService/internal/infrastructure/database"
	"github.com/goNiki/ReviewService/internal/infrastructure/logger"
	"github.com/goNiki/ReviewService/internal/infrastructure/logger/sl"
	"github.com/goNiki/ReviewService/internal/infrastructure/migrator"
)

const MigDir = "./migrations"

var configpath = "./configs/.env"

func main() {

	cfg, err := config.InitConfig(configpath)
	if err != nil {
		log.Fatal(err)
	}

	l := logger.InitLogger(cfg.ServerConfig.Env)

	l.Log.Info("Конфиг иннициализирован")

	postgres, err := database.InitDatabase(cfg.DBConfig)
	if err != nil {
		log.Fatal(err)
		return
	}
	l.Log.Info("Подключение к базе данных выполнено")

	defer postgres.Pool.Close()
	migrator := migrator.NewMigrator(postgres.Pool, MigDir)
	Migration(l.Log, migrator)

}

func Migration(log *slog.Logger, migrator *migrator.Migrator) {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		var choice int64

		fmt.Println("Выберите действие: \n1 - Создать файл для миграции \n2 - Проверить статус миграций \n3 - Миграций UP \n4- Миграций UpTo \n5 - Миграция Down \n6 - Миграция DownTo \n7 - выйти")
		if scanner.Scan(); scanner.Err() != nil {
			fmt.Println("Ошибка ввода")
			continue
		}
		if _, err := fmt.Sscanf(scanner.Text(), "%d", &choice); err != nil || choice < 1 || choice > 7 {
			fmt.Println("Ошибка ввода. Не соответствует условиям")
			continue
		}

		switch choice {
		case 1:
			fmt.Println("Введите название файла ")
			scanner.Scan()
			name := scanner.Text()
			if err := migrator.Create(name, "sql"); err != nil {
				log.Error("Ошибка: ", sl.Error(err))
			}
			log.Info("migration created successfully")

		case 2:
			if err := migrator.Status(); err != nil {
				log.Error("Ошибка", sl.Error(err))
			}
			log.Info("Status migrations apply")
		case 3:
			if err := migrator.Up(); err != nil {
				log.Error("Ошибка", sl.Error(err))
			} else {
				log.Info("UP Migration apply")
			}
		case 4:
			fmt.Println("Введите версию миграции:")
			scanner.Scan()
			var name int64
			if _, err := fmt.Sscanf(scanner.Text(), "%d", &name); err != nil {
				fmt.Println("Ошибка при вводе версии миграции")
				continue
			}
			if err := migrator.UpTo(name); err != nil {
				log.Error("", sl.Error(err))
			}
			log.Info("UpTo migration apply")
		case 5:
			if err := migrator.Down(); err != nil {
				log.Error("", sl.Error(err))
			}
			log.Info("DOWN migration apply")
		case 6:
			fmt.Println("Введите версию миграции:")
			scanner.Scan()
			var name int64
			if _, err := fmt.Sscanf(scanner.Text(), "%d", &name); err != nil {
				fmt.Println("Ошибка при вводе версии миграции")
				continue
			}

			if err := migrator.DownTo(name); err != nil {
				log.Error("", sl.Error(err))
			}
			log.Info("DownTo migration apply")
		case 7:
			return
		}
	}
}
