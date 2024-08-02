package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func CampaignLister(token string) error {

	url := "https://api.customer.io/v1/campaigns"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Set("Authorization", "Bearer "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return err
	}

	// Convert the body to a string and assign it to a variable
	bodyString := string(body)

	fmt.Println(bodyString)
	return nil
}
