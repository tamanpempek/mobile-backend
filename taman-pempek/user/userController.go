package user

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type controller struct {
	userService UserService
}

func NewController(userService UserService) *controller {
	return &controller{userService}
}

func (cn *controller) GetUsers(c *gin.Context) {
	users, err := cn.userService.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   err,
		})
		return
	}

	var usersResponse []UserResponse

	for _, user := range users {
		userResponse := convertToUserResponse(user)

		usersResponse = append(usersResponse, userResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Success!",
		"data":  usersResponse,
	})
}

func (cn *controller) FindUsersByRole(c *gin.Context) {
	role := c.Param("role")
	users, err := cn.userService.FindUsersByRole(role)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   err,
		})
		return
	}

	var usersResponse []UserResponse

	for _, user := range users {
		userResponse := convertToUserResponse(user)

		usersResponse = append(usersResponse, userResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Success!",
		"data":  usersResponse,
	})
}

func (cn *controller) GetUser(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   "Invalid user ID",
		})
		return
	}

	user, err := cn.userService.FindUserByID(id)

	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "User not found" {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, gin.H{
			"error": true,
			"data":  nil,
			"msg":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Success!",
		"data":  convertToUserResponse(user),
	})
}

func (cn *controller) CreateUser(c *gin.Context) {
	var userRequest UserCreateRequest

	err := c.ShouldBindJSON(&userRequest)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on %s field, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   errorMessages,
		})
		return
	}

	user, err := cn.userService.CreateUser(userRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Success!",
		"data":  convertToUserResponse(user),
	})
}

func (cn *controller) UpdateUser(c *gin.Context) {
	var userRequest UserUpdateRequest

	err := c.ShouldBindJSON(&userRequest)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on %s field, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   errorMessages,
		})
		return
	}

	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   "Invalid user ID",
		})
		return
	}

	user, err := cn.userService.UpdateUser(id, userRequest)

	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "User not found" {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, gin.H{
			"error": true,
			"data":  nil,
			"msg":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Success!",
		"data":  convertToUserResponse(user),
	})
}

func (ch *controller) DeleteUser(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   "Invalid user ID",
		})
		return
	}

	user, err := ch.userService.DeleteUser(id)

	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "User not found" {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, gin.H{
			"error": true,
			"data":  nil,
			"msg":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Success!",
		"data":  convertToUserResponse(user),
	})
}

func (ch *controller) Login(c *gin.Context) {
	var request LoginRequest

	if c.Bind(&request) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"token": nil,
			"msg":   "Failed to read request",
		})
		return
	}

	user, _ := ch.userService.FindUserByEmail(request.Email)

	if user.ID == 0 || user.Password != request.Password {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"token": nil,
			"msg":   "Invalid email or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	fmt.Println(tokenString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"token": nil,
			"msg":   "Failed to create token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*7, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Success!",
		"token": tokenString,
		"data":  convertToUserResponse(user),
	})
}

func (cn *controller) Logout(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "", "", false, true)
	c.Set("UserID", nil)
	c.Set("UserName", nil)
	c.Set("UserEmail", nil)

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Logged out successfully!",
	})
}

func convertToUserResponse(user User) UserResponse {
	return UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Whatsapp: user.Whatsapp,
		Gender:   user.Gender,
		Role:     user.Role,
	}
}
