package api

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestHealthHandler(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println(os.Getwd())
		log.Fatal("error loading .env file", err)
	}
	t.Run("GET /api/health", func(t *testing.T) {
		// Create a new request
		req, err := http.NewRequest("GET", "/api/health", nil)
		if err != nil {
			t.Fatal(err)
		}

		// Create a new response recorder
		w := httptest.NewRecorder()

		// Call the health handler
		HealthHandler(w, req)

		// Check the status code
		if status := w.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		// Check the response body
		got := w.Body.String()
		want := `{"message":"Healthy!"}`
		if got != want {
			t.Errorf("handler returned unexpected body: got %v want %v", got, want)
		}
	})
}
