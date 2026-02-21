package controllers

import (
	"grade-management-system/config"
	"grade-management-system/models"
	"grade-management-system/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetStudentCourses returns courses the student is enrolled in
func GetStudentCourses(c *gin.Context) {
	studentID := c.MustGet("userID").(uint)

	var enrollments []models.Enrollment
	if err := config.DB.Preload("Course").Where("student_id = ?", studentID).Find(&enrollments).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch enrolled courses")
		return
	}

	// Extract courses from enrollments for a cleaner response
	var courses []models.Course
	for _, e := range enrollments {
		courses = append(courses, e.Course)
	}

	utils.SuccessResponse(c, http.StatusOK, "Enrolled courses retrieved", courses)
}

// GetStudentGrades returns all grades for the student
func GetStudentGrades(c *gin.Context) {
	studentID := c.MustGet("userID").(uint)

	var grades []models.Grade
	if err := config.DB.Preload("Course").Where("student_id = ?", studentID).Find(&grades).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch grades")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Grades retrieved", grades)
}

// GetStudentGPA calculates and returns the student's GPA
// A=4, B=3, C=2, D=1, F=0
func GetStudentGPA(c *gin.Context) {
	studentID := c.MustGet("userID").(uint)

	var grades []models.Grade
	if err := config.DB.Where("student_id = ?", studentID).Find(&grades).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch grades for GPA calculation")
		return
	}

	if len(grades) == 0 {
		utils.SuccessResponse(c, http.StatusOK, "No grades available to calculate GPA", gin.H{"gpa": 0.0})
		return
	}

	totalPoints := 0.0
	for _, grade := range grades {
		switch grade.GradeLetter {
		case "A":
			totalPoints += 4
		case "B":
			totalPoints += 3
		case "C":
			totalPoints += 2
		case "D":
			totalPoints += 1
		case "F":
			totalPoints += 0
		}
	}

	gpa := totalPoints / float64(len(grades))

	utils.SuccessResponse(c, http.StatusOK, "GPA calculated successfully", gin.H{
		"gpa":           gpa,
		"courses_count": len(grades),
	})
}
