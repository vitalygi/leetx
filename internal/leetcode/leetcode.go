package leetcode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Request struct {
	Query     string      `json:"query"`
	Variables interface{} `json:"variables"`
}

type ProblemDetail struct {
	Data struct {
		Question struct {
			Title         string `json:"title"`
			TitleSlug     string `json:"titleSlug"`
			QuestionId    string `json:"questionId"`
			Content       string `json:"content"`
			QuestionTitle string `json:"questionTitle"`
			Difficulty    string `json:"difficulty"`
			TopicTags     []struct {
				Name string `json:"name"`
				Slug string `json:"slug"`
			} `json:"topicTags"`
		} `json:"question"`
	} `json:"data"`
}

func GetProblem(titleSlug string) (ProblemDetail, error) {
	url := "https://leetcode.com/graphql/"

	query := `
		query questionDetail($titleSlug: String!) {
			question(titleSlug: $titleSlug) {
				title
				titleSlug
				content
				difficulty
				questionTitle
				questionId
			}
		}
	`

	variables := map[string]interface{}{
		"titleSlug": titleSlug,
	}

	requestBody := Request{
		Query:     query,
		Variables: variables,
	}

	requestBodyBytes, _ := json.Marshal(requestBody)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		fmt.Println("Cannot create request", err)
		return ProblemDetail{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Cannot get problem:", err)
		return ProblemDetail{}, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Cannot read response:", err)
		return ProblemDetail{}, err
	}

	var response ProblemDetail
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		fmt.Println("Cannot parse response:", err)
		return ProblemDetail{}, err
	}

	return response, nil

}
