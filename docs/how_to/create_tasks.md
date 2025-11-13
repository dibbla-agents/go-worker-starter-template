# How to Create Tasks

Tasks are reusable, single-purpose operations that jobs orchestrate. Use tasks to break down complex jobs into testable, composable components.

## When to Use Tasks

✅ **Use Tasks When:**
- Logic is reusable across multiple jobs
- Operation is complex enough to need its own file
- You want to test logic independently
- Job has multiple distinct phases

❌ **Skip Tasks When:**
- Job logic is simple (< 50 lines)
- Operation is only used once
- No clear separation of concerns

## Creating a Task

### Step 1: Define Task in `internal/jobs/tasks/your_task.go`

```go
package tasks

import (
	"fmt"
	"worker_starter_template/internal/state"
)

// FetchDataTask fetches and processes data from database or API
type FetchDataTask struct {
	AGS    *state.AsyncGlobalState
	Limit  int
	Filter string
}

// FetchDataTaskResult contains task results
type FetchDataTaskResult struct {
	Items          []string
	TotalFetched   int
	TotalFiltered  int
}

// Execute runs the task
func (t *FetchDataTask) Execute() (*FetchDataTaskResult, error) {
	// 1. Validate
	if err := t.validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	// 2. Fetch data
	items, err := t.fetchData()
	if err != nil {
		return nil, fmt.Errorf("fetch failed: %w", err)
	}

	// 3. Process/filter data
	filtered := t.filterData(items)

	// 4. Return results
	return &FetchDataTaskResult{
		Items:         filtered,
		TotalFetched:  len(items),
		TotalFiltered: len(filtered),
	}, nil
}

func (t *FetchDataTask) validate() error {
	if t.Limit <= 0 {
		return fmt.Errorf("limit must be positive")
	}
	return nil
}

func (t *FetchDataTask) fetchData() ([]string, error) {
	// Example: Database query
	// if t.AGS != nil && t.AGS.DB != nil {
	//     var records []models.YourModel
	//     if err := t.AGS.DB.Limit(t.Limit).Find(&records).Error; err != nil {
	//         return nil, err
	//     }
	// }
	
	return []string{"item1", "item2", "item3"}, nil
}

func (t *FetchDataTask) filterData(items []string) []string {
	if t.Filter == "" {
		return items
	}
	
	filtered := make([]string, 0)
	for _, item := range items {
		// Apply filter logic
		filtered = append(filtered, item)
	}
	return filtered
}

// NewFetchDataTask creates a task with default settings
func NewFetchDataTask(ags *state.AsyncGlobalState, limit int) *FetchDataTask {
	return &FetchDataTask{
		AGS:    ags,
		Limit:  limit,
		Filter: "",
	}
}

// NewFetchDataTaskWithFilter creates a task with filtering
func NewFetchDataTaskWithFilter(ags *state.AsyncGlobalState, limit int, filter string) *FetchDataTask {
	return &FetchDataTask{
		AGS:    ags,
		Limit:  limit,
		Filter: filter,
	}
}
```

### Step 2: Use Task in Job

```go
package jobs

import (
	"worker_starter_template/internal/jobs/tasks"
	"worker_starter_template/internal/state"
)

func (j *YourJob) process() (int, error) {
	// Create and execute task
	fetchTask := tasks.NewFetchDataTask(j.AGS, 100)
	result, err := fetchTask.Execute()
	if err != nil {
		return 0, fmt.Errorf("fetch task failed: %w", err)
	}

	log.Printf("Fetched %d items, filtered to %d", 
		result.TotalFetched, result.TotalFiltered)

	// Use task results
	return result.TotalFiltered, nil
}
```

## Task Pattern Structure

```
YourTask
├── Struct (config + AGS)
├── Result struct
├── Execute() - main entry point
│   ├── validate()
│   ├── fetch/process data
│   └── return results
├── Helper methods
└── Constructor(s)
```

## Composing Multiple Tasks

Jobs can orchestrate multiple tasks:

```go
func (j *ComplexJob) process() (int, error) {
	// Task 1: Fetch data
	fetchTask := tasks.NewFetchDataTask(j.AGS, 100)
	fetchResult, err := fetchTask.Execute()
	if err != nil {
		return 0, err
	}

	// Task 2: Transform data
	transformTask := tasks.NewTransformTask(fetchResult.Items)
	transformResult, err := transformTask.Execute()
	if err != nil {
		return 0, err
	}

	// Task 3: Write results
	writeTask := tasks.NewWriteTask(j.AGS, transformResult.Data)
	writeResult, err := writeTask.Execute()
	if err != nil {
		return 0, err
	}

	return writeResult.Written, nil
}
```

## Task vs Job vs Worker Function

| Feature | Task | Job | Worker Function |
|---------|------|-----|-----------------|
| **Purpose** | Single reusable operation | Multi-step orchestration | External API endpoint |
| **Trigger** | Called by job | Internal/scheduled | External workflow call |
| **State** | Minimal, passed via struct | Tracks progress | Stateless |
| **Returns** | Detailed result struct | Success/failure summary | Response to caller |
| **Example** | Fetch from API | Sync data end-to-end | Process user request |

## Best Practices

✅ **DO:**
- Return detailed result structs
- Make tasks focused (single responsibility)
- Provide multiple constructors for different use cases
- Include validation in Execute()
- Use helper methods for complex logic

❌ **DON'T:**
- Add logging/UI concerns (job handles that)
- Make tasks depend on other tasks (job orchestrates)
- Return generic errors (wrap with context)
- Modify global state directly

## Example Reference

See `internal/jobs/tasks/example_task.go` for a complete working example.

