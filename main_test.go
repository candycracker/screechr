package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Response struct {
	Message string `json:"message"`
}

func TestGetProfileByID(t *testing.T) {
	var jsonStr = []byte(`{"id": 1234567890}`)
	engine := Engine()

	var bearer = "Bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFsaWNlIiwicGFzc3dvcmQiOiIxMjM0NTYiLCJyb2xlIjoyLCJ1aWQiOjk5ODc2NTQzMjEsImlhdCI6MTY2MDI2Njc2MywiaXNzIjoiamlhd2VpIn0.NYlo2U4FWUpTVUgrysKakOy8Bq2pJdmDk5M22ZTRr7k"

	recorder := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/admin/get_profile", bytes.NewBuffer(jsonStr))
	request.Header.Add("Authorization", bearer)
	if err != nil {
		t.Fatalf("building request: %v", err)
	}
	engine.ServeHTTP(recorder, request)

	if recorder.Code != 200 {
		t.Fatalf("bad status code: %d", recorder.Code)
	}
	var response Response
	body := recorder.Body.String()
	if err != nil {
		t.Fatalf("reading response body: %v", err)
	}
	if err := json.Unmarshal([]byte(body), &response); err != nil {
		t.Fatalf("parsing json response: %v", err)
	}
}
