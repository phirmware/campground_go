package controllers

import (
	"campground_go/models"
	"campground_go/services"
	"campground_go/session"
	"campground_go/utils"
	"campground_go/views"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type CampgroundController struct {
	CampgroundDB      *models.CampgroundDB
	CommentDB         *models.CommentDB
	CampgroundService *services.CampgroundService
	Views             *views.Views
	Session           *session.Session
}

type CampgroundForm struct {
	Name        string
	Image       string
	Description string
}

type CampgroundView struct {
	Campground *models.Campground
	Comments   *[]models.Comment
}

func NewCampgorund() *CampgroundController {
	views := views.NewView()
	campground := models.NewCampground()
	comment := models.NewComment()
	cs := services.NewCampgroundService()
	s := session.NewSession()
	return &CampgroundController{
		Views:             views,
		CampgroundDB:      campground,
		CommentDB:         comment,
		CampgroundService: cs,
		Session:           s,
	}
}

func (c *CampgroundController) CreateCampground(w http.ResponseWriter, r *http.Request) {
	var body CampgroundForm
	if err := utils.GetRequestBody(r, &body); err != nil {
		panic(err)
	}

	s, err := c.Session.GetSessionValues(w, r)
	if err != nil {
		panic(err)
	}

	campground := &models.Campground{
		Name:        body.Name,
		Description: body.Description,
		Image:       body.Image,
		OwnerID:     s.UserId,
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

	comments, err := c.CommentDB.FindByCampgroundName(campgroundName)
	if err != nil {
		fmt.Println("*********** Error ***********")
		panic(err)
	}

	fmt.Printf("%+v", comments)

	data := &CampgroundView{
		Comments: comments,
		Campground: campground,
	}

	if err := c.Views.RenderUserPage(w, "campground.html", data); err != nil {
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

func (c *CampgroundController) DeleteCampground(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	campgroundName := vars["name"]

	foundCampground, err := c.CampgroundDB.FindByName(campgroundName)
	if err != nil {
		panic(err)
	}

	if isOwner := c.CampgroundService.IsOwner(w, r, foundCampground.OwnerID); isOwner != true {
		url := fmt.Sprintf("/campground/%s", foundCampground.Name)
		http.Redirect(w, r, url, http.StatusFound)
		return
	}

	if err := c.CampgroundDB.DeleteByName(campgroundName); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/dashboard", http.StatusFound)
}
