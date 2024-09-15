package api

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type APIInterface interface {
	request(string, string, []byte) ([]byte, error)
	Get(string, []byte) ([]byte, error)
	Put(string, []byte) ([]byte, error)
	Post(string, []byte) ([]byte, error)
	Delete(string, []byte) ([]byte, error)
}
type DeckAPI struct {
	URL     string
	User    string
	Pass    string
	LogedIn bool
	Token   string
}

type HTTPHelper interface {
	do(*http.Client, *http.Request) (*http.Response, error)
	newRequest(string, string, []byte) (*http.Request, error)
}
type httpHelper struct{}

func (h httpHelper) newRequest(method string, url string,
	payload []byte) (*http.Request, error) {
	return http.NewRequest(method, url, bytes.NewBuffer(payload))
}

func (h httpHelper) do(client *http.Client,
	req *http.Request) (*http.Response, error) {
	return client.Do(req)
}

func (a DeckAPI) request(url string, method string,
	payload []byte, hhelper HTTPHelper) (data []byte, err error) {

	// We pass an interface that implements this function so that we can test it
	// I really don't find this comfy : (
	req, err := hhelper.newRequest(method, url, payload)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	req.SetBasicAuth(a.User, a.Pass)
	req.Header.Set("OCS-APIRequest", "true")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := hhelper.do(client, req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errStr := fmt.Sprintf("Request failed with status code: %d\n", resp.StatusCode)
		return nil, errors.New(errStr)
	}

	rawData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response: ", err)
		return nil, err
	}

	return rawData, err
}

func (a DeckAPI) Get(endpoint string, payload []byte) (data []byte, err error) {
	url := a.URL + endpoint
	h := httpHelper{}
	return a.request(url, http.MethodGet, []byte("{\"details\":true}"), h)
}

func (a DeckAPI) Post(endpoint string, payload []byte) (data []byte, err error) {
	url := a.URL + endpoint
	h := httpHelper{}
	return a.request(url, http.MethodPost, payload, h)
}

func (a DeckAPI) Delete(endpoint string, payload []byte) (data []byte, err error) {
	url := a.URL + endpoint
	h := httpHelper{}
	return a.request(url, http.MethodDelete, payload, h)
}

func (a DeckAPI) Put(endpoint string, payload []byte) (data []byte, err error) {
	url := a.URL + endpoint
	h := httpHelper{}
	return a.request(url, http.MethodPut, payload, h)
}
