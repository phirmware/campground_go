package main

import (
	"campground_go/controllers"
	"campground_go/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	staticC := controllers.NewStatic()
	userC := controllers.NewUser()
	middleware := middlewares.NewMiddleware()

	r.HandleFunc("/", staticC.HandleHome)
	r.HandleFunc("/campgrounds", staticC.HandleCampgrounds)
	r.HandleFunc("/login", staticC.HandleLogin).Methods("GET")
	r.HandleFunc("/login", userC.LoginUser).Methods("POST")
	r.HandleFunc("/signup", staticC.HandleSignup).Methods("GET")
	r.HandleFunc("/signup", userC.RegisterUser).Methods("POST")
	r.HandleFunc("/dashboard", middleware.IsUserLoggedIn(staticC.HandleDashboard)).Methods("GET")

	http.ListenAndServe(":3000", r)
}
