package routes

import (
	"grade-management-system/controllers"
	"grade-management-system/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// Public routes
	r.GET("/health", controllers.HealthCheck)
	r.POST("/register", controllers.RegisterStudent)
	r.POST("/login", controllers.Login)

	// Protected routes
	api := r.Group("/api")
	api.Use(middleware.AuthRequired())

	// Admin routes
	admin := api.Group("/admin")
	admin.Use(middleware.RoleRequired("admin"))
	{
		admin.POST("/users", controllers.CreateUser)
		admin.POST("/courses", controllers.CreateCourse)
		admin.GET("/students", controllers.ListStudents)
		admin.GET("/courses", controllers.ListCourses)
	}

	// Teacher routes
	teacher := api.Group("/teacher")
	teacher.Use(middleware.RoleRequired("teacher"))
	{
		teacher.GET("/courses", controllers.GetAssignedCourses)
		teacher.POST("/enrollments", controllers.EnrollStudent)
		teacher.POST("/grades", controllers.AddOrUpdateGrade)
		teacher.GET("/courses/:courseId/stats", controllers.GetGradeStatistics)
	}

	// Student routes
	student := api.Group("/student")
	student.Use(middleware.RoleRequired("student"))
	{
		student.GET("/courses", controllers.GetStudentCourses)
		student.GET("/grades", controllers.GetStudentGrades)
		student.GET("/gpa", controllers.GetStudentGPA)
	}

	return r
}
