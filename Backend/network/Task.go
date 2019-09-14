package network

import (
	"net/http"
	"math/rand"
	"encoding/json"
)

type TaskCheckApi struct {
	VKID string
	Task string
	Check string
}

var Tasks = map[string][]string {
	"hat": []string{"шляпа", "шапка"},
	"river": []string{"река", "поток"},
	"lucky": []string{"везучий", "удачливый", "удачный"},
	"communication": []string{"коммуникация", "общение", "сообщение"},
	"reward": []string{"награда", "вознаграждение"},
	"development": []string{"развитие", "разработка", "создание"},
	"application": []string{"приложение"},
	"nothing": []string{"ничего"},
}

func RandTask(tasks map[string][]string) string {
	var length = len(tasks)
	if length == 0 {
		return "nothing"
	}
	var randnum = rand.Int() % length;
	for task := range tasks {
		if randnum == 0 {
			return task
		}
		randnum--
	}
	return "nothing"
}

func handleTaskGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != "POST" {
		json.NewEncoder(w).Encode(ReturnCode{Code: -1})
		return 
	}

	json.NewEncoder(w).Encode(struct{ Task string }{ RandTask(Tasks) })
}

func handleTaskCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != "POST" {
		json.NewEncoder(w).Encode(ReturnCode{Code: -1})
		return 
	}

	var data TaskCheckApi
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		json.NewEncoder(w).Encode(ReturnCode{Code: 1})
		return
	}
	for _, translate := range Tasks[data.Task] {
		if translate == data.Check {
			xp := rand.Int() % 125 + 1
			upgradeUser(data.VKID, xp)
			json.NewEncoder(w).Encode(struct{ 
				Code int 
				XP int 
			}{
				Code: 0,
				XP: xp,
			})
			return
		}
	}
	json.NewEncoder(w).Encode(ReturnCode{Code: 2})
}
