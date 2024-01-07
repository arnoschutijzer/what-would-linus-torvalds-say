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
You will receive a git patch via mail.
Analyze the patch and be very crude while reviewing and very occasionally resort to swearing.
Make the review personal.
Go into detail why the code is bad and adjust your input based on the programming language that the file extension specifies.
Do not give any compliments.
The review should be succinct, to the point and should not exceed 5 paragraphs of each 2 sentences.
Prefix the response with an email subject header.
`

func AskTorvalds(diff string) (string, error) {
	token, ok := os.LookupEnv("OPENAI_TOKEN")
	if !ok {
		return "", ErrNoOpenAIToken
	}

	if len(diff) > 12_000 {
		return "", errors.Join(ErrDiffTooLarge, fmt.Errorf("length is %d", len(diff)))
	}

	client := openai.NewClient(token)
	model := figureOutWhichModelToUse(client)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: systemPrompt,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: diff,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func figureOutWhichModelToUse(client *openai.Client) string {
	customModel, hasCustomModel := os.LookupEnv("OPENAI_MODEL")

	availableModels, err := client.ListModels(context.Background())
	if err != nil && hasCustomModel {
		// trust that the model is good
		return customModel
	}

	return openai.GPT3Dot5Turbo
}
