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

func AvatarsgetbrowserHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		codeVal, ok := args["code"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: code"), nil
		}
		code, ok := codeVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: code"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["width"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("width=%v", val))
		}
		if val, ok := args["height"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("height=%v", val))
		}
		if val, ok := args["quality"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("quality=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/avatars/browsers/%s%s", cfg.BaseURL, code, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Fallback to single auth parameter
		if cfg.APIKey != "" {
			req.Header.Set("X-Appwrite-Key", cfg.APIKey)
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
		var result map[string]interface{}
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

func CreateAvatarsgetbrowserTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_avatars_browsers_code",
		mcp.WithDescription("Get Browser Icon"),
		mcp.WithString("code", mcp.Required(), mcp.Description("Browser Code.")),
		mcp.WithNumber("width", mcp.Description("Image width. Pass an integer between 0 to 2000. Defaults to 100.")),
		mcp.WithNumber("height", mcp.Description("Image height. Pass an integer between 0 to 2000. Defaults to 100.")),
		mcp.WithNumber("quality", mcp.Description("Image quality. Pass an integer between 0 to 100. Defaults to 100.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    AvatarsgetbrowserHandler(cfg),
	}
}
