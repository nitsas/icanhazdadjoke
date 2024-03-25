package client_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	client "nitsas/icanhazdadjoke/client"
)

var _ = Describe("Client.JokeURL", func() {
	var c *client.Client
	baseURL := "sample_base_url"

	When("the client has some given base url", func() {
		BeforeEach(func() {
			c = &client.Client{BaseURL: baseURL}
		})

		When("called with an empty string", func() {
			It("returns the base url (random joke)", func() {
				Expect(c.JokeURL("")).To(Equal(baseURL))
			})
		})

		When("called with the ID of some joke", func() {
			It("returns the joke url with the given ID", func() {
				Expect(c.JokeURL("abcdef")).To(Equal(baseURL + "/j/abcdef"))
			})
		})
	})
})
