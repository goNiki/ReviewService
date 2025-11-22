package errorapp

import "errors"

var (
	ErrInitConfig      = errors.New("error init config")
	ErrInitDatabase    = errors.New("error init database")
	ErrCreateApiServer = errors.New("error create api server")
	ErrListenAndServe  = errors.New("error listenandserve")
)
