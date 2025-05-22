package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

type IPMIRequest struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Raw      string `json:"raw"`
}

type IPMIResponse struct {
	Output string `json:"output"`
	Error  string `json:"error"`
}

func handleIPMIRaw(w http.ResponseWriter, r *http.Request) {
	var req IPMIRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	args := []string{
		"-I", "lanplus",
		"-H", req.Host,
		"-U", req.Username,
		"-P", req.Password,
		"raw",
	}

	// Parse raw string into individual arguments
	args = append(args, splitRawCommand(req.Raw)...)

	fmt.Println(args)

	cmd := exec.Command("ipmitool", args...)
	out, err := cmd.CombinedOutput()

	resp := IPMIResponse{
		Output: string(out),
	}

	if err != nil {
		resp.Error = err.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func splitRawCommand(raw string) []string {
	return append([]string{}, splitAndClean(raw)...)
}

func splitAndClean(raw string) []string {
	// You could sanitize or validate better here
	return filterEmpty(splitBySpace(raw))
}

func splitBySpace(s string) []string {
	return strings.Fields(s)
}

func filterEmpty(input []string) []string {
	var out []string
	for _, i := range input {
		if i != "" {
			out = append(out, i)
		}
	}
	return out
}

func main() {
	http.HandleFunc("/ipmi/raw", handleIPMIRaw)
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
