package client

import "fmt"

func JokeURL(id string) string {
	var url string

	if id == "" {
		url = DEFAULT_BASE_URL
	} else {
		url = fmt.Sprintf("%s/j/%s", DEFAULT_BASE_URL, id)
	}

	return url
}
