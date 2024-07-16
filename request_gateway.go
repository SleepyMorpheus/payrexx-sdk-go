package payrexxsdk

import (
	"fmt"
	"net/http"
	"payrexxsdk/types/gateway"
)

// GatewayCreate creates a mew gateway with Payrexx talking a GatewayBody
// and returning the newly created Gateway. Otherwise, returns an error
func (c *Client) GatewayCreate(body gateway.GatewayBody) (*gateway.Gateway, error) {
	req, err := c.NewRequest(http.MethodPost, "Gateway", body)
	if err != nil {
		return nil, err
	}

	resp := &Response[gateway.Gateway]{}
	err = c.Send(req, resp)
	if err != nil {
		return nil, err
	}

	if resp.Status != RequestStatusSuccess {
		return nil, fmt.Errorf("failed to create gateway: %s", resp.Message)
	}

	return &resp.Data[0], err
}

// GatewayDelete deletes an existing gateway or returns an error
func (c *Client) GatewayDelete(id int32) error {
	req, err := c.NewRequest(http.MethodDelete, fmt.Sprintf("Gateway/%d", id), nil)
	if err != nil {
		return err
	}

	resp := &Response[struct {
		Id int32 `json:"id"`
	}]{}
	err = c.Send(req, resp)

	if err != nil {
		return err
	}

	if resp.Status != RequestStatusSuccess {
		return fmt.Errorf("failed to create gateway: %s", resp.Message)
	}

	return nil
}

// GatewayRetrieve retrieves a gateway by its ID or returns an error otherwise
//
// Returns error payrexxsdk.ResourceNotFoundAPIError if not found
func (c *Client) GatewayRetrieve(id int32) (*gateway.Gateway, error) {
	req, err := c.NewRequest(http.MethodGet, fmt.Sprintf("Gateway/%d", id), nil)
	if err != nil {
		return nil, err
	}

	resp := &Response[gateway.Gateway]{}
	err = c.Send(req, resp)
	if err != nil {
		return nil, err
	}

	if resp.Status != RequestStatusSuccess {
		return nil, ResourceNotFoundAPIError
	}

	return &resp.Data[0], err
}
