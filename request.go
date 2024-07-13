package payrexxsdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type RequestStatus string

const (
	RequestStatusError   RequestStatus = "error"
	RequestStatusSuccess RequestStatus = "success"
)

type Response[T any] struct {
	Status  RequestStatus `json:"status"`
	Message string        `json:"message"`
	Data    []T           `json:"data"`
}

// Send makes a request to the API, the response body will be unmarshalled into v, or if v
// is in an io.Writer, the response will be written to it without decoding
func (c *Client) Send(req *http.Request, v interface{}) error {
	var (
		err  error
		resp *http.Response
	)

	// set default headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", c.Secret)

	resp, err = c.Client.Do(req)
	c.log(req, resp)

	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) error {
		return Body.Close()
	}(resp.Body)

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return errors.New(fmt.Sprintf("Non-200 status code received from Payrexx (%d). %s %s", resp.StatusCode, req.Method, req.URL))
	}

	if v == nil {
		return nil
	}

	if w, ok := v.(io.Writer); ok {
		_, err := io.Copy(w, resp.Body)
		return err
	}

	return json.NewDecoder(resp.Body).Decode(v)
}

// NewRequest creates a request which can be modified and sent later on.
func (c *Client) NewRequest(method, endpoint string, payload interface{}) (*http.Request, error) {
	var buf io.Reader
	if payload != nil {
		b, err := json.Marshal(&payload)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s/", c.ApiUrl, endpoint), buf)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("instance", c.InstanceName)

	req.URL.RawQuery = q.Encode()
	return req, nil
}
