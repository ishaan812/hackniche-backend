package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Question struct {
	QuestionTitle     string `json:"question__title"`
	QuestionTitleSlug string `json:"question__title_slug"`
	Difficulty        string `json:"difficulty"`
}

type Questions struct {
	StatStatusPairs []struct {
		Stat struct {
			QuestionTitle      string `json:"question__title"`
			QuestionTitleSlug  string `json:"question__title_slug"`
			QuestionFrontendID string `json:"frontend_question_id"`
			Difficulty         struct {
				Level int `json:"level"`
			} `json:"difficulty"`
		} `json:"stat"`
		Status string `json:"status"`
	} `json:"stat_status_pairs"`
}

func Evol(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	var evalReq EvalReq
	json.NewDecoder(r.Body).Decode(&evalReq)
	evalRes := Evaluate(evalReq)
	json.NewEncoder(w).Encode(evalRes)
}

func GithubLast30Days(username string) int {
	url := fmt.Sprintf("https://api.github.com/users/%s/repos?per_page=100&page=1", username)

	req, err := http.NewRequest("GET", url, nil)
	// req.Header.Set("Authorization", fmt.Sprintf("token %s", accessToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var repos []struct {
		Name    string `json:"name"`
		HTMLURL string `json:"html_url"`
	}

	err = json.NewDecoder(resp.Body).Decode(&repos)
	if err != nil {
		panic(err)
	}

	// Calculate the sum of the user's contributions over the last 30 days
	contributions := 0
	// today := time.Now().Local().Format("2006-01-02")
	startDate := time.Now().Local().AddDate(0, 0, -30).Format("2006-01-02")
	for _, repo := range repos {
		url := fmt.Sprintf("https://api.github.com/repos/%s/%s/stats/participation", username, repo.Name)
		req, err := http.NewRequest("GET", url, nil)
		fmt.Println(startDate)
		if err != nil {
			fmt.Println(err)

		}
		// req.Header.Set("Authorization", fmt.Sprintf("token %s", accessToken))
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		var stats struct {
			All []int `json:"all"`
		}

		err = json.NewDecoder(resp.Body).Decode(&stats)
		if err != nil {
			panic(err)
		}

		// Find the index of the last 30 days in the stats array
		startIndex := -1
		for i, date := range stats.All {
			if true {
				fmt.Println(date)
				startIndex = i
				break
			}
		}
		if startIndex == -1 {
			panic("Could not find start date in stats")
		}

		endIndex := startIndex + 30
		if endIndex > len(stats.All) {
			endIndex = len(stats.All)
		}

		// Sum the contributions made in the last 30 days
		for i := startIndex; i < endIndex; i++ {
			contributions += stats.All[i]
		}
		return contributions
	}

	fmt.Printf("%s made %d contributions over the last 30 days\n", username, contributions)
	return contributions
}

func ResumeScore(EvalReq EvalReq) float64 {
	score := float64(0)
	m := make(map[string]interface{})
	for i := 0; i < len(EvalReq.HackathonDomains); i++ {
		m[EvalReq.HackathonDomains[i]] = 1
	}
	for i := 0; i < len(EvalReq.Skills); i++ {
		if m[EvalReq.Skills[i]] == 1 {
			score += 1
		}
	}
	score += EvalReq.Experience

	score += float64(len(EvalReq.Qualifications))

	return score
}

func Evaluate(EvalRequest EvalReq) EvalRes {
	GitScore := float64(GithubLast30Days(EvalRequest.GithubUsername))
	LeetcodeRank := EvalRequest.LeetcodeRank
	ResumeScore := ResumeScore(EvalRequest)
	var EvalResult EvalRes
	EvalResult.Score = GitScore * (1 / LeetcodeRank) * ResumeScore
	fmt.Println( GitScore , (1 / LeetcodeRank) , ResumeScore)
	// Get a list of the user's repositories

	return EvalResult
}
