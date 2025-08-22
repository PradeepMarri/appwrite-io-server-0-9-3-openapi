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

func TeamsupdatemembershiprolesHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		url := fmt.Sprintf("%s/teams/%s/memberships/%s", cfg.BaseURL, teamId, membershipId)
		req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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
		var result models.Membership
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

func CreateTeamsupdatemembershiprolesTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_teams_teamId_memberships_membershipId",
		mcp.WithDescription("Update Membership Roles"),
		mcp.WithString("teamId", mcp.Required(), mcp.Description("Team unique ID.")),
		mcp.WithString("membershipId", mcp.Required(), mcp.Description("Membership ID.")),
		mcp.WithArray("roles", mcp.Required(), mcp.Description("Input parameter: Array of strings. Use this param to set the user roles in the team. A role can be any string. Learn more about [roles and permissions](/docs/permissions). Max length for each role is 32 chars.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    TeamsupdatemembershiprolesHandler(cfg),
	}
}
