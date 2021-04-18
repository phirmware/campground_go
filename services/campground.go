package services

import (
	"campground_go/session"
	"net/http"
)

type CampgroundService struct {
	session *session.Session
}

func NewCampgroundService() *CampgroundService {
	s := session.NewSession()
	return &CampgroundService{
		session: s,
	}
}

func (cs *CampgroundService) IsOwner(w http.ResponseWriter, r *http.Request, campgroundOwnerId uint) bool {
	values, err := cs.session.GetSessionValues(w, r)
	if err != nil {
		return false
	}

	if values.UserId == campgroundOwnerId {
		return true
	}

	return false
}
