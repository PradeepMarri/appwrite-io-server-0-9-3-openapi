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

func StoragegetfilepreviewHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		fileIdVal, ok := args["fileId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: fileId"), nil
		}
		fileId, ok := fileIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: fileId"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["width"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("width=%v", val))
		}
		if val, ok := args["height"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("height=%v", val))
		}
		if val, ok := args["gravity"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("gravity=%v", val))
		}
		if val, ok := args["quality"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("quality=%v", val))
		}
		if val, ok := args["borderWidth"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("borderWidth=%v", val))
		}
		if val, ok := args["borderColor"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("borderColor=%v", val))
		}
		if val, ok := args["borderRadius"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("borderRadius=%v", val))
		}
		if val, ok := args["opacity"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("opacity=%v", val))
		}
		if val, ok := args["rotation"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("rotation=%v", val))
		}
		if val, ok := args["background"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("background=%v", val))
		}
		if val, ok := args["output"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("output=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/storage/files/%s/preview%s", cfg.BaseURL, fileId, queryString)
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

func CreateStoragegetfilepreviewTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_storage_files_fileId_preview",
		mcp.WithDescription("Get File Preview"),
		mcp.WithString("fileId", mcp.Required(), mcp.Description("File unique ID")),
		mcp.WithNumber("width", mcp.Description("Resize preview image width, Pass an integer between 0 to 4000.")),
		mcp.WithNumber("height", mcp.Description("Resize preview image height, Pass an integer between 0 to 4000.")),
		mcp.WithString("gravity", mcp.Description("Image crop gravity. Can be one of center,top-left,top,top-right,left,right,bottom-left,bottom,bottom-right")),
		mcp.WithNumber("quality", mcp.Description("Preview image quality. Pass an integer between 0 to 100. Defaults to 100.")),
		mcp.WithNumber("borderWidth", mcp.Description("Preview image border in pixels. Pass an integer between 0 to 100. Defaults to 0.")),
		mcp.WithString("borderColor", mcp.Description("Preview image border color. Use a valid HEX color, no # is needed for prefix.")),
		mcp.WithNumber("borderRadius", mcp.Description("Preview image border radius in pixels. Pass an integer between 0 to 4000.")),
		mcp.WithString("opacity", mcp.Description("Preview image opacity. Only works with images having an alpha channel (like png). Pass a number between 0 to 1.")),
		mcp.WithNumber("rotation", mcp.Description("Preview image rotation in degrees. Pass an integer between 0 and 360.")),
		mcp.WithString("background", mcp.Description("Preview image background color. Only works with transparent images (png). Use a valid HEX color, no # is needed for prefix.")),
		mcp.WithString("output", mcp.Description("Output format type (jpeg, jpg, png, gif and webp).")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    StoragegetfilepreviewHandler(cfg),
	}
}
