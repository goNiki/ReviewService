package domain

const DefoltTeam = 1

type TeamWithMembers struct {
	Team    *Team
	Members []*User
}

type Team struct {
	ID   int64
	Name string
}

type User struct {
	UserId   string
	Username string
	TeamID   int64
	TeamName string
	IsActive bool
}
