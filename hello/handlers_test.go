package main

import (
	"net/http/httptest"
	"testing"
	"io/ioutil"
	"strings"
)

/**
 Below are some basic tests for the handlers used in this tutorial:
 https://thenewstack.io/make-a-restful-json-api-go/
*/

func TestIndex(t *testing.T) {
	tests := []struct {
		description string
		endpoint string
		expectedCode int
		expectedOutput string
	}{
		{
			description: "success",
			endpoint: "/",
			expectedCode: 200,
			expectedOutput: "Welcome!",
		},
	}

	for _, tc := range tests {
		req := httptest.NewRequest("GET", tc.endpoint, nil) // creating an http request
		w := httptest.NewRecorder() // creating a response writer

		Index(w, req) // calling the handler

		if tc.expectedCode != w.Code {
			t.Fatalf("Expected %s, but got %s", tc.expectedCode, w.Code)
		}

		b, err := ioutil.ReadAll(w.Body)
		if err != nil {
			t.Fatalf("Error getting response body: %+v", err)
		}

		output := strings.TrimSpace(string(b))

		if tc.expectedOutput != output {
			t.Fatalf("Expected %s, but got %s", tc.expectedOutput, output)
		}
	}
}

func TestTodoIndex(t *testing.T) {
	tests := []struct {
		description string
		endpoint string
		expectedCode int
		expectedOutput string
	}{
		{
			description: "success",
			endpoint: "/todos",
			expectedCode: 200,
			expectedOutput: `[{"id":0,"name":"Write presentation","completed":false,"due":"0001-01-01T00:00:00Z"},{"id":0,"name":"Host meetup","completed":false,"due":"0001-01-01T00:00:00Z"}]`,
		},

	}

	for _, tc := range tests {
		req  := httptest.NewRequest("GET", "/todos", nil)
		w := httptest.NewRecorder()

		TodoIndex(w, req)

		if tc.expectedCode != w.Code {
			t.Fatalf("Expected %s, but got %s", tc.expectedCode, w.Code)
		}

		b, err := ioutil.ReadAll(w.Body)
		if err != nil {
			t.Fatalf("Error getting response body: %+v", err)
		}

		output := strings.TrimSpace(string(b))

		if tc.expectedOutput != output {
			t.Fatalf("Expected %s, but got %s", tc.expectedOutput, output)
		}
	}
}
