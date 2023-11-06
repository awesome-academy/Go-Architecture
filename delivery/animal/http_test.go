package animal

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAnimalHandler_Handler(t *testing.T) {
	testcases := []struct {
		method             string
		expectedStatusCode int
	}{
		{"GET", http.StatusOK},
		{"POST", http.StatusOK},
		{"DELETE", http.StatusMethodNotAllowed},
	}

	for _, v := range testcases {
		req := httptest.NewRequest(v.method, "/animal", nil)
		w := httptest.NewRecorder()

		a := New(mockDatastore{})
		a.Handler(w, req)

		if w.Code != v.expectedStatusCode {
			t.Errorf("Expected %v\tGot %v", v.expectedStatusCode, w.Code)
		}
	}

}
