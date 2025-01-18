package client

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"time"
)

const (
	maxRetries     = 3
	initialBackoff = 1 * time.Second
	timeout        = 5 * time.Second
)

var client = &http.Client{
	Timeout: timeout,
}

func exponentialBackoff(attempt int) time.Duration {
	backoff := float64(initialBackoff) * math.Pow(2, float64(attempt))
	jitter := time.Duration(float64(time.Millisecond) * (50 + float64(attempt*10)))
	return time.Duration(backoff) + jitter
}

func MakeAPICall(method, url string) (string, error) {
	var responseBody string

	for attempt := 0; attempt <= maxRetries; attempt++ {
		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			return "", fmt.Errorf("failed to create request: %v", err)
		}

		resp, err := client.Do(req)
		if err != nil {
			if attempt < maxRetries {
				log.Printf("Attempt %d failed: %v. Retrying...\n", attempt+1, err)
				time.Sleep(exponentialBackoff(attempt))
				continue
			}
			return "", fmt.Errorf("request failed after %d attempts: %v", attempt+1, err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", fmt.Errorf("failed to read response body: %v", err)
		}
		responseBody = string(body)

		switch resp.StatusCode {
		case http.StatusOK:
			log.Printf("Success: %d\n", resp.StatusCode)
			return responseBody, nil
		case http.StatusBadRequest:
			return "", fmt.Errorf("Bad Request: %d\nBody: %s", resp.StatusCode, responseBody)
		case http.StatusNotFound:
			return "", fmt.Errorf("Not Found: %d\nBody: %s", resp.StatusCode, responseBody)
		case http.StatusTooManyRequests:
			log.Printf("Rate limited: %d\n", resp.StatusCode)
			if attempt < maxRetries {
				time.Sleep(exponentialBackoff(attempt))
				continue
			}
		case http.StatusInternalServerError:
			log.Printf("Server error: %d\n", resp.StatusCode)
			if attempt < maxRetries {
				time.Sleep(exponentialBackoff(attempt))
				continue
			}
		default:
			return "", fmt.Errorf("received status code: %d\nBody: %s", resp.StatusCode, responseBody)
		}
	}

	return responseBody, nil
}
