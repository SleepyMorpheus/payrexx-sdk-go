package payrexxsdk

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"sync"
)

// Client represents a Payrexx REST API Client
type Client struct {
	mu           sync.Mutex
	Client       *http.Client
	LogWriter    io.Writer // if set, all requests will be logged to file. Otherwise stdout
	LogEnabled   bool
	InstanceName string
	Secret       string
	ApiUrl       string
}

const APIBaseDefault = "https://api.payrexx.com/v1.0"

// NewClient returns a new Client struct
// APIBase is a base API URL, you can use payrexxsdk.APIBaseDefault
func NewClient(instanceName string, secret string, APIBase string) (*Client, error) {
	if instanceName == "" || secret == "" {
		return nil, errors.New("InstanceName and Secret cannot be empty")
	}

	return &Client{
		// sync.Mutex
		InstanceName: instanceName,
		Secret:       secret,
		ApiUrl:       APIBase,
		Client:       &http.Client{},
	}, nil
}

func (c *Client) enableLogging() {
	c.LogEnabled = true
}

func (c *Client) disableLogging() {
	c.LogEnabled = false
}

func (c *Client) setLogWriter(w io.Writer) {
	c.LogWriter = w
	if !c.LogEnabled {
		log.Println("Added log writer without enabling logging!")
	}
}

// CheckSignature can be used to check if the provided secret is correct
//
// Eendpoint GET /SignatureCheck/
func (c *Client) CheckSignature() error {
	req, err := c.NewRequest(http.MethodGet, fmt.Sprintf("SignatureCheck"), nil)
	if err != nil {
		return err
	}

	resp := &Response[interface{}]{}
	err = c.Send(req, resp)
	if err != nil {
		return err
	}

	if resp.Status == RequestStatusSuccess {
		return errors.New(resp.Message)
	}

	return nil
}

// log will dump request and response to log file or stdout if not set
func (c *Client) log(req *http.Request, resp *http.Response) {

	if !c.LogEnabled {
		return
	}

	var (
		reqDump  string
		respDump []byte
	)

	// Stringify both request and response
	if req != nil {
		reqDump = fmt.Sprintf("%s %s. Data %s", req.Method, req.URL.String(), req.Form.Encode())
	}
	if resp != nil {
		respDump, _ = httputil.DumpResponse(resp, true)
	}

	// pipe output into correct channels
	if c.LogWriter != nil {
		_, _ = c.LogWriter.Write([]byte(fmt.Sprintf("Request: %s\nResponse: %s\n", reqDump, string(respDump))))
	} else {
		log.Println("Registered")
	}
}
