package network

import (
	"net/http"
	"encoding/json"
	"../settings"
)

type IndexApi struct {
	Head string
	Body string
}

type ReturnCode struct {
	Code int
}

func ServerHTTP() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/api", handleApi)
	http.HandleFunc("/api/battle/close", handleBattleClose)
	http.HandleFunc("/api/battle/data", handleBattleData)
	http.HandleFunc("/api/battle/update", handleBattleUpdate)
	http.HandleFunc("/api/battle/create", handleBattleCreate)
	http.HandleFunc("/api/chat/data", handleChatData)
	http.HandleFunc("/api/chat/create", handleChatCreate)
	http.HandleFunc("/api/task/get", handleTaskGet)
	http.HandleFunc("/api/task/check", handleTaskCheck)
	http.HandleFunc("/api/input/get", handleInputGet)
	http.HandleFunc("/api/input/check", handleInputCheck)
	http.HandleFunc("/api/user/all", handleUserAll)
	http.HandleFunc("/api/user/data", handleUserData)
	http.HandleFunc("/api/user/create", handleUserCreate)
	if err := http.ListenAndServe("0.0.0.0" + settings.Server.Address.Port, nil); err != nil {
		panic("Server is not running")
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(IndexApi{Head: "TEST", Body: "TEST"})
}

func handleApi(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(IndexApi{Head: "TEST", Body: "TEST"})
}
