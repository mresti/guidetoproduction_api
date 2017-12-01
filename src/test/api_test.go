// +build integration
package test

import (
	"api/app"
	"api/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"fmt"
)

var a app.App

func TestMain(m *testing.M) {
	a = app.App{}
	a.Initialize("", "", "")
	code := m.Run()
	os.Exit(code)
}

func TestGetFavicon(t *testing.T) {
	req, _ := http.NewRequest("GET", "/favicon.ico", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)
}

func TestRootEndpoint(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	if body := response.Body.String(); body != `{"message":"Hello DevFest Cordoba 2017"}` {
		t.Errorf("Expected the message. Got %s", body)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}


func TestGetStats(t *testing.T) {
	want := 1
	req, _ := http.NewRequest("GET", "/stats", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var v models.VisitAPI
	json.Unmarshal(response.Body.Bytes(), &v)
	if v.Visits > want {
		t.Errorf("Expected visits to be %v. Got %v", want, v.Visits)
	}
}


func TestGetNumber(t *testing.T) {
	var tests = []struct {
		inputNumber int
		inputMethod string
		wantHTTTPCode int
		wantNumber int
	}{
		{
			inputNumber: 1, inputMethod: http.MethodGet, wantHTTTPCode: http.StatusOK, wantNumber: 1,
		},
		{
			inputNumber: 333, inputMethod: http.MethodGet, wantHTTTPCode: http.StatusOK, wantNumber: 10,
		},
		{
			inputNumber: -1, inputMethod: http.MethodGet, wantHTTTPCode: http.StatusOK, wantNumber: 10,
		},
		{
			inputNumber: -1, inputMethod: http.MethodPut, wantHTTTPCode: http.StatusNotFound, wantNumber: 0,
		},
	}
	for _, test := range tests {
		endpoint := fmt.Sprintf("/numbers/%v",test.inputNumber)
		inputMessage := fmt.Sprintf("%v:%v",test.inputMethod, endpoint)
		t.Run(inputMessage, func(t *testing.T) {
			req, _ := http.NewRequest(test.inputMethod, endpoint, nil)
			response := executeRequest(req)
			checkResponseCode(t, test.wantHTTTPCode, response.Code)
			var got models.NumberAPI
			json.Unmarshal(response.Body.Bytes(), &got)
			if got.Number != test.wantNumber {
				t.Errorf("Expected number %v. Got '%v'", test.wantNumber, got.Number)
			}
		})
	}
}
