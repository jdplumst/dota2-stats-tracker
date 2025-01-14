package api

import (
	"encoding/json"
	"net/http"

	"github.com/jdplumst/dota2-stats-tracker/server/db"
	"github.com/jdplumst/dota2-stats-tracker/server/utils"
)

func HealthHandler(w http.ResponseWriter, req *http.Request) {
	utils.SetCors(w)
	w.Header().Set("Content-Type", "application/json")

	_, err := db.NewConnection()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := make(map[string]string)
		resp["message"] = "Error connecting to the database"
		jsonResp, _ := json.Marshal(resp)
		w.Write(jsonResp)
		return
	}

	w.WriteHeader(http.StatusOK)
	resp := make(map[string]string)
	resp["message"] = "Healthy!"
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
