// Package processbatch demonstrates calling a job from a worker function.
//
// This pattern is useful when:
// - An external workflow triggers a multi-step operation
// - You need to orchestrate complex logic with proper error handling
// - You want to track job execution metrics
package processbatch

import (
	"fmt"

	sdk "github.com/dibbla-agents/sdk-go"
	"github.com/dibbla-agents/go-worker-starter-template/internal/jobs"
)

// Input for the worker function
type ProcessBatchInput struct {
	BatchName string `json:"batch_name"`
	ItemCount int    `json:"item_count"`
}

// Output from the worker function
type ProcessBatchOutput struct {
	Success        bool    `json:"success"`
	ItemsProcessed int     `json:"items_processed"`
	DurationMs     float64 `json:"duration_ms"`
	Error          string  `json:"error,omitempty"`
}

// Register registers the process_batch function with the SDK server.
func Register(server *sdk.Server) {
	fn := sdk.NewSimpleFunction[ProcessBatchInput, ProcessBatchOutput](
		"process_batch",
		"1.0.0",
		"Process a batch of items using a job",
	).WithHandler(func(input ProcessBatchInput) (ProcessBatchOutput, error) {
		// Validate input
		if input.BatchName == "" {
			return ProcessBatchOutput{}, fmt.Errorf("batch_name is required")
		}
		if input.ItemCount <= 0 {
			return ProcessBatchOutput{}, fmt.Errorf("item_count must be positive")
		}

		// Create and execute the job
		job := jobs.NewSimpleJob(input.BatchName, input.ItemCount)
		result := job.Execute()

		// Convert job result to worker output
		output := ProcessBatchOutput{
			Success:        result.Success,
			ItemsProcessed: result.ItemsProcessed,
			DurationMs:     float64(result.Duration.Milliseconds()),
		}

		if result.Error != nil {
			output.Error = result.Error.Error()
			return output, nil // Return error in output, not as function error
		}

		return output, nil
	})

	server.RegisterFunction(fn)
}

