package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"smart-chat/internal/dto"
)

type (
	DeepAiProviderMock struct {
		mock.Mock
	}
)

func (d *DeepAiProviderMock) CallIA(ctx context.Context, text string) (*dto.Answer, error) {
	args := d.Called(ctx, text)

	answer := &dto.Answer{}
	var err error

	if args.Get(0) != nil {
		answer = args.Get(0).(*dto.Answer)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return answer, err
}
