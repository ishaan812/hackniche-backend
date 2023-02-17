package database

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Hackathon struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ID          uuid.UUID      `gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Name        string         `json:"name"`
	Location    string         `json:"location"`
	Offline     bool           `json:"offline"`
	Criteria    string         `json:"criteria"`
	Deadline    string         `json:"deadline"`
	MinTeam     int            `json:"min_team"`
	MaxTeam     int            `json:"max_team"`
	Information string         `json:"information"`
}

type Participant struct {
	CreatedAt        time.Time
	UpdatedAt        time.Time
	ID               uuid.UUID      `gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
	Name             string         `json:"name"`
	College          string         `json:"college"`
	YearOfGraduation string         `json:"year_of_graduation"`
	Resume           string         `json:"resume"`
	Skills           string         `json:"skills"`
	Experience       string         `json:"experience"`
	Qualfications    string         `json:"qualfications"`
	MobileNumber     string         `json:"mobile_number"`
	Gender           string         `json:"gender"`
	Type             string         `json:"type"`
	Email            string         `json:"email"`
}

type Organizer struct {
	CreatedAt        time.Time
	UpdatedAt        time.Time
	ID               uuid.UUID      `gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
	Name             string         `json:"name"`
	Email            string         `json:"email"`
	Password         string         `json:"password"`
	OrganizationName string         `json:"organization_name"`
}
