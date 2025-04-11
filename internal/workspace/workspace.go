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
	ErrCreateDir        = errors.New("cannot create dir")
	ErrCreateFile       = errors.New("cannot create file")
	ErrLanguageNotFound = errors.New("language not found")
)

// createSnippetFile creates a file for the given language snippet in the specified directory.
// It uses a default filename if mainFileName is empty.
func createSnippetFile(dirName string, language string, mainFileName string) (*os.File, error) {
	var fileName string
	var ok bool
	if mainFileName != "" {
		fileName = mainFileName
	} else {
		fileName, ok = langCommonFile[strings.ToLower(language)]
		if !ok {
			return nil, fmt.Errorf("failed to find file name for language: %w", ErrLanguageNotFound)
		}
	}

	if fileName != "" {
		return createFile(filepath.Join(".", dirName, fileName))
	} else {
		return nil, fmt.Errorf("failed to find file name for language: %w", ErrLanguageNotFound)
	}

}

// createFile creates file. If failed, returns error only if it's not os.ErrNotExist
func createFile(name string) (*os.File, error) {
	file, err := os.Create(name)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrCreateFile, err)
	}
	return file, nil
}

func PrepareWorkspace(problem leetcode.Problem, language string, mainFileName string) error {
	if language == "" {
		fmt.Print(problem.GetInfo())
		return nil
	}

	// create dir for problem context
	normalizedProblemTitle := strings.Replace(problem.QuestionTitle, " ", "_", -1)
	dirName := fmt.Sprintf("%v.%v", problem.QuestionId, normalizedProblemTitle)
	err := os.Mkdir(filepath.Join(".", dirName), os.ModePerm)
	if err != nil && os.IsNotExist(err) {
		return ErrCreateDir
	}

	// create file with problem's code snippet
	codeSnippet, isFound := problem.GetCodeSnippet(language)
	if isFound {
		codeFile, err := createSnippetFile(dirName, language, mainFileName)
		if err != nil {
			return fmt.Errorf("cannot create snippet file: %v", err)
		} else {
			_, err := fmt.Fprintln(codeFile, fmt.Sprintf("\n\n%v", codeSnippet.Code))
			if err != nil {
				return fmt.Errorf("cannot create snippet file: %v", err)
			}
		}
	} else {
		fmt.Println("cannot create snippet file, language not found")
	}

	// create file problem.md with problem description
	if problem.Content != "" {
		problemFile, err := createFile(filepath.Join(".", dirName, "problem.md"))
		if err != nil {
			return fmt.Errorf("failed to create file for problem description %w", err)
		}
		_, err = fmt.Fprintf(problemFile, "%s\n%s", problem.GetURL(), problem.Content)
		if err != nil {
			return fmt.Errorf("failed to create file for problem description %w", err)
		}
	} else {
		fmt.Println("Problem description is empty")
	}

	fmt.Print(problem.GetInfo())
	return nil
}
