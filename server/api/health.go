package api

import (
	"encoding/json"
	"net/http"

	utils "github.com/jdplumst/dota2-stats-tracker/server/utils"
)

func HealthHandler(w http.ResponseWriter, req *http.Request) {
	utils.SetCors(w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := make(map[string]string)
	resp["message"] = "Healthy!"
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
