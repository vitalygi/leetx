package leetcode

import (
	"context"
	"testing"
)

func TestGetProblemByUrl(t *testing.T) {
	client := NewClient()
	problem, err := client.GetProblemByURL(context.Background(), "https://leetcode.com/problems/two-sum/")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if problem.Title != "Two Sum" {
		t.Errorf("expected title 'Two Sum', got %s", problem.Title)
	}
}

func TestGetProblemByURL_InvalidURL(t *testing.T) {
	client := NewClient()
	_, err := client.GetProblemByURL(context.Background(), "invalid-url")
	if err == nil {
		t.Error("expected error for invalid URL, got nil")
	}
}
