package torvalds

import (
	"errors"
)

var ErrDiffTooLarge = errors.New("diff too large")

func AskTorvalds(diff *string) (*string, error) {
	if len(*diff) > 10_000 {
		return nil, ErrDiffTooLarge
	}

	return nil, nil
}
