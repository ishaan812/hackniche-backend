package database

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

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
	StartDate   string         `json:"start_date"`
	EndDate     string         `json:"end_date"`
	OrganizerID string         `json:"organizer_id"`
	Organizer   Organizer      `json:"organizer" gorm:"foreignKey:OrganizerID"`
	Description string         `json:"description"`
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
type Team struct {
	Name      string `json:"name"`
	LeaderID  uuid.UUID `json:"team_leader_id"`
	Member1ID uuid.UUID `json:"member1_id"`
	Member2ID uuid.UUID `json:"member2_id"`
	Member3ID uuid.UUID `json:"member3_id"`
}
