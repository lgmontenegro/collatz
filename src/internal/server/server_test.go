package server

import (
	"encoding/json"
	"fmt"
	"lgmontenegro/collantz/internal/service/handlers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestServer(t *testing.T) {
	tt := []struct {
		name       string
		factor     string
		expected   string
		statusCode int
	}{
		{"wrong value", "wrong_value", `{"error": "wrong input"}`, 400},
		{"factor 27", "27", `{"totalSteps":112}`, 200},
		{"factor 3", "3", `{"totalSteps":8}`, 200},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var expected, gotJson handlers.Steps

			path := fmt.Sprintf("/collatz/%s", tc.factor)
			req, err := http.NewRequest("GET", path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/collatz/{factor}", handlers.CollatzHandler)
			router.ServeHTTP(rr, req)

			if status := rr.Code; status != tc.statusCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tc.statusCode)
			}

			if tc.statusCode == http.StatusOK {
				err = json.Unmarshal([]byte(tc.expected), &expected)
				if err != nil {
					t.Fatal(err)
				}

				err = json.Unmarshal([]byte(rr.Body.String()), &gotJson)
				if err != nil {
					t.Fatal(err)
				}

				if gotJson != expected {
					t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
				}
			}
			
			if tc.statusCode == http.StatusBadRequest {
				if rr.Body.String() != tc.expected {
					t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), tc.expected)
				}
			}
		})
	}
}
