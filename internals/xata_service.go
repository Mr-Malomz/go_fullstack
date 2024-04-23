package internals

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var xataAPIKey = GetEnvVariable("XATA_API_KEY")
var baseURL = GetEnvVariable("XATA_DATABASE_URL")

func createRequest(method, url string, bodyData *bytes.Buffer) (*http.Request, error) {
	var req *http.Request
	var err error

	if method == "GET" || method == "DELETE" || bodyData == nil {
		req, err = http.NewRequest(method, url, nil)
	} else {
		req, err = http.NewRequest(method, url, bodyData)
	}

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", xataAPIKey))

	return req, nil
}

func (app *Config) createTodoService(newTodo *TodoRequest) (*TodoResponse, error) {
	createTodo := TodoResponse{}
	jsonData := Todo{
		Description: newTodo.Description,
	}

	bodyData := new(bytes.Buffer)
	json.NewEncoder(bodyData).Encode(jsonData)

	fullURL := fmt.Sprintf("%s:main/tables/Todo/data", baseURL)

	req, err := createRequest("POST", fullURL, bodyData)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&createTodo); err != nil {
		return nil, err
	}

	return &createTodo, nil
}

func (app *Config) deleteTodoService(id string) (string, error) {
	fullURL := fmt.Sprintf("%s:main/tables/Todo/data/%s", baseURL, id)

	client := &http.Client{}
	req, err := createRequest("DELETE", fullURL, nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return id, nil
}

func (app *Config) getAllTodosService() ([]*Todo, error) {
	var todos []*Todo

	fullURL := fmt.Sprintf("%s:main/tables/Todo/query", baseURL)
	client := &http.Client{}
	req, err := createRequest("POST", fullURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response struct {
		Records []*Todo `json:"records"`
	}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&response); err != nil {
		return nil, err
	}

	todos = response.Records

	return todos, nil
}
