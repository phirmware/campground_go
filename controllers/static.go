package controllers

import (
	"campground_go/models"
	"campground_go/session"
	"campground_go/views"
	"fmt"
	"net/http"
)

type Static struct {
	views        *views.Views
	Session      *session.Session
	CampgroundDB *models.CampgroundDB
}

func NewStatic() *Static {
	v := views.NewView()
	s := session.NewSession()
	c := models.NewCampground()
	return &Static{
		views:   v,
		Session: s,
		CampgroundDB: c,
	}
}

func (s *Static) HandleHome(w http.ResponseWriter, r *http.Request) {
	s.views.HomeView.ExecuteTemplate(w, "bootstrap", nil)
}

func (s *Static) HandleCampgrounds(w http.ResponseWriter, r *http.Request) {
	s.views.CampgroundView.ExecuteTemplate(w, "bootstrap", nil)
}

func (s *Static) HandleLogin(w http.ResponseWriter, r *http.Request) {
	s.views.LoginView.ExecuteTemplate(w, "bootstrap", nil)
}

func (s *Static) HandleSignup(w http.ResponseWriter, r *http.Request) {
	s.views.SignupView.ExecuteTemplate(w, "bootstrap", nil)
}

func (s *Static) HandleDashboard(w http.ResponseWriter, r *http.Request) {
	session, err := s.Session.GetSessionValues(w, r)
	if err != nil {
		panic(err)
	}

	fmt.Println("Currently fetching campgrounds for the dashboard view", session.UserId)

	campgrounds, err := s.CampgroundDB.FindUsersCampgroundsByOwnerID(session.UserId)
	fmt.Printf("%+v", campgrounds)
	if err != nil {
		fmt.Println("An error occured here")
		panic(err)
	}

	s.views.DashboardView.ExecuteTemplate(w, "bootstrap", campgrounds)
}
