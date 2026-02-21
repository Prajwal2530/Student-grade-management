package main

import (
	"log"

	"grade-management-system/config"
	"grade-management-system/models"
	"grade-management-system/utils"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	config.ConnectDatabase()

	log.Println("Starting to seed database...")

	// 1. Seed Admin
	adminPassword, _ := utils.HashPassword("admin123")
	admin := models.User{
		Name:     "Dr. Vikram Sharma (HOD)",
		Email:    "hod.cse@university.edu.in",
		Password: adminPassword,
		Role:     "admin",
	}
	if err := config.DB.Where("email = ?", admin.Email).FirstOrCreate(&admin).Error; err != nil {
		log.Fatal("Failed to seed admin:", err)
	}

	// 2. Seed Teacher
	teacherPassword, _ := utils.HashPassword("teacher123")
	teacher := models.User{
		Name:     "Prof. Anjali Desai",
		Email:    "anjali.desai@university.edu.in",
		Password: teacherPassword,
		Role:     "teacher",
	}
	config.DB.Where("email = ?", teacher.Email).FirstOrCreate(&teacher)

	// 3. Seed Student
	studentPassword, _ := utils.HashPassword("student123")
	student := models.User{
		Name:     "Rahul Verma",
		Email:    "rahul.verma@student.edu.in",
		Password: studentPassword,
		Role:     "student",
	}
	config.DB.Where("email = ?", student.Email).FirstOrCreate(&student)

	// 4. Seed Course
	course := models.Course{
		Name:      "Data Structures and Algorithms",
		TeacherID: teacher.ID,
	}
	config.DB.Where("name = ?", course.Name).FirstOrCreate(&course)

	// 5. Seed Enrollment
	enrollment := models.Enrollment{
		StudentID: student.ID,
		CourseID:  course.ID,
	}
	config.DB.Where("student_id = ? AND course_id = ?", student.ID, course.ID).FirstOrCreate(&enrollment)

	// 6. Add a default grade
	grade := models.Grade{
		StudentID:   student.ID,
		CourseID:    course.ID,
		Marks:       85.5,
		GradeLetter: "B",
	}
	config.DB.Where("student_id = ? AND course_id = ?", student.ID, course.ID).FirstOrCreate(&grade)

	log.Println("Database seeded successfully! You can now log in.")
	log.Println("Admin: hod.cse@university.edu.in / admin123")
	log.Println("Teacher: anjali.desai@university.edu.in / teacher123")
	log.Println("Student: rahul.verma@student.edu.in / student123")
}
