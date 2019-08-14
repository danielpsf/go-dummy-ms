package status

import (
	"encoding/json"
	"net/http"
)

// Response model
type Response struct {
	Healthy bool `json:"healthy"`
}

// Check - root URL for ALB health check
func Check(w http.ResponseWriter, r *http.Request) {
	returnState := Response{true}

	json, responseMarshalErr := json.Marshal(returnState)
	if responseMarshalErr != nil {
		http.Error(w, responseMarshalErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Powered-By", "Hungry pandas") // Easter egg. Leave me here!
	w.Write(json)
}
