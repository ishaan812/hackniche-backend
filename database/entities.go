package database

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/lib/pq"
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
	Teams       []*Team        `gorm:"many2many:hackathon_team;foreignKey:ID;joinForeignKey:HackathonID;" json:"teams"`
	Domains     pq.StringArray `gorm:"type:varchar(255)[]" json:"domains"`
	PrizePool   int            `json:"prize_pool"`
	Eligibility int            `json:"eligibility"`
}

type Participant struct {
	CreatedAt        time.Time
	UpdatedAt        time.Time
	ID               uuid.UUID      `gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
	Name             string         `json:"name"`
	Age              int            `json:"age"`
	College          string         `json:"college"`
	YearOfGraduation string         `json:"year_of_graduation"`
	Resume           string         `json:"resume"`
	Skills           string         `json:"skills"`
	Experience       string         `json:"experience"`
	Qualfications    string         `json:"qualifications"`
	MobileNumber     string         `json:"mobile_number"`
	Gender           string         `json:"gender"`
	Type             string         `json:"type"`
	Email            string         `json:"email"`
	Username         string         `json:"username"`
	Password         string         `json:"password"`
	Teams            []*Team        `gorm:"many2many:participant_team;foreignKey:ID;joinForeignKey:ParticipantID;" json:"teams"`
}
type Team struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	ID        uuid.UUID      `gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Name      string         `json:"name"`
	TeamCode  string         `json:"team_code"`
	LeaderID  uuid.UUID      `json:"team_leader_id"`
	Member1ID uuid.UUID      `json:"member1_id"`
	Member2ID uuid.UUID      `json:"member2_id"`
	Member3ID uuid.UUID      `json:"member3_id"`
}

type HackathonTeams struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ID          uuid.UUID      `gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	TeamID      uuid.UUID      `json:"team_id"`
	Team        Team           `json:"team" gorm:"foreignKey:TeamID"`
	HackathonID uuid.UUID      `json:"hackathon_id"`
	Hackathon   Hackathon      `json:"hackathon" gorm:"foreignKey:HackathonID"`
	Accepted    bool           `json:"accepted"`
}
