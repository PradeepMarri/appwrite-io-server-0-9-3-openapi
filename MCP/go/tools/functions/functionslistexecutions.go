package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/appwrite/mcp-server/config"
	"github.com/appwrite/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func FunctionslistexecutionsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		queryParams := make([]string, 0)
		if val, ok := args["search"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("search=%v", val))
		}
		if val, ok := args["limit"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit=%v", val))
		}
		if val, ok := args["offset"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("offset=%v", val))
		}
		if val, ok := args["orderType"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("orderType=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/functions/%s/executions%s", cfg.BaseURL, functionId, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Fallback to single auth parameter
		if cfg.APIKey != "" {
			req.Header.Set("X-Appwrite-Project", cfg.APIKey)
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
		var result models.ExecutionList
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

func CreateFunctionslistexecutionsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_functions_functionId_executions",
		mcp.WithDescription("List Executions"),
		mcp.WithString("functionId", mcp.Required(), mcp.Description("Function unique ID.")),
		mcp.WithString("search", mcp.Description("Search term to filter your list results. Max length: 256 chars.")),
		mcp.WithNumber("limit", mcp.Description("Results limit value. By default will return maximum 25 results. Maximum of 100 results allowed per request.")),
		mcp.WithNumber("offset", mcp.Description("Results offset. The default value is 0. Use this param to manage pagination.")),
		mcp.WithString("orderType", mcp.Description("Order result by ASC or DESC order.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    FunctionslistexecutionsHandler(cfg),
	}
}
