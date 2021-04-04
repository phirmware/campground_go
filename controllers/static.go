package controllers

import (
	"campground_go/views"
	"net/http"
)

type Static struct {
	views *views.Views
}

func NewStatic() *Static {
	v := views.NewView()
	return &Static{
		views: v,
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
	s.views.DashboardView.ExecuteTemplate(w, "bootstrap", nil)
}
