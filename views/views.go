package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type Views struct {
	HomeView       *template.Template
	CampgroundView *template.Template
	LoginView      *template.Template
	SignupView     *template.Template
	DashboardView  *template.Template
}

func getPartialFiles() ([]string, error) {
	files, err := filepath.Glob("views/partials")
	if err != nil {
		return nil, err
	}
	return files, nil
}

func NewView() *Views {
	home, err := template.ParseFiles("views/partials/nav.html", "views/home.html", "views/partials/bootstrap.html")
	campgrounds, err := template.ParseFiles("views/partials/nav.html", "views/campgrounds.html", "views/partials/bootstrap.html")
	login, err := template.ParseFiles("views/partials/nav.html", "views/user/login.html", "views/partials/bootstrap.html")
	signup, err := template.ParseFiles("views/partials/nav.html", "views/user/signup.html", "views/partials/bootstrap.html")
	dashboard, err := template.ParseFiles("views/partials/nav.html", "views/user/dashboard.html", "views/partials/bootstrap.html")
	if err != nil {
		panic(err)
	}
	return &Views{
		HomeView:       home,
		CampgroundView: campgrounds,
		LoginView:      login,
		SignupView:     signup,
		DashboardView:  dashboard,
	}
}

func (v *Views) RenderUserPage(w http.ResponseWriter, path string, data interface{}) error {
	t, err := template.ParseFiles("views/partials/nav.html", "views/user/" + path, "views/partials/bootstrap.html")
	if err != nil {
		return err
	}

	return t.ExecuteTemplate(w, "bootstrap", data)
}
