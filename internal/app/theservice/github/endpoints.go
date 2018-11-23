package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Endpoints struct {
	Zen http.HandlerFunc
}

func NewEndpoints(client *Client) *Endpoints {
	return &Endpoints{
		Zen: newZenEndpoint(client),
	}
}

type zenResponseDto struct {
	Zen Zen `json:"zen"`
}

func newZenEndpoint(client *Client) http.HandlerFunc {

	return func (w http.ResponseWriter, r * http.Request) {

		zen, err := client.GetZen()
		if err != nil {
			http.Error(w, fmt.Sprintf("Backend service error. Root cause: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		dto := &zenResponseDto{
			Zen: zen,
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
}

