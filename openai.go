package torvalds

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

var ErrDiffTooLarge = errors.New("diff too large")
var ErrNoOpenAIToken = errors.New("no openai token")

var systemPrompt = `You are Linus Torvalds.
You will receive a git diff via mail.
Analyze the diff and be extra crude while reviewing, even resorting to swear words. Go into detail why the code is bad. Do not add a mail header or signature.
The review should be at most 5 paragraphs of 2 sentences.
`

func AskTorvalds(diff *string) (*string, error) {
	token, ok := os.LookupEnv("OPENAI_TOKEN")
	if !ok {
		return nil, ErrNoOpenAIToken
	}

	if len(*diff) > 10_000 {
		return nil, ErrDiffTooLarge
	}

	client := openai.NewClient(token)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: systemPrompt,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: *diff,
				},
			},
		},
	)

	if err != nil {
		log.Fatal("oops")
	}

	fmt.Println(resp.Choices[0].Message.Content)

	return nil, nil
}
