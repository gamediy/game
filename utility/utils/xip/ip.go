package xip

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Location struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	Anycast  bool   `json:"anycast"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
}

func GetIpInfoIo(ctx context.Context, ip string) Location {
	var loc Location

	// Build the URL
	url := fmt.Sprintf("https://ipinfo.io/%s?token=752e995a5aa063", ip)

	// Send GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("error fetching data: %v", err)
		return loc
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error reading response body: %v", err)
		return loc
	}

	// Deserialize JSON response
	if err := json.Unmarshal(body, &loc); err != nil {
		log.Printf("error unmarshaling response: %v", err)
		return loc
	}

	return loc
}
