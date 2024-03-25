package client_test

import (
	"encoding/json"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	client "nitsas/icanhazdadjoke/client"
)

var _ = Describe("GetJokeResponse", func() {
	When("unmarshalling a json with fields id (string), joke (string), status (int)", func() {
		It("sets Id from json id, Text from json joke, and StatusCode from json status", func() {
			jsonBytes := []byte(`{"id":"abcdefg","joke":"Donkeys can fly!","status":200}`)

			var resp client.GetJokeResponse
			Expect(json.Unmarshal(jsonBytes, &resp)).To(Succeed())

			Expect(resp.Id).To(Equal("abcdefg"))
			Expect(resp.Text).To(Equal("Donkeys can fly!"))
			Expect(resp.StatusCode).To(Equal(200))
		})
	})
})
