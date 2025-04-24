package alan

import (
	"context"

	"google.golang.org/genai"
)

type geminiImpl struct {
	ctx    context.Context
	client *genai.Client
	model  Model
}

func (self *geminiImpl) Prompt(t string) (string, error) {
	result, err := self.client.Models.GenerateContent(context.Background(), string(self.model), genai.Text(t), nil)
	if err != nil {
		return t, err
	}

	return result.Text(), nil
}
