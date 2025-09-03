// internal/workflow/workflow.go
package workflow

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func BapeWorkflow(ctx workflow.Context, input string) (string, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	var result string
	err := workflow.ExecuteActivity(ctx, BapeActivity, input).Get(ctx, &result)
	if err != nil {
		return "", err
	}

	return "Workflow completed: " + result, nil
}
