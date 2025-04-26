package option

import (
	"github.com/poligono-xyz/alan"
)

func WithTemperature(temperature float32) alan.ChatOption {
	return func(config *alan.ChatConfig) {
		config.Temperature = temperature
	}
}

func WithTopK(topK float32) alan.ChatOption {
	return func(config *alan.ChatConfig) {
		config.TopK = topK
	}
}
func WithTopP(topP float32) alan.ChatOption {
	return func(config *alan.ChatConfig) {
		config.TopP = topP
	}
}
func WithSystemInstruction(systemInstruction string) alan.ChatOption {
	return func(config *alan.ChatConfig) {
		config.SystemInstruction = systemInstruction
	}
}
func WithCandidateCount(candidateCount int) alan.ChatOption {
	return func(config *alan.ChatConfig) {
		config.CandidateCount = candidateCount
	}
}
