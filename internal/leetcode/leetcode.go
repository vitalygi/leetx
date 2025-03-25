package leetcode

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var (
	ErrNotCorrectPrefix = errors.New("URL should starts with https://leetcode.com/problems/")
	ErrIncorrectURL     = errors.New("incorrect leetcode url")
)

const (
	leetCodeGraphQLURL  = "https://leetcode.com/graphql/"
	questionDetailQuery = `
        query questionDetail($titleSlug: String!) {
            question(titleSlug: $titleSlug) {
                title
                titleSlug
                content
                questionFrontendId
                difficulty
                questionTitle
                codeSnippets {
                    lang
                    langSlug
                    code
                }
                topicTags {
                    name
                    slug
                }
            }
        }
    `
)

type Request struct {
	Query     string      `json:"query"`
	Variables interface{} `json:"variables"`
}

type RawProblemResponse struct {
	Data struct {
		Problem Problem `json:"question"`
	} `json:"data"`
}
type Problem struct {
	Title         string `json:"title"`
	TitleSlug     string `json:"titleSlug"`
	QuestionId    string `json:"questionFrontendId"`
	Content       string `json:"content"`
	QuestionTitle string `json:"questionTitle"`
	Difficulty    string `json:"difficulty"`
	TopicTags     []struct {
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"topicTags"`
	CodeSnippets []CodeSnippet `json:"codeSnippets"`
}

type CodeSnippet struct {
	Code     string `json:"code"`
	Lang     string `json:"lang"`
	LangSlug string `json:"langSlug"`
}

func (problem *Problem) GetCodeSnippet(langSlug string) (CodeSnippet, bool) {
	langSlug = strings.ToLower(langSlug)
	for _, snippet := range problem.CodeSnippets {
		isNecessaryLangSlug := strings.ToLower(snippet.LangSlug) == langSlug
		isNecessaryLang := strings.ToLower(snippet.Lang) == langSlug
		if isNecessaryLangSlug || isNecessaryLang {
			return snippet, true
		}
	}
	return CodeSnippet{}, false
}

func checkURLPrefix(url string) error {
	if strings.HasPrefix(url, "https://leetcode.com/problems/") {
		return nil
	} else {
		return ErrNotCorrectPrefix
	}
}

func getProblemTitle(url string) (string, error) {
	urlParts := strings.Split(url, "/")
	if len(urlParts) > 4 {
		return urlParts[4], nil
	} else {
		return "", ErrIncorrectURL
	}

}
func GetProblemByURL(url string) (Problem, error) {
	err := checkURLPrefix(url)
	if err != nil {
		fmt.Println(err)
		return Problem{}, err
	}

	problemTitle, err := getProblemTitle(url)
	if err != nil {
		fmt.Println(ErrIncorrectURL)
		return Problem{}, err
	}
	return getProblem(map[string]interface{}{
		"titleSlug": problemTitle}, questionDetailQuery)
}

func getProblem(queryVars map[string]interface{}, questionDetailQuery string) (Problem, error) {
	url := leetCodeGraphQLURL
	query := questionDetailQuery

	requestBody := Request{
		Query:     query,
		Variables: queryVars,
	}

	requestBodyBytes, _ := json.Marshal(requestBody)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		fmt.Println("Cannot create request", err)
		return Problem{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Cannot get problem:", err)
		return Problem{}, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Cannot read response:", err)
		return Problem{}, err
	}
	var response RawProblemResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		fmt.Println("Cannot parse response:", err)
		return Problem{}, err
	}

	return response.Data.Problem, nil

}
