package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

type MockProcessList struct {
	listFunc func(*MockProcessList) []*Process
}

func (l *MockProcessList) List() []*Process {
	return l.listFunc(l)
}

func TestHTTPHandlerServeHTTP(t *testing.T) {
	// initialize the mock process list with custom list function
	processes := []*Process{}
	processList := &MockProcessList{
		listFunc: func(l *MockProcessList) []*Process {
			return processes
		},
	}

	// initialize the http handler for testing
	handler := &HTTPHandler{processLister: processList}

	// initialize a mock http request
	var requestBody bytes.Buffer
	req := httptest.NewRequest("GET", "/ps", &requestBody)

	// initialize a mock http response writer
	recorder := httptest.NewRecorder()

	// run the actual ServeHTTP function
	handler.ServeHTTP(recorder, req)

	// pull out the http response from the recorder
	resp := recorder.Result()

	// read the body into an array of bytes
	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	// marshal the expected response for comparison
	expectedRespBody, err := json.Marshal(processes)

	if err != nil {
		t.Fatal(err)
	}

	// compare the response body to the expected response body
	if bytes.Compare(respBody[:len(respBody)-1], expectedRespBody) != 0 {
		t.Fatal("did not match\nexpected: %s\ngot: %s\n", expectedRespBody, respBody)
	}
}
