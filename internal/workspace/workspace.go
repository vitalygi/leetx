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

func PrepareWorkspace(problem leetcode.ProblemDetail) error {
	normalizedProblemTitle := strings.Replace(problem.Data.Question.QuestionTitle, " ", "_", -1)
	dirName := fmt.Sprintf("%v.%v", problem.Data.Question.QuestionId, normalizedProblemTitle)
	err := os.Mkdir(filepath.Join(".", dirName), os.ModePerm)
	if err != nil && os.IsNotExist(err) {
		fmt.Println(ErrWhileCreateDir, err)
		return err
	}
	_, err = os.Create(filepath.Join(".", dirName, "main_test.go"))
	if err != nil && os.IsNotExist(err) {
		fmt.Println(ErrWhileCreateFile, err)
		return err
	}
	return nil
}
