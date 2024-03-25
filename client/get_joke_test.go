package client_test

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	client "nitsas/icanhazdadjoke/client"
)

var _ = Describe("GetJoke", func() {
	var c *client.Client

	var mockServer *httptest.Server
	var mockHandler *http.HandlerFunc
	var mockResponse []byte

	JustBeforeEach(func() {
		mockServer = httptest.NewServer(*mockHandler)
		c = &client.Client{BaseURL: mockServer.URL}
	})

	AfterEach(func() {
		mockServer.Close()
	})

	When("asking for a random Joke (param id == \"\")", func() {
		BeforeEach(func() {
			mockResponse = []byte(`{"id":"abcd","joke":"some joke","status":200}`)
			mockResponseStatus := 200

			mockHandlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				Expect(r.Method).To(Equal("GET"), "performs a GET request")
				Expect(r.URL.Path).To(Equal("/"), "performs a request to the root URL")

				w.WriteHeader(mockResponseStatus)
				n, err := w.Write(mockResponse)
				Expect(err).NotTo(HaveOccurred())
				Expect(n).To(Equal(len(mockResponse)))
			})
			mockHandler = &mockHandlerFunc
		})

		It("performs the request correctly and returns a correct Joke", func() {
			joke, err := c.GetJoke("")
			Expect(err).NotTo(HaveOccurred())
			Expect(joke.Id).To(Equal("abcd"))
			Expect(joke.Text).To(Equal("some joke"))
		})

		When("the server responds with HTTP status 200 but the response body has status 500", func() {
			BeforeEach(func() {
				mockResponse = []byte(`{"id":"","joke":"","status":500}`)
			})

			It("returns an error", func() {
				_, err := c.GetJoke("")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("500"))
				Expect(err.Error()).To(ContainSubstring("Internal Server Error"))
			})
		})
	})

	When("asking for a specific Joke by id (param id == \"some_id\")", func() {
		BeforeEach(func() {
			mockResponse = []byte(`{"id":"some_id","joke":"some specific joke","status":200}`)
			mockResponseStatus := 200

			mockHandlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				Expect(r.Method).To(Equal("GET"), "performs a GET request")
				Expect(r.URL.Path).To(Equal("/j/some_id"), "performs a request to the specific joke's URL")

				w.WriteHeader(mockResponseStatus)
				n, err := w.Write(mockResponse)
				Expect(err).NotTo(HaveOccurred())
				Expect(n).To(Equal(len(mockResponse)))
			})
			mockHandler = &mockHandlerFunc
		})

		It("performs the request correctly and returns a correct Joke", func() {
			joke, err := c.GetJoke("some_id")
			Expect(err).NotTo(HaveOccurred())
			Expect(joke.Id).To(Equal("some_id"))
			Expect(joke.Text).To(Equal("some specific joke"))
		})

		When("the server responds with HTTP status 200 but the response body has status 404", func() {
			BeforeEach(func() {
				mockResponse = []byte(`{"id":"","joke":"","status":404}`)
			})

			It("returns an error", func() {
				_, err := c.GetJoke("some_id")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("404"))
				Expect(err.Error()).To(ContainSubstring("Not Found"))
			})
		})
	})
})
