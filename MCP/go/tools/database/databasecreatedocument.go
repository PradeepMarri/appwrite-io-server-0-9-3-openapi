package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/appwrite/mcp-server/config"
	"github.com/appwrite/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func DatabasecreatedocumentHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		// Create properly typed request body using the generated schema
		var requestBody map[string]interface{}
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/database/collections/%s/documents", cfg.BaseURL, collectionId)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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

func CreateDatabasecreatedocumentTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_database_collections_collectionId_documents",
		mcp.WithDescription("Create Document"),
		mcp.WithString("collectionId", mcp.Required(), mcp.Description("Collection unique ID. You can create a new collection with validation rules using the Database service [server integration](/docs/server/database#createCollection).")),
		mcp.WithString("parentPropertyType", mcp.Description("Input parameter: Parent document property connection type. You can set this value to **assign**, **append** or **prepend**, default value is assign. Use when you want your new document to be a child of a parent document.")),
		mcp.WithArray("read", mcp.Description("Input parameter: An array of strings with read permissions. By default only the current user is granted with read permissions. [learn more about permissions](/docs/permissions) and get a full list of available permissions.")),
		mcp.WithArray("write", mcp.Description("Input parameter: An array of strings with write permissions. By default only the current user is granted with write permissions. [learn more about permissions](/docs/permissions) and get a full list of available permissions.")),
		mcp.WithObject("data", mcp.Required(), mcp.Description("Input parameter: Document data as JSON object.")),
		mcp.WithString("parentDocument", mcp.Description("Input parameter: Parent document unique ID. Use when you want your new document to be a child of a parent document.")),
		mcp.WithString("parentProperty", mcp.Description("Input parameter: Parent document property name. Use when you want your new document to be a child of a parent document.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DatabasecreatedocumentHandler(cfg),
	}
}
