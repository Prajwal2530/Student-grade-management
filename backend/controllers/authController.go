package controllers

import (
	"grade-management-system/config"
	"grade-management-system/models"
	"grade-management-system/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// RegisterStudent allows anyone to register, but ONLY as a Student.
// Admin and Teacher accounts must be created by an Admin or Seeded.
func RegisterStudent(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Check if user already exists
	var existingUser models.User
	if err := config.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		utils.ErrorResponse(c, http.StatusConflict, "Email already in use")
		return
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	user := models.User{
		Name:     input.Name,
		Email:    strings.ToLower(input.Email),
		Password: hashedPassword,
		Role:     "student", // Forced to student
	}

	if err := config.DB.Create(&user).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create user")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Student registered successfully", user)
}

// Login authenticates a user and returns a JWT
func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", strings.ToLower(input.Email)).First(&user).Error; err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Login successful", gin.H{
		"token": token,
		"user":  user,
	})
}
