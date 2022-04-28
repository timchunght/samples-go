package helloworld

import (
	"context"
	"fmt"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"

	// TODO(cretz): Remove when tagged
	_ "go.temporal.io/sdk/contrib/tools/workflowcheck/determinism"
)

// Workflow is a Hello World workflow definition.
func Workflow(ctx workflow.Context, name string) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 100 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("HelloWorld workflow started", "name", name)

	var result string
	for i := 1; i < 5; i++ {
		err := workflow.ExecuteActivity(ctx, Activity, name, i).Get(ctx, &result)
		if err != nil {
			logger.Error("Activity failed.", "Error", err)
			return "", err
		}
	}

	logger.Info("HelloWorld workflow completed.", "result", result)

	return result, nil
}

func Activity(ctx context.Context, name string, i int) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Activity", "name", name)
	logger.Info(fmt.Sprintf("About to Sleeping for 10 sec at index: %v", i))
	time.Sleep(10 * time.Second)
	logger.Info(fmt.Sprintf("Sleep ended at index: %v", i))
	return "Hello " + name + "!", nil
}
