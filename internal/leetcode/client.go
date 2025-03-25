package leetcode

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

var (
	ErrNotCorrectPrefix = errors.New("URL should starts with https://leetcode.com/problems/")
	ErrIncorrectURL     = errors.New("incorrect leetcode url")
	ErrProblemNotFound  = errors.New("problem not found")
)

type LeetCodeClient interface {
	GetProblemByURL(ctx context.Context, url string) (Problem, error)
}

type client struct{}

func NewClient() LeetCodeClient {
	return &client{}
}

func (c *client) GetProblemByURL(ctx context.Context, url string) (Problem, error) {
	err := checkURLPrefix(url)
	if err != nil {
		return Problem{}, err
	}

	problemTitle, err := getProblemTitle(url)
	if err != nil {
		return Problem{}, ErrIncorrectURL
	}
	return c.getProblem(ctx, map[string]interface{}{
		"titleSlug": problemTitle}, questionDetailQuery)
}

func (c *client) getProblem(ctx context.Context, queryVars map[string]interface{}, questionDetailQuery string) (Problem, error) {
	url := leetCodeGraphQLURL
	query := questionDetailQuery

	requestBody := Request{
		Query:     query,
		Variables: queryVars,
	}

	requestBodyBytes, _ := json.Marshal(requestBody)
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return Problem{}, fmt.Errorf("cannot create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return Problem{}, fmt.Errorf("timeout exceeded =(, maybe no internet?:%v", err)
		} else {
			return Problem{}, fmt.Errorf("cannot get problem:%v", err)
		}

	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Problem{}, fmt.Errorf("cannot read response:%v", err)
	}
	var response RawProblemResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return Problem{}, fmt.Errorf("cannot parse response:%v", err)
	}
	if response.Data.Problem.QuestionId == "" {
		return Problem{}, ErrProblemNotFound
	}
	return response.Data.Problem, nil

}
