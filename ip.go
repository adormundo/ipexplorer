package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
)

// IPResponse represents the structure of the response from the IP API.
type IPResponse struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Latitude    float64 `json:"lat"`
	Longitude   float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	ISP         string  `json:"isp"`
	Org         string  `json:"org"`
	AS          string  `json:"as"`
	IP          string  `json:"ip"` // Field for the IP found by https://api.ipify.org
}

// getLocalIP retrieves the local IP address of the machine.
func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "Error obtaining IP address"
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			return ipnet.IP.String()
		}
	}
	return "IP address not found"
}

// getClientIP retrieves the client's IP address from the request.
func getClientIP(r *http.Request) string {
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = r.Header.Get("X-Forwarded-For")
	}
	if ip == "" {
		ip = r.RemoteAddr
	}
	ip = strings.Split(ip, ":")[0]
	return ip
}

// fetchIPFromAPI fetches the IP address from the ipify API.
func fetchIPFromAPI(clientIP string) (string, error) {
	ipifyURL := fmt.Sprintf("https://api.ipify.org?format=json&ip=%s", clientIP)
	resp, err := http.Get(ipifyURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var ipResponse struct {
		IP string `json:"ip"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&ipResponse); err != nil {
		return "", err
	}
	return ipResponse.IP, nil
}

// fetchIPDetails fetches detailed IP information from the ip-api.
func fetchIPDetails(ip string) (IPResponse, error) {
	ipAPIURL := fmt.Sprintf("http://ip-api.com/json/%s", ip)
	resp, err := http.Get(ipAPIURL)
	if err != nil {
		return IPResponse{}, err
	}
	defer resp.Body.Close()

	var ipDetails IPResponse
	if err := json.NewDecoder(resp.Body).Decode(&ipDetails); err != nil {
		return IPResponse{}, err
	}
	ipDetails.IP = ip
	return ipDetails, nil
}
