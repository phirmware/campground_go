package session

import (
	"campground_go/models"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key         = []byte("super-secret-key")
	cookie_name = "secret-cookie-name"
)

type Session struct {
	UserDB *models.UserDB
	store  *sessions.CookieStore
}

type SessionValues struct {
	Authenticated bool
	Username      string
	UserId        uint
}

func NewSession() *Session {
	store := sessions.NewCookieStore(key)
	userDB := models.NewUser()
	return &Session{
		store:  store,
		UserDB: userDB,
	}
}

func (s *Session) CreateSession(w http.ResponseWriter, r *http.Request, user *models.User) error {
	session, err := s.store.Get(r, cookie_name)
	if err != nil {
		return err
	}

	session.Values["authenticated"] = true
	session.Values["username"] = user.Username
	session.Values["id"] = user.ID
	session.Save(r, w)
	return nil
}

func (s *Session) DestroySession(w http.ResponseWriter, r *http.Request) error {
	session, err := s.store.Get(r, cookie_name)
	if err != nil {
		return nil
	}

	session.Values["authenticated"] = false
	session.Save(r, w)
	return nil
}

func (s *Session) AuthenticateSession(w http.ResponseWriter, r *http.Request) error {
	session, _ := s.store.Get(r, cookie_name)

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		return errors.New("Invalid session")
	}

	id, _ := session.Values["id"].(uint)

	user, err := s.UserDB.FindUserByID(uint(id))
	if err != nil {
		return err
	}

	fmt.Printf("From session authenticator - %+v", user)
	return nil
}

func (s *Session) GetSessionValues(w http.ResponseWriter, r *http.Request) (SessionValues, error) {
	session, _ := s.store.Get(r, cookie_name)

	username := session.Values["username"]
	authenticated, _ := session.Values["authenticated"].(bool)
	id, _ := session.Values["id"].(uint)

	value := SessionValues{
		Username:      fmt.Sprintf("%s", username),
		Authenticated: authenticated,
		UserId:        id,
	}

	return value, nil

}
