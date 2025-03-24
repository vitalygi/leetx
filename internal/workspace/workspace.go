package workspace

import (
	"LeetX/internal/leetcode"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	ErrWhileCreateDir  = errors.New("cannot create dir for problem")
	ErrWhileCreateFile = errors.New("cannot create file for problem")
)

func createFile(name string) *os.File {
	file, err := os.Create(name)
	if err != nil && os.IsNotExist(err) {
		fmt.Println(ErrWhileCreateFile, err)
		return nil
	}
	return file
}

func PrepareWorkspace(problem leetcode.ProblemDetail, language string) error {
	normalizedProblemTitle := strings.Replace(problem.Data.Question.QuestionTitle, " ", "_", -1)
	dirName := fmt.Sprintf("%v.%v", problem.Data.Question.QuestionId, normalizedProblemTitle)
	err := os.Mkdir(filepath.Join(".", dirName), os.ModePerm)
	if err != nil && os.IsNotExist(err) {
		fmt.Println(ErrWhileCreateDir, err)
		return err
	}
	codeFile := createFile(filepath.Join(".", dirName, "main_test.go"))
	if codeFile != nil && language != "" {
		codeSnippet, isFound := problem.GetCodeSnippet(language)
		if isFound {
			fmt.Fprintln(codeFile, fmt.Sprintf("\n\n%v", codeSnippet.Code))
		}
	}

	if problem.Data.Question.Content != "" {
		problemFile := createFile(filepath.Join(".", dirName, "problem.md"))
		if problemFile != nil {
			fmt.Fprintln(problemFile, problem.Data.Question.Content)
		}
	}
	return nil
}
