package controllers

import (
	"campground_go/models"
	"campground_go/services"
	"campground_go/session"
	"campground_go/utils"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type UserController struct {
	UserDB      *models.UserDB
	CommentDB   *models.CommentDB
	UserService *services.UserService
	Session     *session.Session
}

type UserForm struct {
	Username string
	Password string
}

type CommentForm struct {
	Comment string
}

func NewUser() *UserController {
	userDB := models.NewUser()
	commentDB := models.NewComment()
	userService := services.NewUserService()
	session := session.NewSession()
	return &UserController{
		UserDB:      userDB,
		CommentDB:   commentDB,
		UserService: userService,
		Session:     session,
	}
}

func (u *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var body UserForm
	if err := utils.GetRequestBody(r, &body); err != nil {
		panic(err)
	}

	if body.Password == "" || body.Username == "" {
		fmt.Println("No username or password")
		http.Redirect(w, r, "/signup", http.StatusFound)
		return
	}

	user := &models.User{
		Username: body.Username,
		Password: body.Password,
	}

	if err := u.UserDB.Create(user); err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/signup", http.StatusFound)
		return
	}

	if err := u.authenticateUser(w, r, user); err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/signup", http.StatusFound)
		return
	}

	http.Redirect(w, r, "/dashboard", http.StatusFound)
}

func (u *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	var body UserForm
	if err := utils.GetRequestBody(r, &body); err != nil {
		panic(err)
	}

	user := &models.User{
		Username: body.Username,
		Password: body.Password,
	}

	if err := u.authenticateUser(w, r, user); err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	http.Redirect(w, r, "/dashboard", http.StatusFound)
}

func (u *UserController) authenticateUser(w http.ResponseWriter, r *http.Request, user *models.User) error {
	foundUser, err := u.UserService.Authenticate(user)
	if err != nil {
		return err
	}

	if err := u.Session.CreateSession(w, r, foundUser); err != nil {
		return err
	}

	return nil
}

func (u *UserController) CreateComment(w http.ResponseWriter, r *http.Request) {
	session, err := u.Session.GetSessionValues(w, r)
	if err != nil {
		panic(err)
	}

	params := mux.Vars(r)

	var body CommentForm
	if err := utils.GetRequestBody(r, &body); err != nil {
		panic(err)
	}

	comment := &models.Comment{
		Comment:        body.Comment,
		UserID:         session.UserId,
		CampgroundName: params["name"],
	}

	if comment.Comment == "" {
		http.Redirect(w, r, "/campground/" + params["name"], http.StatusFound)
		return
	}

	if err := u.CommentDB.CreateComment(comment); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/campground/" + params["name"], http.StatusFound)
}
