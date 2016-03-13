package instagram_test

import (
	"../go-instagram"
	"os"
	"strconv"
	"testing"
)

func fetchEnvVars(t *testing.T) (token string, count int) {
	token = os.Getenv("INSTAGRAM_TEST_TOKEN")
	if len(token) <= 0 {
		t.Fatalf("Test cannot run because cannot get environment variable INSTAGRAM_TEST_TOKEN")
	}
	countStr := os.Getenv("INSTAGRAM_TEST_COUNT")
	if len(countStr) <= 0 {
		t.Fatalf("Test cannot run because cannot get environment variable INSTAGRAM_TEST_COUNT")
	}
	count, err := strconv.Atoi(countStr)
	if err != nil {
		t.Fatalf(err.Error())
	}
	return token, count
}

func TestNewClient(t *testing.T) {
	_, count := fetchEnvVars(t)
	_, err := instagram.NewClient("", count)
	if err == nil {
		t.Errorf("Expected an error, due to missing API Key.")
	}
}

func TestGetRecentMedia(t *testing.T) {
	token, count := fetchEnvVars(t)
	client, _ := instagram.NewClient(token, count)
	_, err := client.GetRecentMedia()
	if err != nil {
		t.Errorf("Did not expect any error, but get: %s", err.Error())
	}
}
