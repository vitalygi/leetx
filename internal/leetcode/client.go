package leetcode

import (
	"context"
	"errors"
	"strings"
)

var (
	ErrNotCorrectPrefix = errors.New("URL should starts with https://leetcode.com/problems/")
	ErrIncorrectURL     = errors.New("incorrect leetcode url")
	ErrProblemNotFound  = errors.New("problem not found")
)

type Client interface {
	GetProblem(ctx context.Context, query string) (Problem, error)
	getProblemByURL(ctx context.Context, url string) (Problem, error)
	getProblemByQuery(ctx context.Context, query string) (Problem, error)
}

type client struct{}

func NewClient() Client {
	return &client{}
}

func (c *client) GetProblem(ctx context.Context, query string) (Problem, error) {
	var problem Problem
	var err error
	if strings.HasPrefix(query, "https://") {
		problem, err = c.getProblemByURL(ctx, query)
	} else {
		problem, err = c.getProblemByQuery(ctx, query)
	}
	return problem, err
}

func (c *client) getProblemByURL(ctx context.Context, url string) (Problem, error) {
	err := checkURLPrefix(url)
	if err != nil {
		return Problem{}, err
	}

	problemTitle, err := getProblemTitle(url)
	if err != nil {
		return Problem{}, ErrIncorrectURL
	}
	response, err := fetchProblem[RawProblemResponse](ctx, map[string]interface{}{
		"titleSlug": problemTitle}, questionDetailQuery)
	if err != nil {
		return Problem{}, err
	}
	return response.Data.Problem, err
}

func (c *client) getProblemByQuery(ctx context.Context, query string) (Problem, error) {
	queryVars := map[string]interface{}{
		"categorySlug": "all-code-essentials",
		"skip":         0,
		"limit":        1,
		"filters": map[string]interface{}{
			"searchKeywords": query,
		},
	}

	response, err := fetchProblem[RawProblemSetQuestionListResponse](ctx, queryVars, problemsetQuestionListQuery)
	if err != nil {
		return Problem{}, err
	}
	problems := response.Data.Problemset.Problems
	if len(problems) > 0 {
		return problems[0], err
	} else {
		return Problem{}, ErrProblemNotFound
	}
}
