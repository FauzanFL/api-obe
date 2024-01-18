package middleware

import (
	"api-obe/model"
	"api-obe/repository"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware interface {
	RequireAuth(c *gin.Context)
	RequireAdminAuth(c *gin.Context)
	RequireDosenAuth(c *gin.Context)
}

type authMiddleware struct {
	userRepo repository.UserRepository
}

func NewAuthMiddleware(userRepo repository.UserRepository) AuthMiddleware {
	return &authMiddleware{userRepo}
}

func (am *authMiddleware) RequireAuth(c *gin.Context) {
	// Get Cookie
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		user, _ := am.userRepo.GetUserById(int(claims["sub"].(float64)))
		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("user", user)

		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func (am *authMiddleware) RequireAdminAuth(c *gin.Context) {
	am.RequireAuth(c)

	user, exists := c.Get("user")
	if !exists {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userExist := user.(model.User)

	if userExist.Role != "admin" {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	c.Next()
}

func (am *authMiddleware) RequireDosenAuth(c *gin.Context) {
	am.RequireAuth(c)

	user, exists := c.Get("user")
	if !exists {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userExist := user.(model.User)

	if userExist.Role != "dosen" {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	c.Next()
}
