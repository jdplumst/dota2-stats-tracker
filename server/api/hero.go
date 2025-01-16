package api

import (
	"encoding/json"
	"net/http"

	"github.com/jdplumst/dota2-stats-tracker/server/database"
	"github.com/jdplumst/dota2-stats-tracker/server/services"
)

func HeroHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.URL.Query())
	heroes, err := services.GetHeroes()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := make(map[string]string)
		resp["message"] = "Error getting heroes"
		jsonResp, _ := json.Marshal(resp)
		w.Write(jsonResp)
		return
	}

	w.WriteHeader(http.StatusOK)
	resp := make(map[string][]database.Hero)
	resp["heroes"] = heroes
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
