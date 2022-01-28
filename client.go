package hetzner

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

const (
	// UserAgent contains user agent identifier
	UserAgent = "intboat/go-hetzner"

	// DefaultEndpoint to be used
	DefaultEndpoint = "https://robot-ws.your-server.de"
)

type Client struct {
	// HTTP client used to communicate with Hetzner API.
	client *http.Client

	// Base URL for API requests.
	BaseURL string

	// User agent for client
	UserAgent string

	// Debug will dump http request and response
	Debug bool

	headers map[string]string

	Boot     BootService
	Ordering OrderingService
	Reset    ResetService
	Server   ServerService
	SSHKey   SSHKeyService
	Failover FailoverService
	IP       IPService
	Subnet   SubnetService
	Wol      WOLService
	Traffic  TrafficService
	RDNS     RDNSService
}

func NewClient(username, password string) *Client {
	c := &Client{}
	c.client = &http.Client{Timeout: time.Second * 120}
	c.BaseURL = DefaultEndpoint
	c.headers = map[string]string{
		"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte(username+":"+password)),
	}
	c.UserAgent = UserAgent

	c.Boot = &BootServiceImpl{client: c}
	c.Ordering = &OrderingServiceImpl{client: c}
	c.Reset = &ResetServiceImpl{client: c}
	c.Server = &ServerServiceImpl{client: c}
	c.SSHKey = &SSHKeyServiceImpl{client: c}
	c.Failover = &FailoverServiceImpl{client: c}
	c.IP = &IPServiceImpl{client: c}
	c.Subnet = &SubnetServiceImpl{client: c}
	c.Wol = &WOLServiceImpl{client: c}
	c.Traffic = &TrafficServiceImpl{client: c}
	c.RDNS = &RDNSServiceImpl{client: c}
	return c
}

// WithUserAgent Send this API client UserAgent to server
func (c *Client) WithUserAgent(ua string) *Client {
	c.UserAgent = ua
	return c
}

// WithTimeout Set a custom timeout
func (c *Client) WithTimeout(timeout time.Duration) *Client {
	c.client.Timeout = timeout
	return c
}

// NewRequest creates an API request. A relative URL can be provided in urlStr, which will be resolved to the BaseURL
// of the Client. If specified, the value pointed to by request is www-form-urlencoded and included in as the request body.
func (c *Client) NewRequest(method, path string, request interface{}) (*http.Request, error) {
	var u *url.URL
	var err error

	if c.BaseURL == "" {
		return nil, errors.New("baseURL must not be empty")
	}

	u, err = url.Parse(c.BaseURL)
	if err != nil {
		return nil, err
	}

	qv, err := query.Values(request)
	if err != nil {
		return nil, err
	}

	qs := encode(qv)
	if qs != "" {
		if method == http.MethodGet || method == http.MethodDelete {
			if strings.Contains(path, "?") {
				path += "&" + qs
			} else {
				path += "?" + qs
			}
		}
	}

	if path != "" {
		rel, err := url.Parse(path)
		if err != nil {
			return nil, err
		}
		if u != nil {
			u = u.ResolveReference(rel)
		} else {
			u = rel
		}
	}
	if u == nil {
		return nil, errors.New("no URL is provided")
	}

	req, err := http.NewRequest(method, u.String(), bytes.NewBufferString(qs))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", c.UserAgent)
	for k, v := range c.headers {
		req.Header.Add(k, v)
	}
	if method == http.MethodPost || method == http.MethodPut {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Length", strconv.Itoa(len(qs)))
	}
	if c.Debug {
		dump, err := httputil.DumpRequestOut(req, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	return req, nil
}

// Do send an API request and returns the API response. The API response is JSON decoded and stored in the value
// pointed to by v, or returned as an error if an API error has occurred.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	var resp *http.Response
	var err error

	if c.Debug {
		fmt.Println(req.URL.String())
	}

	resp, err = c.client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if c := resp.StatusCode; c == 500 || c >= 502 && c <= 599 {
		return nil, errors.New("5xx Server Error")
	}

	if c.Debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}

	err = checkResponse(resp)
	if err != nil {
		return resp, err
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
		if err != nil {
			return nil, err
		}
	}
	return resp, err
}

// CheckResponse checks the API response for errors, and returns them if present. A response is considered an
// error if it has a status code outside the 200 range. API error responses are expected to have either no response
// body, or a JSON response body that maps to APIError.
func checkResponse(r *http.Response) error {
	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	apiErr := &APIError{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && len(data) > 0 {
		type E struct {
			Error *APIError `json:"error"`
		}
		e := E{}
		if err = json.Unmarshal(data, &e); err == nil {
			e.Error.Response = r
			apiErr = e.Error
		}
	}
	return apiErr
}

// Call Send request to API
func (c *Client) Call(method, path string, reqBody, resType interface{}) (*http.Response, error) {
	req, err := c.NewRequest(method, path, reqBody)
	if err != nil {
		return nil, err
	}
	return c.Do(req, resType)
}

// Encode encodes the values into ``URL encoded'' form
// ("bar=baz&foo=bar") sorted by key.
// Unlike std lib, this avoids escaping keys.
func encode(v url.Values) string {
	if v == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		prefix := k + "=" // QueryEscape(k)
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(prefix)
			buf.WriteString(url.QueryEscape(v))
		}
	}
	return buf.String()
}
