package middlewares

import (
	"campground_go/session"
	"net/http"
)

type Middleware struct {
	session *session.Session
}

func NewMiddleware() *Middleware {
	session := session.NewSession()
	return &Middleware{
		session: session,
	}
}

func (m *Middleware) IsUserLoggedIn(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := m.session.AuthenticateSession(w, r); err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		f(w, r)
	}
}
