package echo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Endpoints struct {
	Headers http.HandlerFunc
	Body    http.HandlerFunc
}

func NewEndpoints() *Endpoints {
	return &Endpoints{
		Headers: echoHeaders,
		Body:    echoBody,
	}
}

func echoBody(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.WriteHeader(http.StatusOK)
	w.Write(bodyBytes)
}

func echoHeaders(w http.ResponseWriter, r *http.Request) {

	jsonPayload, err := json.Marshal(r.Header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonPayload)
}
