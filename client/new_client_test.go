package client_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	client "nitsas/icanhazdadjoke/client"
)

var _ = Describe("NewClient", func() {
	It("returns a Client with BaseURL equal to DEFAULT_BASE_URL", func() {
		c := client.NewClient()
		Expect(c.BaseURL).To(Equal(client.DEFAULT_BASE_URL))
	})
})
