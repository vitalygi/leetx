package leetcode

import "strings"

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

type CodeSnippet struct {
	Code     string `json:"code"`
	Lang     string `json:"lang"`
	LangSlug string `json:"langSlug"`
}
