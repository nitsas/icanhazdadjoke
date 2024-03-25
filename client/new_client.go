package client

func NewClient() Client {
	return Client{BaseURL: DEFAULT_BASE_URL}
}
