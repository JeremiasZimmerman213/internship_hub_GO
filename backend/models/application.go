package models

import "time"

type ApplicationStatus uint8

const (
	StatusApplied ApplicationStatus = iota
	StatusOAReceived
	StatusInterviewing
	StatusAccepted
	StatusRejected
)

func (s ApplicationStatus) String() string {
	switch s {
	case StatusApplied:
		return "Applied"
	case StatusOAReceived:
		return "OA Received"
	case StatusInterviewing:
		return "Interviewing"
	case StatusAccepted:
		return "Accepted"
	case StatusRejected:
		return "Rejected"
	default:
		return "Unknown"
	}
}

type Application struct {
	ID          uint              `gorm:"primaryKey" json:"id"`
	Company     string            `json:"company"`
	Position    string            `json:"position"`
	Status      ApplicationStatus `json:"status"` // e.g., "Applied", "Interview", etc.
	Location    string            `json:"location"`
	AppliedDate time.Time         `json:"applied_date"`
	Term        string            `json:"term"`                            // e.g., "Summer 2025"
	Note        string            `gorm:"size:1048" json:"note,omitempty"` // Optional note, up to 1048 chars
	ResumeURL   string            `json:"resume_url"`
	UserID      uint              `json:"user_id"`                                 // Set automatically by server
	User        User              `gorm:"foreignKey:UserID" json:"user,omitempty"` // Only in responses
}
