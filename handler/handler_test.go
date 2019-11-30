package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fmusayev/golang-excel/types"
	"github.com/gorilla/mux"
)

func TestHealthCheckHandler(t *testing.T) {
	router := mux.NewRouter()
	h := NewHandler(router)

	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(h.Health)

	handler.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestReadinessCheckHandler(t *testing.T) {
	router := mux.NewRouter()
	h := NewHandler(router)

	req, err := http.NewRequest("GET", "/readiness", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(h.Health)

	handler.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestDownloadExcelHandler(t *testing.T) {
	router := mux.NewRouter()
	h := NewHandler(router)

	b := []types.ExcelForm{
		types.ExcelForm{
			ID:           "01",
			FirstName:    "FUAD",
			LastName:     "MUSAYEV",
			Username:     "fuadmusayev",
			Email:        "fuadmusayev_@outlook.com",
			AddressLine:  "Azerbaijan, Baku, Unknown st. 15 ",
			ContactPhone: "+99450*******",
		},
	}

	json, err := json.Marshal(b)
	if err != nil {
		println(err)
		t.Fail()
	}

	req, err := http.NewRequest("POST", "/v1/generate-excel", bytes.NewBuffer(json))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(h.Health)

	handler.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
