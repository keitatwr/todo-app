package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keitatwr/todo-app/domain"
	"golang.org/x/crypto/bcrypt"
)

type SignupController struct {
	SignupUsecase domain.SignupUsecase
}

func (sc *SignupController) Signup(c *gin.Context) {
	var request domain.SignupRequest

	// binding json request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: err.Error()})
		return
	}

	// check if user already exists
	_, err := sc.SignupUsecase.GetUserByEmail(c, request.Email)
	if err == nil {
		c.JSON(http.StatusConflict, domain.ErrorResponse{
			Message: fmt.Sprintf("user with email %s already exists", request.Email)})
		return
	}

	// password hashing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "failed to hash password"})
		return
	}
	request.Password = string(hashedPassword)

	// create user
	err = sc.SignupUsecase.Create(c, request.Name, request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: fmt.Sprintf("failed to create user: %v", err)})
		return
	}

	c.JSON(http.StatusCreated, domain.SuccessResponse{Message: "user created"})

}
