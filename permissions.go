package slack

import "context"

const PermissionCanAtChannel = "CAN_AT_CHANNEL"

type GetPermissionsParameters struct {
	Token       string   `json:"token"`
	UserID      string   `json:"user_id"`
	Permissions []string `json:"permissions"`
	ChannelID   string   `json:"channel_id,omitempty"`
}

type GetPermissionsResponse struct {
	SlackResponse
	Allowed []string `json:"allowed"`
}

func (api *Client) GetPermissionsContext(ctx context.Context, teamID string, params GetPermissionsParameters) (*GetPermissionsResponse, error) {
	if api == nil {
		return nil, ErrNilClient
	}
	params.Token = api.token
	var resp GetPermissionsResponse
	err := api.postEdgeAPI(ctx, teamID, "permissions.info", params, &resp)
	if err != nil {
		return nil, err
	}
	err = resp.Err()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
