package repomodels

import "time"

type User struct {
	UserID    string
	UserName  string
	TeamID    int64
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserWithTeamName struct {
	UserID   string
	UserName string
	TeamName string
	IsActive bool
}

type Team struct {
	Id         int64
	Name       string
	Created_at time.Time
}
