package velora

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/auroradevllc/apiclient"
)

const veloraBase = "https://api.velora.tv"

func New() Client {
	return &client{}
}

type Client interface {
	Stream(username string) (*UserStream, error)
	LiveStreams() ([]ListStream, error)
	FeaturedStreams() ([]ListStream, error)
	TopStreams() ([]ListStream, error)
	GrowingStreams() ([]ListStream, error)
	Categories() ([]Category, error)
	User(username string) (*User, error)
}

var (
	_           Client = (*client)(nil)
	ErrNotFound        = errors.New("resource not found")
)

type client struct {
	token string
}

// Stream retrieves a stream for a user
// Endpoint: /api/streams/user/<username>
func (c *client) Stream(username string) (*UserStream, error) {
	var stream UserStream

	if err := c.sendRequestJSON("/api/streams/user/"+url.PathEscape(username), nil, &stream); err != nil {
		return nil, err
	}

	return &stream, nil
}

// LiveStreams retrieves all streams that are live
// Endpoint: /api/streams/live
func (c *client) LiveStreams() ([]ListStream, error) {
	var res ListStreamResponse

	if err := c.sendRequestJSON("/api/streams/live", nil, &res); err != nil {
		return nil, err
	}

	return res.Streams, nil
}

// FeaturedStreams retrieves all featured streams
// Endpoint: /api/streams/featured
func (c *client) FeaturedStreams() ([]ListStream, error) {
	var res ListStreamResponse

	if err := c.sendRequestJSON("/api/streams/featured", nil, &res); err != nil {
		return nil, err
	}

	return res.Streams, nil
}

// TopStreams retrieves the top streams
func (c *client) TopStreams() ([]ListStream, error) {
	var res ListStreamResponse

	if err := c.sendRequestJSON("/api/streams/top", nil, &res); err != nil {
		return nil, err
	}

	return res.Streams, nil
}

// GrowingStreams returns the "growing" stream category
func (c *client) GrowingStreams() ([]ListStream, error) {
	var res ListStreamResponse

	if err := c.sendRequestJSON("/api/streams/growing", nil, &res); err != nil {
		return nil, err
	}

	return res.Streams, nil
}

// Categories retrieves all known categories on the platform
func (c *client) Categories() ([]Category, error) {
	var categories []Category

	if err := c.sendRequestJSON("/api/categories", nil, &categories); err != nil {
		return nil, err
	}

	return categories, nil
}

// User retrieves user information for a username
func (c *client) User(username string) (*User, error) {
	var user User

	if err := c.sendRequestJSON("/api/users/"+url.PathEscape(username), nil, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (c *client) sendRequestJSON(path string, body, out any) error {
	res, err := c.sendRequest(path, body, out)

	if err != nil {
		return err
	}

	return res.Unmarshal(out)
}

func (c *client) sendRequest(path string, body any, out any) (*apiclient.Response, error) {
	var opts []apiclient.Option

	if c.token != "" {
		opts = append(opts, apiclient.WithHeader("Authorziation", "Bearer "+c.token))
	}

	if body != nil {
		opts = append(opts, apiclient.WithBody(body))
	}

	req, err := apiclient.NewRequest(veloraBase+path, opts...)

	if err != nil {
		return nil, err
	}

	res, err := req.Send()

	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusNotFound {
		defer res.Close()
		return nil, ErrNotFound
	}

	return res, err
}
