package ixon

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	UserAgent           = "ixapi-go v1"
	CompanyHeader       = "IXapi-Company"
	ApplicationHeader   = "IXapi-Application"
	ApiVersionHeader    = "IXapi-Version"
	AuthorizationHeader = "Authorization"
	UserAgentHeader     = "User-Agent"
	ContentTypeHeader   = "Content-Type"
	ContentType         = "application/json"
	ApiVersion          = 1
)

type Client struct {
	client       *http.Client
	apiEndpoints map[string]string

	applicationId string
	companyId     string
	accessToken   string

	authenticated bool

	common service

	Discovery *DiscoveryService
	Auth      *AuthService
	Agent     *AgentService
}

type service struct {
	client *Client
}

func NewClient(applicationId string, companyId string) *Client {
	httpClient := http.DefaultClient

	c := &Client{client: httpClient, applicationId: applicationId, companyId: companyId}

	c.common.client = c

	c.Discovery = (*DiscoveryService)(&c.common)
	c.Auth = (*AuthService)(&c.common)
	c.Agent = (*AgentService)(&c.common)

	return c
}

func (c *Client) MakeRequest(method string, endpointUrl string, customHeaders map[string]string,
	body interface{}, customQuery map[string]string) (*http.Response, error) {

	var buf io.ReadWriter

	// Encode body
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	// Encode query params
	params := url.Values{}

	for k, v := range customQuery {
		params.Add(k, v)
	}

	encodedParams := params.Encode()

	urlWithParams := endpointUrl

	if len(encodedParams) != 0 {
		urlWithParams = fmt.Sprintf("%s?%s", endpointUrl, encodedParams)
	}

	// Create request
	req, err := http.NewRequest(method, urlWithParams, buf)

	if err != nil {
		return nil, err
	}

	// Add required headers
	req.Header.Add(ApiVersionHeader, strconv.Itoa(ApiVersion))
	req.Header.Add(ApplicationHeader, c.applicationId)
	req.Header.Add(UserAgentHeader, UserAgent)
	req.Header.Add(ContentTypeHeader, ContentType)

	// Add authorization and company header if authenticated
	if c.authenticated {
		req.Header.Add(CompanyHeader, c.companyId)
		req.Header.Add(AuthorizationHeader, fmt.Sprintf("Bearer %s", c.accessToken))
	}

	// Add custom headers
	for k, v := range customHeaders {
		req.Header.Add(k, v)
	}

	// Make request
	res, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) ParseEndpoint(name string, replacements map[string]string) (string, error) {
	endpoint := c.apiEndpoints[name]

	if endpoint == "" {
		return "", errors.New("could not find endpoint")
	}

	for k, v := range replacements {
		endpoint = strings.Replace(endpoint, fmt.Sprintf("{%s}", k), v, -1)
	}

	return endpoint, nil
}

func (c *Client) LogApiError(res *http.Response, message string) {
	var parsed ApiErrorResponse

	err := json.NewDecoder(res.Body).Decode(&parsed)

	if err != nil {
		return
	}

	errorString := ""

	for i, err := range parsed.Data {
		errorString = errorString + err.Message
		if i != len(parsed.Data)-1 {
			errorString = errorString + ", "
		}
	}

	logrus.WithField("error", errorString).Error(message)
}
