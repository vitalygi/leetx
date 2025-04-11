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

// checkURLPrefix validates URL
func checkURLPrefix(url string) error {
	if strings.HasPrefix(url, "https://leetcode.com/problems/") {
		return nil
	} else {
		return ErrNotCorrectPrefix
	}
}

// getProblemTitle extracts problem title from url
func getProblemTitle(url string) (string, error) {
	urlParts := strings.Split(url, "/")
	if len(urlParts) > 4 {
		return urlParts[4], nil
	} else {
		return "", ErrIncorrectURL
	}
}

// fetchProblem fetches problem from leetcode and returns necessary struct.
// Requires queryVars and questionDetailQuery (graphql query)
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
