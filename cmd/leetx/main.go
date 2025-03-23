package main

import (
	"LeetX/internal/leetcode"
	"LeetX/internal/workspace"
	"errors"
	"flag"
	"fmt"
	"strings"
)

var (
	ErrNotCorrectPrefix = errors.New("URL should starts with https://leetcode.com/problems/")
	ErrIncorrectURL     = errors.New("incorrect leetcode url")
	ErrProblemNotFound  = errors.New("problem not found")
	ErrEmptyURL         = errors.New("enter url of problem you need")
)

func checkURLPrefix(url string) error {
	if strings.HasPrefix(url, "https://leetcode.com/problems/") {
		return nil
	} else {
		return ErrNotCorrectPrefix
	}
}

func getProblemTitle(url string) (string, error) {
	urlParts := strings.Split(url, "/")
	if len(urlParts) > 4 {
		return urlParts[4], nil
	} else {
		return "", ErrIncorrectURL
	}

}

func main() {
	problemURL := flag.String("get", "", "Enter correct URL to get problem")
	language := flag.String("l", "", "Enter programming language for start code")
	flag.Parse()
	if *problemURL == "" {
		fmt.Println(ErrEmptyURL)
		return
	}
	err := checkURLPrefix(*problemURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	problemTitle, err := getProblemTitle(*problemURL)
	if err != nil {
		fmt.Println(ErrIncorrectURL)
		return
	}
	problem, err := leetcode.GetProblem(problemTitle)
	if err != nil {
		return
	}
	if problem.Data.Question.QuestionId == "" {
		fmt.Println(ErrProblemNotFound)
		return
	}
	err = workspace.PrepareWorkspace(problem, *language)
	if err != nil {
		fmt.Println(err)
		return
	}

}
