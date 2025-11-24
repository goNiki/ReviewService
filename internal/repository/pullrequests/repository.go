package pullrequests

import (
	"github.com/goNiki/ReviewService/internal/infrastructure/database"
)

type Repository struct {
	db *database.Db
}

func NewPullRequestRepository(db *database.Db) *Repository {
	return &Repository{
		db: db,
	}
}
