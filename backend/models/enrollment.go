package models

// Enrollment represents a student enrolled in a course
// Includes unique index to prevent duplicate enrollments
type Enrollment struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	StudentID uint `gorm:"uniqueIndex:idx_student_course;not null" json:"student_id"`
	CourseID  uint `gorm:"uniqueIndex:idx_student_course;not null" json:"course_id"`
	
	// Relationships
	Student User   `gorm:"foreignKey:StudentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"student,omitempty"`
	Course  Course `gorm:"foreignKey:CourseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"course,omitempty"`
}
