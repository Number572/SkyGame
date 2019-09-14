package network

import (
	"net/http"
	"encoding/json"
)

var Chat []string

type ChatCreateApi struct {
	Message string
}

func handleChatCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != "POST" {
		json.NewEncoder(w).Encode(ReturnCode{Code: -1})
		return 
	}

	var data ChatCreateApi
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		json.NewEncoder(w).Encode(ReturnCode{Code: 1})
		return
	}
	Chat = append(Chat, data.Message)
	json.NewEncoder(w).Encode(ReturnCode{Code: 0})
}

func handleChatData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != "POST" {
		json.NewEncoder(w).Encode(ReturnCode{Code: -1})
		return 
	}

	json.NewEncoder(w).Encode(struct{ Chat []string }{ Chat })
}
