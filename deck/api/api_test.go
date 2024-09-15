package api

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestPut(t *testing.T) {
	// Create a new test server to handle the request
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the HTTP method used
		if r.Method != http.MethodPut {
			t.Errorf("Expected PUT request, got %s", r.Method)
		}

		// Read request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("Error reading request body: %v", err)
		}

		// Check request body content
		expectedPayload := []byte("test payload")
		if string(body) != string(expectedPayload) {
			t.Errorf("Expected request body %s, got %s", expectedPayload, body)
		}

		// Respond with a dummy response
		w.WriteHeader(http.StatusOK)
	}))

	// Close the server when done
	defer server.Close()

	// Create a new instance of DeckAPI with the test server's URL
	apiClient := DeckAPI{
		URL: server.URL,
	}

	// Make the PUT request
	_, err := apiClient.Put("/test", []byte("test payload"))
	if err != nil {
		t.Fatalf("PUT request failed: %v", err)
	}
}

func TestGet(t *testing.T) {
	expectedPayload := []byte("{\"details\":true}")
	expectedResponse := "{\"json\":\"response\"}"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("Error reading request body: %v", err)
		}

		if string(body) != string(expectedPayload) {
			t.Errorf("Expected request body %s, got %s", expectedPayload, body)
		}

		w.Write([]byte("{\"json\":\"response\"}"))
	}))
	defer server.Close()

	apiClient := DeckAPI{
		URL: server.URL,
	}

	resp, err := apiClient.Get("/test", nil)
	if err != nil {
		t.Fatalf("GET request failed: %v", err)
	}

	if string(resp) != expectedResponse {
		t.Fatalf("GET request failed with response %s, expected %s.",
			resp, expectedResponse)
	}
}

func TestPost(t *testing.T) {
	expectedPayload := []byte("test payload")
	expectedResponse := "{\"json\":\"response\"}"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("Error reading request body: %v", err)
		}

		if string(body) != string(expectedPayload) {
			t.Errorf("Expected request body %s, got %s", expectedPayload, body)
		}

		w.Write([]byte("{\"json\":\"response\"}"))
	}))
	defer server.Close()

	apiClient := DeckAPI{
		URL: server.URL,
	}

	resp, err := apiClient.Post("/test", []byte("test payload"))
	if err != nil {
		t.Fatalf("GET request failed: %v", err)
	}

	if string(resp) != expectedResponse {
		t.Fatalf("GET request failed with response %s, expected %s.",
			resp, expectedResponse)
	}
}

func TestDelete(t *testing.T) {
	expectedPayload := []byte("test payload")
	expectedResponse := "{\"json\":\"response\"}"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE request, got %s", r.Method)
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("Error reading request body: %v", err)
		}

		if string(body) != string(expectedPayload) {
			t.Errorf("Expected request body %s, got %s", expectedPayload, body)
		}

		w.Write([]byte("{\"json\":\"response\"}"))
	}))
	defer server.Close()

	apiClient := DeckAPI{
		URL: server.URL,
	}

	resp, err := apiClient.Delete("/test", []byte("test payload"))
	if err != nil {
		t.Fatalf("GET request failed: %v", err)
	}

	if string(resp) != expectedResponse {
		t.Fatalf("GET request failed with response %s, expected %s.",
			resp, expectedResponse)
	}
}

type MockHttpHelper struct {
	mock.Mock
}

func (m *MockHttpHelper) newRequest(method string, url string,
	payload []byte) (*http.Request, error) {

	args := m.Called(method, url, payload)
	return args.Get(0).(*http.Request), args.Error(1)
}

func (m *MockHttpHelper) do(client *http.Client,
	req *http.Request) (*http.Response, error) {

	args := m.Called(client, req)
	return args.Get(0).(*http.Response), args.Error(1)
}

// If everything goes alright
func TestRequest(t *testing.T) {
	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the request method and URL
		if r.Method != "POST" {
			t.Errorf("Expected method POST, got %s", r.Method)
		}
		if r.URL.Path != "/test" {
			t.Errorf("Expected URL /test, got %s", r.URL.Path)
		}

		// Write a mock response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("mock response"))
	}))
	defer server.Close()

	// Create a new instance of DeckAPI
	api := DeckAPI{
		URL: server.URL,
	}

	// Make a request to the mock server
	h := httpHelper{}
	resp, err := api.request(server.URL+"/test", http.MethodPost, []byte("test payload"), h)
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}

	expected := "mock response"
	if string(resp) != expected {
		t.Errorf("Expected response %q, got %q", expected, string(resp))
	}
}

// If it fails to create a new request
func TestRequest_FailsToCreateRequest(t *testing.T) {
	api := DeckAPI{
		URL: "http://example.com/test1",
	}
	mockHelper := new(MockHttpHelper)

	expectedReq1, _ := http.NewRequest("GET",
		"http://example.com/test1", bytes.NewBuffer([]byte("test payload")),
	)
	mockHelper.
		On("newRequest", "GET", "http://example.com/test1", []byte("test payload")).
		Return(expectedReq1, errors.New("An error"))

	_, err := api.request("http://example.com/test1", "GET", []byte("test payload"), mockHelper)
	if err == nil {
		t.Fatalf("First request succeed when it should've failed: %v", err)
	}
}

// This is not finished yet
func TestRequest_FailsToDoRequest(t *testing.T) {
	api := DeckAPI{
		URL: "http://example.com/test1",
	}
	mockHelper := new(MockHttpHelper)

	expectedReq1, _ := http.NewRequest("GET",
		"http://example.com/test1", bytes.NewBuffer([]byte("test payload")),
	)

	mockHelper.
		On("newRequest", "GET", "http://example.com/test1", []byte("test payload")).
		Return(expectedReq1, nil)
	mockHelper.
		On("do", mock.Anything, mock.Anything).
		Return(&http.Response{}, errors.New("Failed in the do function."))

	_, err := api.request("http://example.com/test1", "GET", []byte("test payload"), mockHelper)
	if err == nil {
		t.Fatalf("First request succeed when it should've failed: %v", err)
	}
}

func TestRequest_RequestIsNotSuccessful(t *testing.T) {
	api := DeckAPI{
		URL: "http://example.com/test1",
	}
	mockHelper := new(MockHttpHelper)

	expectedRequest, _ := http.NewRequest("GET",
		"http://example.com/test1", bytes.NewBuffer([]byte("test payload")),
	)
	expectedResponse := &http.Response{
		StatusCode: 405,
		Body:       io.NopCloser(bytes.NewBuffer([]byte("test response"))),
	}

	mockHelper.
		On("newRequest", "GET", "http://example.com/test1", []byte("test payload")).
		Return(expectedRequest, nil)
	mockHelper.
		On("do", mock.Anything, expectedRequest).
		Return(expectedResponse, nil)

	_, err := api.request("http://example.com/test1", "GET", []byte("test payload"), mockHelper)
	if err == nil {
		t.Fatalf("Request shoud've failed: %v", err)
	}
}

// If it fails to read the response
type ErrorReader struct{}

func (e *ErrorReader) Read(p []byte) (int, error) {
	return 0, errors.New("Failed to read response body")
}
func (e *ErrorReader) Close() error {
	return nil
}

func TestRequest_FailsToReadResponse(t *testing.T) {
	api := DeckAPI{
		URL: "http://example.com/test1",
	}
	mockHelper := new(MockHttpHelper)

	expectedRequest, _ := http.NewRequest("GET",
		"http://example.com/test1", bytes.NewBuffer([]byte("test payload")),
	)
	expectedResponse := &http.Response{
		StatusCode: 200,
		Body:       &ErrorReader{},
	}

	mockHelper.
		On("newRequest", "GET", "http://example.com/test1", []byte("test payload")).
		Return(expectedRequest, nil)
	mockHelper.
		On("do", mock.Anything, expectedRequest).
		Return(expectedResponse, nil)

	_, err := api.request("http://example.com/test1", "GET", []byte("test payload"), mockHelper)
	if err == nil {
		t.Fatalf("Request shoud've failed: %v", err)
	}
}
