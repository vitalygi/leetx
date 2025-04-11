package leetcode

import (
	"fmt"
	"strings"
)

type Request struct {
	Query     string      `json:"query"`
	Variables interface{} `json:"variables"`
}

type RawProblemSetQuestionListResponse struct {
	Data struct {
		Problemset struct {
			Problems []Problem `json:"questions"`
		} `json:"problemsetQuestionList"`
	} `json:"data"`
}
type RawProblemResponse struct {
	Data struct {
		Problem Problem `json:"question"`
	} `json:"data"`
}

type CodeSnippet struct {
	Code     string `json:"code"`
	Lang     string `json:"lang"`
	LangSlug string `json:"langSlug"`
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

func (problem *Problem) GetURL() string {
	return fmt.Sprintf("https://leetcode.com/problems/%s", problem.TitleSlug)
}

func (problem *Problem) GetInfo() string {
	return fmt.Sprintf(
		"Problem: %s.%s\nDifficulty: %s\nURL:%s\n",
		problem.QuestionId,
		problem.Title,
		problem.Difficulty,
		problem.GetURL())
}
