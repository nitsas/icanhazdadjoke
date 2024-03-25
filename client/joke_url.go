package client

import "fmt"

func JokeURL(id string) string {
	var url string

	if id == "" {
		url = BASE_URL
	} else {
		url = fmt.Sprintf("%s/j/%s", BASE_URL, id)
	}

	return url
}
