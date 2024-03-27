package main

import (
	"fmt"
	"os"

	client "nitsas/icanhazdadjoke/client"
)

func main() {
	jokeId := ""
	if len(os.Args) > 1 {
		jokeId = os.Args[1]
	}

	c := client.NewClient()
	joke, err := c.GetJoke(jokeId)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Joke %s:\n%s\n", joke.Id, joke.Text)
}
