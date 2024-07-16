package payrexxsdk

import (
	"fmt"
	"net/http"
	"payrexxsdk/types/paylink"
)

// PaylinkCreate creates a mew Paylink with Payrexx talking a PaylinkBody
// and returning the newly created Paylink. Otherwise, returns an error
func (c *Client) PaylinkCreate(body paylink.PaylinkBody) (*paylink.Paylink, error) {
	req, err := c.NewRequest(http.MethodPost, "Invoice", body)
	if err != nil {
		return nil, err
	}

	resp := &Response[paylink.Paylink]{}
	err = c.Send(req, resp)
	if err != nil {
		return nil, err
	}

	if resp.Status != RequestStatusSuccess {
		return nil, fmt.Errorf("failed to create Paylink: %s", resp.Message)
	}

	return &resp.Data[0], err
}

// PaylinkDelete deletes an existing Paylink or returns an error
func (c *Client) PaylinkDelete(id int32) error {
	req, err := c.NewRequest(http.MethodDelete, fmt.Sprintf("Invoice/%d", id), nil)
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
		return fmt.Errorf("failed to create Paylink: %s", resp.Message)
	}

	return nil
}

// PaylinkRetrieve retrieves a Paylink by its ID or returns an error otherwise
//
// Returns error payrexxsdk.ResourceNotFoundAPIError if not found
func (c *Client) PaylinkRetrieve(id int32) (*paylink.Paylink, error) {
	req, err := c.NewRequest(http.MethodGet, fmt.Sprintf("Invoice/%d", id), nil)
	if err != nil {
		return nil, err
	}

	resp := &Response[paylink.Paylink]{}
	err = c.Send(req, resp)
	if err != nil {
		return nil, err
	}

	if resp.Status != RequestStatusSuccess {
		return nil, ResourceNotFoundAPIError
	}

	return &resp.Data[0], err
}
