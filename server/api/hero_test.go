package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jdplumst/dota2-stats-tracker/server/database"
	"github.com/joho/godotenv"
)

func TestHeroHandler(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal("error loading .env file", err)
	}
	database.SeedDatabase()

	t.Run("GET /api/hero", func(t *testing.T) {
		// Create a new request
		req, err := http.NewRequest("GET", "/api/hero", nil)
		if err != nil {
			t.Fatal(err)
		}

		// Create a new response recorder
		w := httptest.NewRecorder()

		// Call the hero handler
		HeroHandler(w, req)

		// Check the status code
		if w.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, w.Code)
		}

		// Check the response body
		jsonResp := make(map[string][]database.Hero)
		err = json.Unmarshal(w.Body.Bytes(), &jsonResp)
		if err != nil {
			t.Errorf("error unmarshalling response body: %v", err)
		}

		got := len(jsonResp["heroes"])
		want := 126
		if got != want {
			t.Errorf("handler returned unexpected body, got length %v want length %v", got, want)
		}
	})
}
