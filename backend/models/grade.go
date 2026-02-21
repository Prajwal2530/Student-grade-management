package models

// Grade represents the marks a student received in a course.
// Includes unique index so a student only has one grade per course.
type Grade struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	StudentID   uint   `gorm:"uniqueIndex:idx_grading;not null" json:"student_id"`
	CourseID    uint   `gorm:"uniqueIndex:idx_grading;not null" json:"course_id"`
	Marks       float64 `gorm:"not null;check:marks >= 0 AND marks <= 100" json:"marks"`
	GradeLetter string `gorm:"not null;size:1" json:"grade_letter"`
	
	// Relationships
	Student User   `gorm:"foreignKey:StudentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"student,omitempty"`
	Course  Course `gorm:"foreignKey:CourseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"course,omitempty"`
}
