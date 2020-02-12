package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func main() {
	repoOwner := os.Getenv("REPOSITORY_OWNER")
	repoName := os.Getenv("REPOSITORY_NAME")
	issueNumber := os.Getenv("ISSUE_NUMBER")
	commentBody := os.Getenv("ISSUE_COMMENT_BODY")
	commentUser := os.Getenv("ISSUE_COMMENT_USER")
	commentURL := os.Getenv("ISSUE_COMMENTS_URL")

	fmt.Printf("%s/%s #%s\n", repoOwner, repoName, issueNumber)
	fmt.Printf("<%s> by %s\n", commentBody, commentUser)
	fmt.Printf("url: %s\n", commentURL)

	fmt.Printf("hello: %+v\n", errors.New("err"))
	fmt.Printf("Good bye!\n")
}
