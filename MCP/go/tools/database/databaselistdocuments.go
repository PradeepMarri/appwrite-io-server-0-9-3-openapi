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

func DatabaselistdocumentsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		queryParams := make([]string, 0)
		if val, ok := args["filters"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filters=%v", val))
		}
		if val, ok := args["limit"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit=%v", val))
		}
		if val, ok := args["offset"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("offset=%v", val))
		}
		if val, ok := args["orderField"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("orderField=%v", val))
		}
		if val, ok := args["orderType"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("orderType=%v", val))
		}
		if val, ok := args["orderCast"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("orderCast=%v", val))
		}
		if val, ok := args["search"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("search=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/database/collections/%s/documents%s", cfg.BaseURL, collectionId, queryString)
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
		var result models.DocumentList
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

func CreateDatabaselistdocumentsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_database_collections_collectionId_documents",
		mcp.WithDescription("List Documents"),
		mcp.WithString("collectionId", mcp.Required(), mcp.Description("Collection unique ID. You can create a new collection with validation rules using the Database service [server integration](/docs/server/database#createCollection).")),
		mcp.WithArray("filters", mcp.Description("Array of filter strings. Each filter is constructed from a key name, comparison operator (=, !=, >, <, <=, >=) and a value. You can also use a dot (.) separator in attribute names to filter by child document attributes. Examples: 'name=John Doe' or 'category.$id>=5bed2d152c362'.")),
		mcp.WithNumber("limit", mcp.Description("Maximum number of documents to return in response.  Use this value to manage pagination. By default will return maximum 25 results. Maximum of 100 results allowed per request.")),
		mcp.WithNumber("offset", mcp.Description("Offset value. The default value is 0. Use this param to manage pagination.")),
		mcp.WithString("orderField", mcp.Description("Document field that results will be sorted by.")),
		mcp.WithString("orderType", mcp.Description("Order direction. Possible values are DESC for descending order, or ASC for ascending order.")),
		mcp.WithString("orderCast", mcp.Description("Order field type casting. Possible values are int, string, date, time or datetime. The database will attempt to cast the order field to the value you pass here. The default value is a string.")),
		mcp.WithString("search", mcp.Description("Search query. Enter any free text search. The database will try to find a match against all document attributes and children. Max length: 256 chars.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DatabaselistdocumentsHandler(cfg),
	}
}
