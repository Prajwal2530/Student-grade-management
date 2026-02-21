package controllers

import (
	"grade-management-system/config"
	"grade-management-system/models"
	"grade-management-system/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// HealthCheck endpoint to verify server availability
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "Server is running",
	})
}

// CreateUser allowed only for Admins to create Teachers or Students
type CreateUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"required,oneof=teacher student"`
}

func CreateUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

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
		Role:     input.Role,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create user")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "User created successfully", user)
}

type CreateCourseInput struct {
	Name      string `json:"name" binding:"required"`
	TeacherID uint   `json:"teacher_id" binding:"required"`
}

func CreateCourse(c *gin.Context) {
	var input CreateCourseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Verify teacher exists and is actually a teacher
	var teacher models.User
	if err := config.DB.First(&teacher, input.TeacherID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Teacher not found")
		return
	}
	if teacher.Role != "teacher" {
		utils.ErrorResponse(c, http.StatusBadRequest, "Assigned user is not a teacher")
		return
	}

	course := models.Course{
		Name:      input.Name,
		TeacherID: input.TeacherID,
	}

	if err := config.DB.Create(&course).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create course. Ensure course name is unique.")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Course created successfully", course)
}

// ListStudents returns all students with basic pagination
func ListStudents(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	var students []models.User
	if err := config.DB.Where("role = ?", "student").Limit(limit).Offset(offset).Find(&students).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch students")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Students fetched successfully", students)
}

// ListCourses returns all courses with basic pagination
func ListCourses(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	var courses []models.Course
	// Preload Teacher to show who teaches it
	if err := config.DB.Preload("Teacher").Limit(limit).Offset(offset).Find(&courses).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch courses")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Courses fetched successfully", courses)
}
