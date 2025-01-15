package services

import (
	"fmt"

	"github.com/jdplumst/dota2-stats-tracker/server/database"
)

func GetHeroes() ([]database.Hero, error) {
	db, err := database.NewConnection()
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}
	defer db.Close()

	var heroes []database.Hero
	err = db.Select(&heroes, "SELECT * FROM hero")
	if err != nil {
		return nil, fmt.Errorf("error getting heroes: %w", err)
	}

	return heroes, nil
}
