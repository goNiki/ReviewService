package users

import "github.com/goNiki/ReviewService/internal/infrastructure/database"

type Repository struct {
	db *database.Db
}

func NewUsersRepository(db *database.Db) *Repository {
	return &Repository{
		db: db,
	}
}
