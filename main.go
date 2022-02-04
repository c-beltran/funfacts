package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/c-beltran/funfacts/facts/apis"
)

func main() {
	httpClient := http.Client{
		Timeout:   http.DefaultClient.Timeout,
		Transport: http.DefaultTransport,
	}

	apiClient := apis.NewClient(&httpClient, "https://dog-api.kinduff.com")

	dogFact, err := apiClient.FindDogFact(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Dog fact: ", dogFact.Fact)
}
