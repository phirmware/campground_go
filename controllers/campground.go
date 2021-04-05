package controllers

import (
	"campground_go/models"
	"campground_go/utils"
	"campground_go/views"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type CampgroundController struct {
	CampgroundDB *models.CampgroundDB
	Views        *views.Views
}

type CampgroundForm struct {
	Name        string
	Image       string
	Description string
}

func NewCampgorund() *CampgroundController {
	views := views.NewView()
	campground := models.NewCampground()
	return &CampgroundController{
		Views: views,
		CampgroundDB: campground,
	}
}

func (c *CampgroundController) CreateCampground(w http.ResponseWriter, r *http.Request) {
	var body CampgroundForm
	if err := utils.GetRequestBody(r, &body); err != nil {
		panic(err)
	}

	campground := &models.Campground{
		Name:        body.Name,
		Description: body.Description,
		Image:       body.Image,
	}

	if err := c.CampgroundDB.Create(campground); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/campground/"+campground.Name, http.StatusFound)
}

func (c *CampgroundController) DisplayCampground(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	campgroundName := vars["name"]

	campground, err := c.CampgroundDB.FindByName(campgroundName)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", campground)

	if err := c.Views.RenderUserPage(w, "campground.html", campground); err != nil {
		panic(err)
	}
}

func (c *CampgroundController) DisplayCampgorunds(w http.ResponseWriter, r *http.Request) {
	campgrounds, err := c.CampgroundDB.Find()
	if err != nil {
		panic(err)
	}

	views.NewView().RenderUserPage(w, "campgrounds.html", campgrounds)
}
