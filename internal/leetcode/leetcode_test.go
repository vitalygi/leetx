package leetcode

import (
	"testing"
)

func TestGetProblemByUrl(t *testing.T) {
	problem, err := GetProblemByURL("https://leetcode.com/problems/two-sum/")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if problem.Title != "Two Sum" {
		t.Errorf("expected title 'Two Sum', got %s", problem.Title)
	}
}
