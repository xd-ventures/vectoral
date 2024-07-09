package openai

import (
	"context"
	"log"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func CalculateOpenAIAdaEmbedding(text string) ([]float32, error) {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	queryReq := openai.EmbeddingRequest{
		Input: []string{text},
		Model: openai.AdaEmbeddingV2,
	}

	queryResponse, err := client.CreateEmbeddings(context.Background(), queryReq)
	if err != nil {
		log.Fatal("Error creating query embedding:", err)
	}
	return queryResponse.Data[0].Embedding, err
}
