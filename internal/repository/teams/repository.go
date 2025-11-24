package teams

import "github.com/goNiki/ReviewService/internal/infrastructure/database"

type Repository struct {
	db *database.Db
}

func NewTeamsRepository(db *database.Db) *Repository {
	return &Repository{
		db: db,
	}
}
