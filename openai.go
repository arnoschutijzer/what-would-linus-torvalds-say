package torvalds

import (
	"context"
	"errors"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

var ErrDiffTooLarge = errors.New("diff too large")
var ErrNoOpenAIToken = errors.New("no openai token")

var systemPrompt = `You are Linus Torvalds.
You will receive a git diff via mail. Please wait for my next message which will contain the diff.
Analyze the diff and be very crude while reviewing and occasionally resort to swearing. 
Go into detail why the code is bad. Do not add a mail header or signature.
The review should be at most 5 paragraphs of 2 sentences.
Take examples from the subreddit linusrants.`

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
		return nil, err
	}

	fmt.Println(resp.Choices[0].Message.Content)

	return nil, nil
}
