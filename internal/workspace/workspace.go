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

func PrepareWorkspace(problem leetcode.ProblemDetail, language string) error {
	normalizedProblemTitle := strings.Replace(problem.Data.Question.QuestionTitle, " ", "_", -1)
	dirName := fmt.Sprintf("%v.%v", problem.Data.Question.QuestionId, normalizedProblemTitle)
	err := os.Mkdir(filepath.Join(".", dirName), os.ModePerm)
	if err != nil && os.IsNotExist(err) {
		fmt.Println(ErrWhileCreateDir, err)
		return err
	}
	file, err := os.Create(filepath.Join(".", dirName, "main_test.go"))
	if err != nil && os.IsNotExist(err) {
		fmt.Println(ErrWhileCreateFile, err)
		return err
	}
	if language != "" {
		codeSnippet, isFound := problem.GetCodeSnippet(language)
		if isFound {
			fmt.Fprintln(file, fmt.Sprintf("\n\n%v", codeSnippet.Code))
		}
	}
	return nil
}
