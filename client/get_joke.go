package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const BASE_URL = "https://icanhazdadjoke.com/"
const USER_AGENT = "https://github.com/nitsas/icanhazdadjoke"

type JokeResponse struct {
	Id         string `json:"id"`
	Text       string `json:"joke"`
	StatusCode int    `json:"status"`
}

type Joke struct {
	Id   string `json:"id"`
	Text string `json:"joke"`
}

type ResponseStatusError struct {
	StatusCode int
}

func (e *ResponseStatusError) Error() string {
	return fmt.Sprintf("Got response status %d (%s)", e.StatusCode, http.StatusText(e.StatusCode))
}

func GetJoke(id string) (Joke, error) {
	req, err := http.NewRequest("GET", jokeURL(id), nil)
	if err != nil {
		return Joke{}, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", USER_AGENT)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Joke{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return Joke{}, &ResponseStatusError{resp.StatusCode}
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return Joke{}, err
	}
	var jokeResp JokeResponse
	err = json.Unmarshal(bodyBytes, &jokeResp)
	if err != nil {
		return Joke{}, err
	}

	if jokeResp.StatusCode < http.StatusOK || jokeResp.StatusCode >= http.StatusMultipleChoices {
		return Joke{}, &ResponseStatusError{jokeResp.StatusCode}
	}

	joke := Joke{jokeResp.Id, jokeResp.Text}

	return joke, nil
}

func jokeURL(id string) string {
	var url string

	if id == "" {
		url = BASE_URL
	} else {
		url = fmt.Sprintf("%s/j/%s", BASE_URL, id)
	}

	return url
}
