package services

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestHeroService(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println(os.Getwd())
		log.Fatal("error loading .env file", err)
	}
	t.Run("GetHeroes", func(t *testing.T) {
		got, err := GetHeroes()
		if err != nil {
			t.Errorf("service returned an error: %v", err)
		}

		want := 126

		if len(got) != want {
			t.Errorf("service returned unexpected data: got length %v\n want length %v", len(got), want)
		}
	})
}
