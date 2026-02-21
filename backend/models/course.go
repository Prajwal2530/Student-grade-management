package models

// Course represents a class taught by a Teacher
type Course struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"not null;unique" json:"name"`
	TeacherID uint   `gorm:"not null" json:"teacher_id"`
	// Relationship
	Teacher   User   `gorm:"foreignKey:TeacherID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"teacher,omitempty"`
}
