package errorapp

import "errors"

var (
	ErrInitConfig      = errors.New("error init config")
	ErrInitDatabase    = errors.New("error init database")
	ErrCreateApiServer = errors.New("error create api server")
	ErrListenAndServe  = errors.New("error listenandserve")

	ErrDatabaseQuery      = errors.New("database query error")
	ErrDatabaseConnection = errors.New("database connection error")
	ErrTransactionFailed  = errors.New("transaction failed")

	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")

	ErrTeamNotFound      = errors.New("team not found")
	ErrTeamAlreadyExists = errors.New("team already exists")

	ErrPullRequestNotFound = errors.New("pull request not found")
	ErrPRExists            = errors.New("pull request already exists")
	ErrInvalidReviewer     = errors.New("invalid reviewer")
	ErrNoCandidate         = errors.New("No candidate ")
	ErrPRAlreadyMerged     = errors.New("pull request already merged")
)
