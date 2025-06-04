package payrexxsdk

import (
	"bytes"
	"context"
	"encoding/json"
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

type APIError struct {
	StatusCode int
	Body       []byte
	Err        error
}

func (e *APIError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("status %d: %v", e.StatusCode, e.Err)
	}
	return fmt.Sprintf("status %d: %s", e.StatusCode, string(e.Body))
}

// Send makes a request to the API, the response body will be unmarshalled into v, or if v
// is in an io.Writer, the response will be written to it without decoding
func (c *Client) Send(req *http.Request, v interface{}) (err error) {
	var resp *http.Response

	// set default headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", c.Secret)

	resp, err = c.Client.Do(req)

	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		data, _ := io.ReadAll(resp.Body)

		if resp.StatusCode == 422 {
			return &APIError{
				StatusCode: resp.StatusCode,
				Body:       data,
				Err:        fmt.Errorf("wrong instance name"),
			}
		} else if resp.StatusCode == 403 {
			return &APIError{
				StatusCode: resp.StatusCode,
				Body:       data,
				Err:        fmt.Errorf("wrong api secret"),
			}
		} else {
			return &APIError{
				StatusCode: resp.StatusCode,
				Body:       data,
			}
		}
	}

	if v == nil {
		return nil
	}

	if w, ok := v.(io.Writer); ok {
		_, err := io.Copy(w, resp.Body)
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))

	if err := json.Unmarshal(body, v); err != nil {
		return fmt.Errorf("failed to decode JSON (bytes=%q): %w", string(body), err)
	}
	return nil
}

// NewRequest creates a request which can be modified and sent later on.
func (c *Client) NewRequest(
	ctx context.Context,
	method, endpoint string,
	payload interface{},
) (*http.Request, error) {
	var buf io.Reader
	if payload != nil {
		b, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
	}

	url := fmt.Sprintf("%s/%s/", c.ApiUrl, endpoint)
	req, err := http.NewRequestWithContext(ctx, method, url, buf)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("instance", c.InstanceName)
	req.URL.RawQuery = q.Encode()

	return req, nil
}
