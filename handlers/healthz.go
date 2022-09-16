package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	fs "github.com/tmccoy14/mlb/internal/firestore"
)

type HealthCheckResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func HealthCheck(c echo.Context) error {

	// Set resp to the go struct pointer
	resp := HealthCheckResponse{
		Message: "OK!",
		Status:  http.StatusOK,
	}

	// Get a Firestore client
	ctx := context.Background()
	client, err := fs.CreateClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	// Create Health Check Reponse dataset for collection
	dataset := map[string]interface{}{
		"message": resp.Message,
		"status":  resp.Status,
	}

	// Add dataset to Firestore collection
	if err := fs.AddCollection(client, ctx, dataset); err != nil {
		return err
	}

	// Read documentID from Firestore collection
	documentID, err := fs.ReadCollection(client, ctx)
	if err != nil {
		return err
	}

	// Update dataset in Firestore collection
	if err := fs.UpdateCollection(client, ctx, documentID); err != nil {
		return err
	}

	// Return health check message and status code
	return c.JSON(http.StatusOK, resp)
}
