package main

import (
	"campground_go/controllers"
	"campground_go/middlewares"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const PORT = ":3000"

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
	r.HandleFunc("/campground/delete/{name}", middleware.IsUserLoggedIn(campgroundC.DeleteCampground)).Methods("POST")
	r.HandleFunc("/campground/{name}/comment", middleware.IsUserLoggedIn(userC.CreateComment)).Methods("POST")
	r.HandleFunc("/login", staticC.HandleLogin).Methods("GET")
	r.HandleFunc("/login", userC.LoginUser).Methods("POST")
	r.HandleFunc("/signup", staticC.HandleSignup).Methods("GET")
	r.HandleFunc("/signup", userC.RegisterUser).Methods("POST")
	r.HandleFunc("/dashboard", middleware.IsUserLoggedIn(staticC.HandleDashboard)).Methods("GET")

	fmt.Println(fmt.Sprintf("Server running on port %s from the cake branch yo!", PORT))
	http.ListenAndServe(PORT, r)
}
