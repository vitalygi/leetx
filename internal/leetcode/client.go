package leetcode

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

var (
	ErrNotCorrectPrefix = errors.New("URL should starts with https://leetcode.com/problems/")
	ErrIncorrectURL     = errors.New("incorrect leetcode url")
	ErrProblemNotFound  = errors.New("problem not found")
)

type LeetCodeClient interface {
	GetProblem(ctx context.Context, query string) (Problem, error)
	getProblemByURL(ctx context.Context, url string) (Problem, error)
	getProblemByQuery(ctx context.Context, query string) (Problem, error)
}

type client struct{}

func NewClient() LeetCodeClient {
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
	problems := response.Data.Problemset.Problems
	if len(problems) > 0 {
		return problems[0], err
	} else {
		if err != nil {
			return Problem{}, err
		} else {
			return Problem{}, ErrProblemNotFound
		}

	}
}

func fetchProblem[T RawProblemResponse | RawProblemSetQuestionListResponse](ctx context.Context, queryVars map[string]interface{}, questionDetailQuery string) (T, error) {
	url := leetCodeGraphQLURL
	query := questionDetailQuery
	var result T
	requestBody := Request{
		Query:     query,
		Variables: queryVars,
	}

	requestBodyBytes, _ := json.Marshal(requestBody)
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return result, fmt.Errorf("cannot create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return result, fmt.Errorf("timeout exceeded =(, maybe no internet?:%v", err)
		} else {
			return result, fmt.Errorf("cannot get problem:%v", err)
		}

	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, fmt.Errorf("cannot read response:%v", err)
	}

	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return result, fmt.Errorf("cannot parse response:%v", err)
	}
	return result, nil

}
