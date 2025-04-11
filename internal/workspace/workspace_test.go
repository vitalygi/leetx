package workspace

import (
	"LeetX/internal/leetcode"
	"os"
	"path/filepath"
	"testing"
)

func TestPrepareWorkspace(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "lcget-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(originalDir)

	err = os.Chdir(tempDir)
	if err != nil {
		t.Fatal(err)
	}

	problem := leetcode.Problem{
		Title:         "Two Sum",
		TitleSlug:     "two-sum",
		QuestionId:    "1",
		Content:       "<p>Given an array of integers nums and an integer target...</p>",
		QuestionTitle: "Two Sum",
		Difficulty:    "Easy",
		TopicTags: []struct {
			Name string `json:"name"`
			Slug string `json:"slug"`
		}{
			{Name: "Array", Slug: "array"},
			{Name: "Hash Table", Slug: "hash-table"},
		},
		CodeSnippets: []leetcode.CodeSnippet{
			{
				Lang:     "go",
				LangSlug: "golang",
				Code:     `func twoSum(nums []int, target int) []int {}`,
			},
		},
	}

	err = PrepareWorkspace(problem, "go", "")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	dirName := "1.Two_Sum"
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		t.Errorf("directory %s does not exist", dirName)
	}

	codeFileName := filepath.Join(dirName, "main.go")
	if _, err := os.Stat(codeFileName); os.IsNotExist(err) {
		t.Errorf("file %s does not exist", codeFileName)
	}

	problemFileName := filepath.Join(dirName, "problem.md")
	if _, err := os.Stat(problemFileName); os.IsNotExist(err) {
		t.Errorf("file %s does not exist", problemFileName)
	}
}
