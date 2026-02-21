package controllers

import (
	"grade-management-system/config"
	"grade-management-system/models"
	"grade-management-system/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAssignedCourses returns courses assigned to the logged-in teacher
func GetAssignedCourses(c *gin.Context) {
	teacherID := c.MustGet("userID").(uint)

	// Basic pagination
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	var courses []models.Course
	if err := config.DB.Where("teacher_id = ?", teacherID).Limit(limit).Offset(offset).Find(&courses).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch courses")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Courses fetched successfully", courses)
}

type EnrollStudentInput struct {
	StudentID uint `json:"student_id" binding:"required"`
	CourseID  uint `json:"course_id" binding:"required"`
}

// EnrollStudent enrolls a student to a course (teacher must own the course)
func EnrollStudent(c *gin.Context) {
	teacherID := c.MustGet("userID").(uint)

	var input EnrollStudentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Verify course belongs to this teacher
	var course models.Course
	if err := config.DB.First(&course, input.CourseID).Error; err != nil || course.TeacherID != teacherID {
		utils.ErrorResponse(c, http.StatusForbidden, "Course not found or you don't have access")
		return
	}

	enrollment := models.Enrollment{
		StudentID: input.StudentID,
		CourseID:  input.CourseID,
	}

	if err := config.DB.Create(&enrollment).Error; err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to enroll student. They may already be enrolled.")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Student enrolled successfully", enrollment)
}

type GradeInput struct {
	StudentID uint    `json:"student_id" binding:"required"`
	CourseID  uint    `json:"course_id" binding:"required"`
	Marks     float64 `json:"marks" binding:"min=0,max=100"`
}

func calculateGradeLetter(marks float64) string {
	if marks >= 90 {
		return "A"
	} else if marks >= 80 {
		return "B"
	} else if marks >= 70 {
		return "C"
	} else if marks >= 60 {
		return "D"
	}
	return "F"
}

// AddOrUpdateGrade allows teacher to grade a student in their course
func AddOrUpdateGrade(c *gin.Context) {
	teacherID := c.MustGet("userID").(uint)

	var input GradeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Verify course belongs to this teacher
	var course models.Course
	if err := config.DB.First(&course, input.CourseID).Error; err != nil || course.TeacherID != teacherID {
		utils.ErrorResponse(c, http.StatusForbidden, "Course not found or you don't have access")
		return
	}

	// Check if enrollment exists
	var enrollment models.Enrollment
	if err := config.DB.Where("student_id = ? AND course_id = ?", input.StudentID, input.CourseID).First(&enrollment).Error; err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Student is not enrolled in this course")
		return
	}

	gradeLetter := calculateGradeLetter(input.Marks)

	// Check if grade already exists
	var grade models.Grade
	err := config.DB.Where("student_id = ? AND course_id = ?", input.StudentID, input.CourseID).First(&grade).Error
	
	if err != nil {
		// Grade doesn't exist, create it
		grade = models.Grade{
			StudentID:   input.StudentID,
			CourseID:    input.CourseID,
			Marks:       input.Marks,
			GradeLetter: gradeLetter,
		}
		if createErr := config.DB.Create(&grade).Error; createErr != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to add grade")
			return
		}
		utils.SuccessResponse(c, http.StatusCreated, "Grade added successfully", grade)
	} else {
		// Grade exists, update it
		grade.Marks = input.Marks
		grade.GradeLetter = gradeLetter
		if updateErr := config.DB.Save(&grade).Error; updateErr != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update grade")
			return
		}
		utils.SuccessResponse(c, http.StatusOK, "Grade updated successfully", grade)
	}
}

// GetGradeStatistics returns count of A, B, C, D, F for a specific course
func GetGradeStatistics(c *gin.Context) {
	teacherID := c.MustGet("userID").(uint)
	courseIDParam := c.Param("courseId")
	courseID, err := strconv.Atoi(courseIDParam)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid course ID")
		return
	}

	// Check access
	var course models.Course
	if err := config.DB.First(&course, courseID).Error; err != nil || course.TeacherID != teacherID {
		utils.ErrorResponse(c, http.StatusForbidden, "Course not found or access denied")
		return
	}

	type StatsResult struct {
		GradeLetter string `json:"grade_letter"`
		Count       int    `json:"count"`
	}
	var results []StatsResult

	if err := config.DB.Model(&models.Grade{}).
		Select("grade_letter, count(id) as count").
		Where("course_id = ?", courseID).
		Group("grade_letter").
		Scan(&results).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to compute statistics")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Grade statistics retrieved", results)
}
