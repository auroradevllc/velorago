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
	User(username string) (*User, error)
	Stream(username string) (*Stream, error)
}

var (
	_           Client = (*client)(nil)
	ErrNotFound        = errors.New("resource not found")
)

type client struct {
	token string
}

func (c *client) Stream(username string) (*Stream, error) {
	var stream Stream

	if err := c.sendRequestJSON("/api/streams/user/"+url.PathEscape(username), nil, &stream); err != nil {
		return nil, err
	}

	return &stream, nil
}

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
