package iam

import "errors"

var (
	ErrUnauthenticated   = errors.New("unauthenticated")
	ErrPermissionDenied  = errors.New("permission denied")
	ErrInvalidCredential = errors.New("invalid credential")
	ErrInvalidPolicy     = errors.New("invalid policy")
	ErrInvalidPrincipal  = errors.New("invalid principal")
)
