package main

import (
	"LeetX/internal/leetcode"
	"LeetX/internal/workspace"
	"context"
	"flag"
	"fmt"
)

func main() {
	problemID := flag.String("get", "", "URL or ID or title of the LeetCode problem to fetch (required)")
	language := flag.String("l", "", "Programming language for the code snippet (e.g., go, python)")
	mainFileName := flag.String("f", "", "File name for the code snippet (default based on language)")
	flag.Parse()
	if *problemID == "" || *mainFileName != "" && *language == "" {
		flag.Usage()
		return
	}

	client := leetcode.NewClient()

	problem, err := client.GetProblem(context.Background(), *problemID)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = workspace.PrepareWorkspace(problem, *language, *mainFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
}
