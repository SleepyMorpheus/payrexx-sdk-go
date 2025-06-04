package payrexxsdk

import (
	"context"
	"fmt"
	"github.com/SleepyMorpheus/payrexx-sdk-go/types/transaction"
	"net/http"
)

func (c *Client) TransactionRetrieve(ctx context.Context, id int32) (*transaction.Transaction, error) {
	req, err := c.NewRequest(ctx, http.MethodGet, fmt.Sprintf("Gateway/%d", id), nil)
	if err != nil {
		return nil, err
	}

	resp := &Response[transaction.Transaction]{}
	err = c.Send(req, resp)
	if err != nil {
		return nil, err
	}

	if resp.Status != RequestStatusSuccess {
		return nil, ResourceNotFoundAPIError
	}

	return &resp.Data[0], nil
}

func (c *Client) TransactionRetrieveMany(ctx context.Context, args transaction.RetrieveManyArguments) (*[]transaction.Transaction, error) {
	req, err := c.NewRequest(ctx, http.MethodGet, "Gateway", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	for k, v := range args.ToMap() {
		q.Add(k, fmt.Sprintf("%v", v))
	}
	req.URL.RawQuery = q.Encode()

	resp := &Response[transaction.Transaction]{}
	err = c.Send(req, resp)
	if err != nil {
		return nil, err
	}

	if resp.Status != RequestStatusSuccess {
		return nil, UnknownAPIError
	}

	return &resp.Data, nil
}

func (c *Client) TransactionCashCreate() {
	// todo
}
