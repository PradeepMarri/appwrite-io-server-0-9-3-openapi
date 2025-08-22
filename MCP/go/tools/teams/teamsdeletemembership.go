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

func TeamsdeletemembershipHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		teamIdVal, ok := args["teamId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: teamId"), nil
		}
		teamId, ok := teamIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: teamId"), nil
		}
		membershipIdVal, ok := args["membershipId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: membershipId"), nil
		}
		membershipId, ok := membershipIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: membershipId"), nil
		}
		url := fmt.Sprintf("%s/teams/%s/memberships/%s", cfg.BaseURL, teamId, membershipId)
		req, err := http.NewRequest("DELETE", url, nil)
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

func CreateTeamsdeletemembershipTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_teams_teamId_memberships_membershipId",
		mcp.WithDescription("Delete Team Membership"),
		mcp.WithString("teamId", mcp.Required(), mcp.Description("Team unique ID.")),
		mcp.WithString("membershipId", mcp.Required(), mcp.Description("Membership ID.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    TeamsdeletemembershipHandler(cfg),
	}
}
