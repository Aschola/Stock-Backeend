package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"unique;not null"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	FirstName string    `gorm:"not null"`
	LastName  string    `gorm:"not null"`
	RoleID    uint      `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
