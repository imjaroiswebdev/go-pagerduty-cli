package pagerduty

import (
	"context"
	"fmt"

	"github.com/google/go-querystring/query"
)

type WorkflowIntegrationConnectionHealth struct {
	IsHealthy     bool   `json:"is_healthy"`
	HealthMessage string `json:"health_message"`
	LastCheckedAt string `json:"last_checked_at"`
}

type WorkflowIntegrationConnection struct {
	ID              string                               `json:"id,omitempty"`
	Type            string                               `json:"type,omitempty"`
	IntegrationID   string                               `json:"integration_id,omitempty"`
	Name            string                               `json:"name"`
	ServiceURL      string                               `json:"service_url,omitempty"`
	ExternalID      string                               `json:"external_id,omitempty"`
	ExternalIDLabel string                               `json:"external_id_label,omitempty"`
	Scopes          []string                             `json:"scopes,omitempty"`
	IsDefault       *bool                                `json:"is_default,omitempty"`
	Health          *WorkflowIntegrationConnectionHealth `json:"health,omitempty"`
	Configuration   map[string]interface{}               `json:"configuration,omitempty"`
	Secrets         map[string]interface{}               `json:"secrets,omitempty"`
	Teams           []WorkflowIntegrationConnectionTeam  `json:"teams,omitempty"`
	CreatedAt       string                               `json:"created_at,omitempty"`
	CreatedBy       *APIObject                           `json:"created_by,omitempty"`
}

type WorkflowIntegrationConnectionTeam struct {
	TeamID string `json:"team_id"`
	Type   string `json:"type"`
}

type ListWorkflowIntegrationConnectionsOptions struct {
	Limit  uint   `url:"limit,omitempty"`
	Cursor string `url:"cursor,omitempty"`
	Name   string `url:"name,omitempty"`
}

type ListWorkflowIntegrationConnectionsResponse struct {
	Limit       uint                            `json:"limit"`
	NextCursor  string                          `json:"next_cursor"`
	Connections []WorkflowIntegrationConnection `json:"connections"`
}

func (c *Client) ListWorkflowIntegrationConnections(ctx context.Context, integrationID string, o ListWorkflowIntegrationConnectionsOptions) (*ListWorkflowIntegrationConnectionsResponse, error) {
	v, err := query.Values(o)
	if err != nil {
		return nil, err
	}

	resp, err := c.get(ctx, fmt.Sprintf("/workflows/integrations/%s/connections?%s", integrationID, v.Encode()), nil)
	if err != nil {
		return nil, err
	}

	var result ListWorkflowIntegrationConnectionsResponse
	if err := c.decodeJSON(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateWorkflowIntegrationConnection(ctx context.Context, integrationID string, connection *WorkflowIntegrationConnection) (*WorkflowIntegrationConnection, error) {
	resp, err := c.post(ctx, fmt.Sprintf("/workflows/integrations/%s/connections", integrationID), connection, nil)
	if err != nil {
		return nil, err
	}

	var result WorkflowIntegrationConnection
	if err := c.decodeJSON(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetWorkflowIntegrationConnection(ctx context.Context, integrationID, connectionID string) (*WorkflowIntegrationConnection, error) {
	resp, err := c.get(ctx, fmt.Sprintf("/workflows/integrations/%s/connections/%s", integrationID, connectionID), nil)
	if err != nil {
		return nil, err
	}

	var result WorkflowIntegrationConnection
	if err := c.decodeJSON(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateWorkflowIntegrationConnection(ctx context.Context, integrationID string, connection *WorkflowIntegrationConnection) (*WorkflowIntegrationConnection, error) {
	resp, err := c.patch(ctx, fmt.Sprintf("/workflows/integrations/%s/connections/%s", integrationID, connection.ID), connection, nil)
	if err != nil {
		return nil, err
	}

	var result WorkflowIntegrationConnection
	if err := c.decodeJSON(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteWorkflowIntegrationConnection(ctx context.Context, integrationID, connectionID string) error {
	_, err := c.delete(ctx, fmt.Sprintf("/workflows/integrations/%s/connections/%s", integrationID, connectionID))
	return err
}
