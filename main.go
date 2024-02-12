package main

import (
	"context"
	"fmt"
	"github.com/vinodhalaharvi/gitai/config"
	"github.com/vinodhalaharvi/gitai/git"
	"log"
	"os"
	"strings"
)

func main() {
	// check if OPENAI_API_KEY is set in the environment

	if os.Getenv("OPENAI_API_KEY") == "" {
		fmt.Println("OPENAI_API_KEY environment variable not set")
		return

	}

	args := os.Args
	gitCommand := strings.Join(args[1:], " ")
	config.SetupConfig()

	// setup GitAi
	gitAi := git.NewGitAi()
	c := context.Background()
	command, err := gitAi.InterpretCommand(c, gitCommand)
	if err != nil {
		return
	}

	applyCommand, err := gitAi.ApplyCommand(c, command)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("\n%s\n", applyCommand)
}
