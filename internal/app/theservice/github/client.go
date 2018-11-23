package github

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

type Zen string

type Client struct {
	GetZen func() (Zen, error)
}

type ClientConfig struct {
	GitHubApiUrl string
}

func NewClient(config *ClientConfig) *Client {

	zenUrl := config.GitHubApiUrl + "/zen"

	return &Client{
		GetZen: newGetZen(zenUrl),
	}
}

func newGetZen(zenUrl string) func() (Zen, error) {

	return func() (Zen, error) {

		resp, err := http.Get(zenUrl)
		if err != nil {
			return "", errors.Wrapf(err, "unable to call endpoint '%s'", zenUrl)
		}

		if resp.StatusCode != http.StatusOK {
			return "", fmt.Errorf("endpoint '%s' respond with an error (%s)", zenUrl, resp.Status)
		}

		defer resp.Body.Close()
		respPayload, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", errors.Wrapf(err, "unable to read response from '%s'", zenUrl)
		}

		return Zen(respPayload), nil
	}
}

