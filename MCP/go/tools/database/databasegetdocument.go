package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/appwrite/mcp-server/config"
	"github.com/appwrite/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func DatabasegetdocumentHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		collectionIdVal, ok := args["collectionId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: collectionId"), nil
		}
		collectionId, ok := collectionIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: collectionId"), nil
		}
		documentIdVal, ok := args["documentId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: documentId"), nil
		}
		documentId, ok := documentIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: documentId"), nil
		}
		url := fmt.Sprintf("%s/database/collections/%s/documents/%s", cfg.BaseURL, collectionId, documentId)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Fallback to single auth parameter
		if cfg.APIKey != "" {
			req.Header.Set("X-Appwrite-JWT", cfg.APIKey)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.Document
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateDatabasegetdocumentTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_database_collections_collectionId_documents_documentId",
		mcp.WithDescription("Get Document"),
		mcp.WithString("collectionId", mcp.Required(), mcp.Description("Collection unique ID. You can create a new collection with validation rules using the Database service [server integration](/docs/server/database#createCollection).")),
		mcp.WithString("documentId", mcp.Required(), mcp.Description("Document unique ID.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DatabasegetdocumentHandler(cfg),
	}
}
