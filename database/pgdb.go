package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is a global variable to hold db connection

var DB *gorm.DB

func InitialMigration(DNS string) *gorm.DB {
	DB, err := gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	if err := DB.SetupJoinTable(&Hackathon{}, "Teams", &HackathonTeams{}); err != nil {
		println(err.Error())
		panic("Failed to setup join table: HackathonTeams")
	}
	// if err := DB.SetupJoinTable(&ParticipantTeams{}); err != nil {
	// 	println(err.Error())
	// 	panic("Failed to setup join table: HackathonTeams")
	// }
	DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	DB.AutoMigrate(Organizer{}, Hackathon{}, Participant{}, Team{})
	return DB
}
