# OpenAI Embeddings Package

This package provides a client for generating text embeddings using the OpenAI API.

## Setup

1. **Set your API key** in `.env`:
   ```env
   OPENAI_API_KEY=sk-your_openai_api_key_here
   ```

2. **Install the dependency** (if not already in go.mod):
   ```bash
   go get github.com/sashabaranov/go-openai
   ```

## Usage

### Basic Example - Single Text

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "your-module/internal/embeddings"
)

func main() {
    // Create embeddings client
    client, err := embeddings.NewClient()
    if err != nil {
        log.Fatal(err)
    }
    
    // Generate embedding for a single text
    ctx := context.Background()
    text := "Hello, this is a sample text to embed"
    
    embedding, err := client.GenerateEmbedding(ctx, text)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Generated embedding with %d dimensions\n", len(embedding))
}
```

### Batch Processing - Multiple Texts

```go
// Generate embeddings for multiple texts efficiently
texts := []string{
    "First document to embed",
    "Second document to embed",
    "Third document to embed",
}

results, err := client.GenerateEmbeddingsBatch(ctx, texts)
if err != nil {
    log.Fatal(err)
}

for i, result := range results {
    if result.Error != nil {
        fmt.Printf("Error embedding text %d: %v\n", i, result.Error)
        continue
    }
    fmt.Printf("Text %d: %d dimensions\n", i, len(result.Embedding))
}
```

### Using with AsyncGlobalState

To integrate with the worker's shared state, add the client to `AsyncGlobalState`:

**1. Update `internal/state/async_global_state.go`:**

```go
type AsyncGlobalState struct {
    // ... existing fields ...
    EmbeddingsClient *embeddings.Client
}

func InitializeAsyncGlobalState(ctx context.Context) (*AsyncGlobalState, error) {
    // ... existing initialization ...
    
    // Initialize embeddings client
    embeddingsClient, err := embeddings.NewClient()
    if err != nil {
        log.Printf("⚠️  Failed to initialize embeddings client: %v", err)
        // Optionally continue without embeddings or return error
    }
    
    return &AsyncGlobalState{
        // ... existing fields ...
        EmbeddingsClient: embeddingsClient,
    }, nil
}
```

**2. Use in worker functions:**

```go
func MyFunction(ctx context.Context, ags *state.AsyncGlobalState, input MyInput) (MyOutput, error) {
    // Generate embedding for input text
    embedding, err := ags.EmbeddingsClient.GenerateEmbedding(ctx, input.Text)
    if err != nil {
        return MyOutput{}, fmt.Errorf("failed to generate embedding: %w", err)
    }
    
    // Use the embedding (e.g., store in database, compare similarity, etc.)
    // ...
    
    return MyOutput{Embedding: embedding}, nil
}
```

## Features

- **Automatic retries**: Built-in retry logic with exponential backoff (3 attempts)
- **Batch processing**: Efficiently process multiple texts in a single API call
- **Error handling**: Comprehensive error messages and validation
- **Cost estimation**: Helper functions to estimate token usage and costs
- **Model selection**: Uses `text-embedding-3-small` (1536 dimensions) by default

## Cost Estimation

```go
// Estimate tokens in text
text := "Your text here..."
tokens := embeddings.EstimateTokens(text)

// Estimate cost
cost := embeddings.EstimateCost(tokens)
fmt.Printf("Estimated cost: $%.4f\n", cost)
```

## Model Information

**Current Model**: `text-embedding-3-small`
- **Dimensions**: 1536
- **Cost**: $0.020 per 1M tokens
- **Performance**: Good balance of quality and cost

To use a different model, modify the `model` field in the `Client` struct initialization.

## Batch Size Limits

- Maximum batch size: **100 texts** (configurable in code)
- OpenAI API limit: ~2048 texts per request
- For larger datasets, split into multiple batches

## Common Use Cases

1. **Semantic search**: Store embeddings in a vector database (pgvector, Pinecone, etc.)
2. **Similarity comparison**: Calculate cosine similarity between embeddings
3. **Clustering**: Group similar texts based on embedding vectors
4. **Recommendation systems**: Find similar content based on embeddings
5. **Classification**: Use embeddings as features for ML models

## Error Handling

The client handles common errors:
- Missing API key
- Empty text input
- API failures (with retries)
- Invalid response format

Always check for errors in production:

```go
embedding, err := client.GenerateEmbedding(ctx, text)
if err != nil {
    // Handle error appropriately
    log.Printf("Embedding generation failed: %v", err)
    return err
}
```

## Resources

- [OpenAI Embeddings Guide](https://platform.openai.com/docs/guides/embeddings)
- [go-openai Library](https://github.com/sashabaranov/go-openai)
- [Text Embeddings Pricing](https://openai.com/pricing)

