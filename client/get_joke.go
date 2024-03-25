package client

import (
	"encoding/json"
	"io"
	"net/http"
)

type GetJokeResponse struct {
	Id         string `json:"id"`
	Text       string `json:"joke"`
	StatusCode int    `json:"status"`
}

type Joke struct {
	Id   string `json:"id"`
	Text string `json:"joke"`
}

func GetJoke(id string) (Joke, error) {
	req, err := http.NewRequest("GET", JokeURL(id), nil)
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
		return Joke{}, &ResponseStatusError{StatusCode: resp.StatusCode}
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return Joke{}, err
	}
	var jokeResp GetJokeResponse
	err = json.Unmarshal(bodyBytes, &jokeResp)
	if err != nil {
		return Joke{}, err
	}

	if jokeResp.StatusCode < http.StatusOK || jokeResp.StatusCode >= http.StatusMultipleChoices {
		return Joke{}, &ResponseStatusError{StatusCode: jokeResp.StatusCode}
	}

	joke := Joke{jokeResp.Id, jokeResp.Text}

	return joke, nil
}
