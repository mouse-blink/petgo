package main

import (
	"fmt"
	"net/http"
	"os"
	"testing"
)

func TestAppRunning(t *testing.T) {
	url := os.Getenv("BASE_URL")
	if url == "" {
		url = "http://localhost:8080"
	}

	resp, err := http.Get(url)
	if err != nil {
		t.Fatalf("Failed to reach %s: %v", url, err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", resp.StatusCode)
	} else {
		fmt.Printf("✅ App is running and returned 200 OK at %s\n", url)
	}
}
