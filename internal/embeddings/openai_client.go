package embeddings

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sashabaranov/go-openai"
)

// Client wraps the OpenAI API client for embedding generation
type Client struct {
	client *openai.Client
	model  openai.EmbeddingModel
}

// EmbeddingResult represents the result of an embedding operation
type EmbeddingResult struct {
	Text      string
	Embedding []float32
	Error     error
}

// NewClient creates a new OpenAI embeddings client
func NewClient() (*Client, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("OPENAI_API_KEY environment variable is not set")
	}

	client := openai.NewClient(apiKey)

	return &Client{
		client: client,
		model:  openai.SmallEmbedding3, // text-embedding-3-small (1536 dimensions)
	}, nil
}

// GenerateEmbedding generates an embedding for a single text string
func (c *Client) GenerateEmbedding(ctx context.Context, text string) ([]float32, error) {
	if text == "" {
		return nil, fmt.Errorf("text cannot be empty")
	}

	// Create embedding request with retry logic
	var resp openai.EmbeddingResponse
	var err error

	maxRetries := 3
	for attempt := 1; attempt <= maxRetries; attempt++ {
		resp, err = c.client.CreateEmbeddings(ctx, openai.EmbeddingRequest{
			Input: []string{text},
			Model: c.model,
		})

		if err == nil {
			break
		}

		if attempt < maxRetries {
			waitTime := time.Duration(attempt) * time.Second
			log.Printf("⚠️  Embedding request failed (attempt %d/%d), retrying in %v: %v",
				attempt, maxRetries, waitTime, err)
			time.Sleep(waitTime)
		}
	}

	if err != nil {
		return nil, fmt.Errorf("failed to create embedding after %d attempts: %w", maxRetries, err)
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no embedding data returned from API")
	}

	return resp.Data[0].Embedding, nil
}

// GenerateEmbeddingsBatch generates embeddings for multiple texts in a single API call
// This is more efficient than calling GenerateEmbedding multiple times
func (c *Client) GenerateEmbeddingsBatch(ctx context.Context, texts []string) ([]EmbeddingResult, error) {
	if len(texts) == 0 {
		return nil, fmt.Errorf("texts cannot be empty")
	}

	// OpenAI API has a limit on batch size, typically around 2048 texts
	// We'll use a conservative batch size of 100
	maxBatchSize := 100
	if len(texts) > maxBatchSize {
		return nil, fmt.Errorf("batch size %d exceeds maximum of %d", len(texts), maxBatchSize)
	}

	// Filter out empty texts and keep track of original indices
	validTexts := make([]string, 0, len(texts))
	textIndices := make([]int, 0, len(texts))
	for i, text := range texts {
		if text != "" {
			validTexts = append(validTexts, text)
			textIndices = append(textIndices, i)
		}
	}

	if len(validTexts) == 0 {
		return nil, fmt.Errorf("no valid texts to embed")
	}

	// Create embedding request with retry logic
	var resp openai.EmbeddingResponse
	var err error

	maxRetries := 3
	for attempt := 1; attempt <= maxRetries; attempt++ {
		resp, err = c.client.CreateEmbeddings(ctx, openai.EmbeddingRequest{
			Input: validTexts,
			Model: c.model,
		})

		if err == nil {
			break
		}

		if attempt < maxRetries {
			waitTime := time.Duration(attempt) * time.Second
			log.Printf("⚠️  Batch embedding request failed (attempt %d/%d), retrying in %v: %v",
				attempt, maxRetries, waitTime, err)
			time.Sleep(waitTime)
		}
	}

	if err != nil {
		return nil, fmt.Errorf("failed to create batch embeddings after %d attempts: %w", maxRetries, err)
	}

	if len(resp.Data) != len(validTexts) {
		return nil, fmt.Errorf("expected %d embeddings, got %d", len(validTexts), len(resp.Data))
	}

	// Build results array maintaining original order
	results := make([]EmbeddingResult, len(texts))
	for i, validIdx := range textIndices {
		results[validIdx] = EmbeddingResult{
			Text:      texts[validIdx],
			Embedding: resp.Data[i].Embedding,
			Error:     nil,
		}
	}

	// Mark empty texts with errors
	for i, text := range texts {
		if text == "" {
			results[i] = EmbeddingResult{
				Text:      text,
				Embedding: nil,
				Error:     fmt.Errorf("text is empty"),
			}
		}
	}

	return results, nil
}

// EstimateCost estimates the cost of embedding a given number of tokens
// Based on OpenAI's pricing: text-embedding-3-small costs $0.020 per 1M tokens
func EstimateCost(numTokens int) float64 {
	costPer1MTokens := 0.020 // USD
	return (float64(numTokens) / 1000000.0) * costPer1MTokens
}

// EstimateTokens roughly estimates the number of tokens in a text
// This is a simple approximation: ~4 characters per token
func EstimateTokens(text string) int {
	return len(text) / 4
}

