package client_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	client "nitsas/icanhazdadjoke/client"
)

var _ = Describe("ResponseStatusError", func() {
	var statusCodeP *int
	var subject *client.ResponseStatusError

	JustBeforeEach(func() {
		subject = &client.ResponseStatusError{StatusCode: *statusCodeP}
	})

	When("initialized with param 500", func() {
		BeforeEach(func() {
			statusCode := 500
			statusCodeP = &statusCode
		})

		Describe("Error()", func() {
			It("contains the correct status code and status text", func() {
				Expect(subject.Error()).To(ContainSubstring("500"))
				Expect(subject.Error()).To(ContainSubstring("Internal Server Error"))
			})
		})
	})

	When("initialized with param 404", func() {
		BeforeEach(func() {
			statusCode := 404
			statusCodeP = &statusCode
		})

		Describe("Error()", func() {
			It("contains the correct status code and status text", func() {
				Expect(subject.Error()).To(ContainSubstring("404"))
				Expect(subject.Error()).To(ContainSubstring("Not Found"))
			})
		})
	})
})
