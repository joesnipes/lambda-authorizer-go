package main

import (
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	var (
		response events.APIGatewayProxyResponse
		request  events.APIGatewayProxyRequest
		err      error
	)

	request.Path = "/dev/hello"
	response, err = handler(request)

	if err != nil {
		t.Errorf("TEST FAILED. Expected %v, got %v", nil, err)
	} else {
		t.Logf("Test PASSED. Expected %v, got %v", nil, err)
	}

	if response.Body != "Hello, World!" {
		t.Errorf("Test FAILED. Expected %s, got %s", "Hello, World!", response.Body)
	} else {
		t.Logf("Test PASSED. Expected %s, got %s", "Hello, World!", response.Body)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("Test FAILED. Expected %d, got %d", 200, response.StatusCode)
	} else {
		t.Logf("Test PASSED. Expected %d, got %d", 200, response.StatusCode)
	}

	request.Path = "/dev/"
	response, _ = handler(request)

	if response.StatusCode != http.StatusOK {
		t.Errorf("Test FAILED. Expected %d, got %d", 200, response.StatusCode)
	} else {
		t.Logf("Test PASSED. Expected %d, got %d", 200, response.StatusCode)
	}

	request.Path = "/fail"
	response, _ = handler(request)

	if response.StatusCode != http.StatusForbidden {
		t.Errorf("Test FAILED. Expected %d, got %d", 403, response.StatusCode)
	} else {
		t.Logf("Test PASSED. Expected %d, got %d", 403, response.StatusCode)
	}
}
