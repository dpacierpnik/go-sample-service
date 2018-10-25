package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const githubZenEndpoint = "https://api.github.com/zen"

type zenResponseDto struct {
	Quotation string `json:"quotation"`
}

func Zen(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get(githubZenEndpoint)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to call '%s'. Root cause: %s", githubZenEndpoint, err.Error()), http.StatusInternalServerError)
		return
	}

	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("Github endpoint '%s' respond with an error: %d.", githubZenEndpoint, resp.StatusCode), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()
	respPayload, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to read response from '%s'. Root cause: %s", githubZenEndpoint, err.Error()), http.StatusInternalServerError)
		return
	}

	dto := &zenResponseDto{
		Quotation: string(respPayload),
	}
	dtoJsonBytes, err := json.Marshal(dto)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to marshall to JSON: '%v'. Root cause: %s", dto, err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(dtoJsonBytes)
}
