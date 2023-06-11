package auth

import (
	"context"

	"github.com/gorilla/sessions"
)

type Authenticator interface {
	Authenticate(ctx context.Context)
}

type service struct {
	store sessions.Store
}
