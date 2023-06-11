package auth

import (
	"github.com/gorilla/sessions"
)

func SessionStore() sessions.Store {
	// Note: Don't store your key in your source code. Pass it via an
	// environmental variable, or flag (or both), and don't accidentally commit it
	// alongside your code. Ensure your key is sufficiently random - i.e. use Go's
	// crypto/rand or securecookie.GenerateRandomKey(32) and persist the result.
	return sessions.NewCookieStore([]byte("AUTH_SESSION KEY"))
}
