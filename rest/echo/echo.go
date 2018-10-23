package echo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Body(w http.ResponseWriter, r *http.Request) {

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

func Headers(w http.ResponseWriter, r *http.Request) {

	jsonPayload, err := json.Marshal(r.Header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonPayload)
}
