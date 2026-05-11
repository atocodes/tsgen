package network

import (
	"net/http"
	"time"
)

func HasInternetConnection() bool {
	client := http.Client{
		Timeout: 3 * time.Second,
	}

	resp, err := client.Get("https://atocodes-portfolio.vercel.app/")

	if err != nil {
		return false
	}

	defer resp.Body.Close()

	return resp.StatusCode == 204
}
