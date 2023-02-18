package controllers

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v4"
	"github.com/lib/pq"
)

type Claims struct {
	SAPID      string    `json:"sap_id"`
	UserID     int       `json:"user_id"`
	Email      string    `json:"email"`
	Name       string    `json:"name"`
	Department string    `json:"department"`
	Expires    time.Time `json:"expires"`
	jwt.RegisteredClaims
}

type AddTeamToHackathonReq struct {
	TeamID      uuid.UUID `json:"team_id"`
	HackathonID uuid.UUID `json:"hackathon_id"`
}

type EvalReq struct {
	HackathonDomains pq.StringArray `json:"hackathon_domains"`
	Skills           pq.StringArray `json:"skills"`
	Experience       float64        `json:"experience"`
	Qualifications   pq.StringArray `json:"qualifications"`
	LeetcodeRank     float64        `json:"leetcode_rank"`
	GithubUsername   string         `json:"github_username"`
}

type EvalRes struct {
	Score float64 `json:"score"`
}
