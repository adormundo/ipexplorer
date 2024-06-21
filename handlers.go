package main

import (
	"encoding/json"
	"net/http"
)

// handleIPRequest handles the request to fetch and return IP details.
func handleIPRequest(w http.ResponseWriter, r *http.Request) {
	clientIP := getClientIP(r)
	ip, err := fetchIPFromAPI(clientIP)
	if err != nil {
		http.Error(w, "Error fetching IP from api.ipify.org", http.StatusInternalServerError)
		return
	}
	ipDetails, err := fetchIPDetails(ip)
	if err != nil {
		http.Error(w, "Error fetching details from ip-api.com", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ipDetails)
}