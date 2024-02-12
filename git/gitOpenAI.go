package git

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/vinodhalaharvi/gitai/config"
	ai "github.com/vinodhalaharvi/gitai/openai"
	"github.com/vinodhalaharvi/gitai/utils"
	os "os"
	"os/exec"
	"regexp"
	"strings"
)

type GitAi struct {
	GitBinPath string
}

func NewGitAi() *GitAi {

	gitAi := GitAi{}
	gitAi.GitBinPath = config.C.Git.Command.Path
	return &gitAi
}

func (g *GitAi) validateCommand(cmd string) bool {
	pattern := config.C.Git.Validation.Pattern
	matched, _ := regexp.MatchString(pattern, cmd)
	return matched
}

func (g *GitAi) InterpretCommand(
	c context.Context,
	command string,
) (string, error) {
	return interpretCommand(c, command)
}

func (g *GitAi) ApplyCommand(ctx context.Context, command string) (string, error) {
	updatedPath := os.Getenv("PATH") + ":" + g.GitBinPath

	cmd := exec.CommandContext(ctx, "bash", "-c", command)
	cmd.Env = append(os.Environ(), "PATH="+updatedPath)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%sfailed to execute command:%s %v, output: %s", err, string(output))
	}

	return string(output), nil
}

func interpretCommand(ctx context.Context, command string) (string, error) {
	// Call the helper function from utils.go
	respBytes, err := utils.CallOpenAI(config.C.Git.Prompt.Template, command)
	if err != nil {
		return "", err
	}

	// Parse the response
	var response ai.OpenAIResponse
	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		return "", err
	}

	if len(response.Choices) > 0 {
		return strings.TrimSpace(response.Choices[0].Message.Content), nil
	}

	return "", fmt.Errorf("no response from OpenAI")

}
