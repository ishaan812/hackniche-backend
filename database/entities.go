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
	Image            string         `json:"image"`
	Email            string         `json:"email"`
	Password         string         `json:"password"`
	OrganizationName string         `json:"organization_name"`
}
type Hackathon struct {
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ID           uuid.UUID      `gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	Name         string         `json:"name"`
	Image        string         `json:"image"`
	Location     string         `json:"location"`
	Offline      bool           `json:"offline"`
	Criteria     string         `json:"criteria"`
	Deadline     string         `json:"deadline"`
	MinTeam      int            `json:"min_team"`
	MaxTeam      int            `json:"max_team"`
	StartDate    string         `json:"start_date"`
	EndDate      string         `json:"end_date"`
	OrganizerID  string         `json:"organizer_id"`
	Organizer    Organizer      `json:"organizer" gorm:"foreignKey:OrganizerID"`
	Description  string         `json:"description"`
	Teams        []*Team        `gorm:"many2many:hackathon_team;foreignKey:ID;joinForeignKey:HackathonID;" json:"teams"`
	Domains      pq.StringArray `gorm:"type:varchar(255)[]" json:"domains"`
	PrizePool    int            `json:"prize_pool"`
	Eligibility  int            `json:"eligibility"`
	Announcement string         `json:"announcenents"`
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
	Image            string         `json:"image"`
	Teams            []*Team        `gorm:"many2many:participant_team;foreignKey:ID;joinForeignKey:ParticipantID;" json:"teams"`
}
type Team struct {
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ID           uuid.UUID      `gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	Name         string         `json:"team_name"`
	Participants []*Participant `gorm:"many2many:participant_team;foreignKey:ID;joinForeignKey:TeamID;" json:"participants"`
}

type TeamsParticipant struct {
	CreatedAt      time.Time
	UpdatedAt      time.Time
	ID             uuid.UUID      `gorm:"primarykey;type:uuid;default:uuid_generate_v4()" json:"id"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
	ParticipantID  uuid.UUID      `json:"participant_id"`
	Participant    Participant    `json:"participant" gorm:"foreignKey:ParticipantID"`
	TeamID         uuid.UUID      `json:"team_id"`
	Team           Team           `json:"team" gorm:"foreignKey:TeamID"`
	TeamLeader     bool           `json:"team_leader"`
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
	Registered  bool           `json:"registered"`
	Accepted    bool           `json:"accepted"`
	// Messages    []*Message     `json:"messages"`
}

type Message struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	ID        uuid.UUID      `gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Sender    string         `json:"sender"`
	Content   string         `json:"content"`
}
