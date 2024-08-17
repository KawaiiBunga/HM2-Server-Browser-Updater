package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://master.iw4.zip/servers#"

	fmt.Println("Grabbing servers for H2M...")

	// Send a GET request to the webpage
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to fetch the webpage:", err)
		return
	}
	defer resp.Body.Close()

	// Parse the HTML content of the page using goquery
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("Failed to parse HTML:", err)
		return
	}

	// List to store IP:Port pairs
	var ipPortList []string

	// Find the H2M tab by its ID and then find all server rows within it
	doc.Find("#H2M_servers .server-row").Each(func(i int, s *goquery.Selection) {
		ip, exists := s.Attr("data-ip")
		if !exists {
			return
		}
		port, exists := s.Attr("data-port")
		if !exists {
			return
		}
		ipPortList = append(ipPortList, fmt.Sprintf("%s:%s", ip, port))
	})

	// Convert the list to JSON without spaces
	jsonData, err := json.Marshal(ipPortList)
	if err != nil {
		fmt.Println("Failed to marshal JSON:", err)
		return
	}

	// Save the results to a JSON file
	err = ioutil.WriteFile("favourites.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Failed to write to file:", err)
		return
	}

	fmt.Println("H2M server list has been updated!")

	// Prompt the user to press Enter to exit
    fmt.Println("Press Enter to exit...")
    fmt.Scanln()
}
