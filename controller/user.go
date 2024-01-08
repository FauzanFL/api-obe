package controller

import (
	"api-obe/model"
	repo "api-obe/repository"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserController interface {
	GetUser(c *gin.Context)
	AddUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
}

type userController struct {
	userRepo repo.UserRepository
}

func NewUserController(userRepo repo.UserRepository) UserController {
	return &userController{userRepo}
}

func (u *userController) GetUser(c *gin.Context) {
	user, err := u.userRepo.GetUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (u *userController) AddUser(c *gin.Context) {
	var body struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if body.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email is empty"})
		return
	}
	if body.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password is empty"})
		return
	}

	recordUser, err := u.userRepo.GetUserByEmail(body.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if recordUser.Email != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email already exists"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user model.User
	user.Email = body.Email
	user.Password = string(hash)
	user.Role = "dosen"

	if err := u.userRepo.Add(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User added successfully"})
}

func (u *userController) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var body struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if body.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email is empty"})
		return
	}
	if body.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password is empty"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user model.User
	user.Email = body.Email
	user.Password = string(hash)
	user.ID = id

	if err := u.userRepo.Update(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (u *userController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := u.userRepo.Delete(idInt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delete Account Successful"})
}

func (u *userController) Login(c *gin.Context) {
	var body struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if body.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email is empty"})
		return
	}
	if body.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password is empty"})
		return
	}

	recordUser, err := u.userRepo.GetUserByEmail(body.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if recordUser.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email or password"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(recordUser.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email or password"})
		return
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": recordUser.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to create token"})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Login success"})
}

func (u *userController) Logout(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logout success"})
}
