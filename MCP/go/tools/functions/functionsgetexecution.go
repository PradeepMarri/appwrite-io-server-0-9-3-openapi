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

func FunctionsgetexecutionHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		functionIdVal, ok := args["functionId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: functionId"), nil
		}
		functionId, ok := functionIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: functionId"), nil
		}
		executionIdVal, ok := args["executionId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: executionId"), nil
		}
		executionId, ok := executionIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: executionId"), nil
		}
		url := fmt.Sprintf("%s/functions/%s/executions/%s", cfg.BaseURL, functionId, executionId)
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
		var result models.Execution
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

func CreateFunctionsgetexecutionTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_functions_functionId_executions_executionId",
		mcp.WithDescription("Get Execution"),
		mcp.WithString("functionId", mcp.Required(), mcp.Description("Function unique ID.")),
		mcp.WithString("executionId", mcp.Required(), mcp.Description("Execution unique ID.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    FunctionsgetexecutionHandler(cfg),
	}
}
