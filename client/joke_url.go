package client

import "fmt"

func (c *Client) JokeURL(id string) string {
	var url string

	if id == "" {
		url = c.BaseURL
	} else {
		url = fmt.Sprintf("%s/j/%s", c.BaseURL, id)
	}

	return url
}
