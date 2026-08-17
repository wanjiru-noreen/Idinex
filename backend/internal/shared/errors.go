package shared

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid_credentials")
	ErrInvalidRequest     = errors.New("invalid_request")
	ErrEmailAlreadyExists = errors.New("email_already_exists")
	ErrInternalServer     = errors.New("internal_server_error")
)
