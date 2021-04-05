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
	campgroundC := controllers.NewCampgorund()
	middleware := middlewares.NewMiddleware()

	r.HandleFunc("/", staticC.HandleHome)
	r.HandleFunc("/campgrounds", campgroundC.DisplayCampgorunds).Methods("GET")
	r.HandleFunc("/campground", middleware.IsUserLoggedIn(campgroundC.CreateCampground)).Methods("POST")
	r.HandleFunc("/campground/{name}", middleware.IsUserLoggedIn(campgroundC.DisplayCampground)).Methods("GET")
	r.HandleFunc("/login", staticC.HandleLogin).Methods("GET")
	r.HandleFunc("/login", userC.LoginUser).Methods("POST")
	r.HandleFunc("/signup", staticC.HandleSignup).Methods("GET")
	r.HandleFunc("/signup", userC.RegisterUser).Methods("POST")
	r.HandleFunc("/dashboard", middleware.IsUserLoggedIn(staticC.HandleDashboard)).Methods("GET")

	http.ListenAndServe(":3000", r)
}
