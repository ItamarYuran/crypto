package crypto

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetTxt() string {
	// Define the URL
	url := "https://cryptopals.com/static/challenge-data/4.txt" // Replace with your target URL

	// Send HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return ""
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return ""
	}

	// Print the text from the URL
	//fmt.Println(string(body))
	return string(body)
}
