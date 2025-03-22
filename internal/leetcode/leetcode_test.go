package leetcode

import (
	"testing"
)

func TestGetProblem(t *testing.T) {
	problem, err := GetProblem("two-sum")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if problem.Data.Question.Title != "Two Sum" {
		t.Errorf("expected title 'Two Sum', got %s", problem.Data.Question.Title)
	}
}
