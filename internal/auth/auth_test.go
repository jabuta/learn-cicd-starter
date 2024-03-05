package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	header := http.Header{"Authorization": []string{"ApiKey gettestkey"}}
	if key, err := GetAPIKey(header); err != nil {
		t.Errorf("Expected no error, got %s.", err)
	} else if key != "gettestkey" {
		t.Errorf("Wrong key extraction, expected: 'gettestkey'; got: %s.", key)
	}
	header = http.Header{}
	if _, err := GetAPIKey(header); !errors.Is(err, ErrNoAuthHeaderIncluded) {
		t.Errorf("Expected %s, got %s.", ErrNoAuthHeaderIncluded, err)
	}
	header = http.Header{"Authorization": []string{"Bearer gettestkey"}}
	if _, err := GetAPIKey(header); err.Error() != "malformed authorization header" {
		t.Errorf("Expected 'malformed authorization header', got %s", err)
	}
	header = http.Header{"Authorization": []string{"Bearer gettestkey asdf"}}
	if _, err := GetAPIKey(header); err.Error() != "malformed authorization header" {
		t.Errorf("Expected 'malformed authorization header', got %s", err)
	}

}
