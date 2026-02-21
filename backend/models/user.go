package models

import (
	"time"
)

// User represents Admin, Teacher, or Student in the system
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"` // Don't return password in JSON
	Role      string    `gorm:"not null;check:role IN ('admin', 'teacher', 'student')" json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
