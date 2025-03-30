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

func createMainFile(dirName string, language string, mainFileName string) *os.File {
	var fileName string
	if mainFileName != "" {
		fileName = mainFileName
	} else {
		fileName, _ = langCommonFile[strings.ToLower(language)]
	}

	if fileName != "" {
		return createFile(filepath.Join(".", dirName, fileName))
	} else {
		return nil
	}

}

func createFile(name string) *os.File {
	file, err := os.Create(name)
	if err != nil && os.IsNotExist(err) {
		fmt.Println(ErrWhileCreateFile, err)
		return nil
	}
	return file
}

func PrepareWorkspace(problem leetcode.Problem, language string, mainFileName string) error {
	normalizedProblemTitle := strings.Replace(problem.QuestionTitle, " ", "_", -1)
	dirName := fmt.Sprintf("%v.%v", problem.QuestionId, normalizedProblemTitle)
	err := os.Mkdir(filepath.Join(".", dirName), os.ModePerm)
	if err != nil && os.IsNotExist(err) {
		return ErrWhileCreateDir
	}
	codeFile := createMainFile(dirName, language, mainFileName)
	if codeFile != nil && language != "" {
		codeSnippet, isFound := problem.GetCodeSnippet(language)
		if isFound {
			fmt.Fprintln(codeFile, fmt.Sprintf("\n\n%v", codeSnippet.Code))
		}
	}

	if problem.Content != "" {
		problemFile := createFile(filepath.Join(".", dirName, "problem.md"))
		if problemFile != nil {
			fmt.Fprintf(problemFile, "%s\n%s", problem.GetURL(), problem.Content)
		}
	} else {
		fmt.Println("Problem description is empty")
	}
	fmt.Print(problem.GetInfo())
	return nil
}
