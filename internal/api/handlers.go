// internal/api/handlers.go
package api

import (
	"context"
	"net/http"

	"BAPE/internal/workflow"

	"github.com/gin-gonic/gin"
	"github.com/pborman/uuid"
	"go.temporal.io/sdk/client"
)

func StartWorkflowHandler(c *gin.Context) {
	temporalClient, err := client.Dial(client.Options{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create Temporal client"})
		return
	}
	defer temporalClient.Close()

	options := client.StartWorkflowOptions{
		ID:        "bape-workflow_" + uuid.New(),
		TaskQueue: workflow.BapeTaskQueue,
	}

	we, err := temporalClient.ExecuteWorkflow(context.Background(), options, workflow.BapeWorkflow, "BAPE")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to start workflow"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"workflow_id": we.GetID(),
		"run_id":      we.GetRunID(),
	})
}
