package models

import "time"

type Application struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Company     string    `json:"company"`
	Position    string    `json:"position"`
	Status      string    `json:"status"` // e.g., "Applied", "Interview", etc.
	Location    string    `json:"location"`
	AppliedDate time.Time `json:"applied_date"`
	ResumeURL   string    `json:"resume_url"`
	UserID      uint      `json:"user_id"`                                 // Set automatically by server
	User        User      `gorm:"foreignKey:UserID" json:"user,omitempty"` // Only in responses
}
