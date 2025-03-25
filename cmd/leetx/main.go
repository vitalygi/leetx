package main

import (
	"LeetX/internal/leetcode"
	"LeetX/internal/workspace"
	"errors"
	"flag"
	"fmt"
)

var (
	ErrProblemNotFound = errors.New("problem not found")
	ErrEmptyURL        = errors.New("enter url of problem you need")
)

func main() {
	problemID := flag.String("get", "", "Enter correct URL to get problem")
	language := flag.String("l", "", "Enter programming language for code snippet")
	mainFileName := flag.String("f", "", "Enter file name for file with code snippet. "+
		"Default will be get from language")
	flag.Parse()
	if *problemID == "" {
		fmt.Println(ErrEmptyURL)
		return
	}
	var problem leetcode.Problem
	var err error

	problem, err = leetcode.GetProblemByURL(*problemID)
	if err != nil {
		return
	}
	if problem.QuestionId == "" {
		fmt.Println(ErrProblemNotFound)
		return
	}
	err = workspace.PrepareWorkspace(problem, *language, *mainFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
}
