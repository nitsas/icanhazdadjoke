package client_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	client "nitsas/icanhazdadjoke/client"
)

var _ = Describe("JokeURL", func() {
	When("called with an empty string", func() {
		It("returns the base url (random joke)", func() {
			Expect(client.JokeURL("")).To(Equal(client.BASE_URL))
		})
	})

	When("called with the ID of some joke", func() {
		It("returns the joke url with the given ID", func() {
			Expect(client.JokeURL("abcdef")).To(Equal(client.BASE_URL + "/j/abcdef"))
		})
	})
})
