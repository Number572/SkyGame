package network

import (
	"net/http"
	"math/rand"
	"encoding/json"
)

type InputCheckApi struct {
	VKID string
	Check []string
}

var Input = [][]string {
	[]string{"went", "I _ (go) to Italy three years ago."},
	[]string{"drunk", "That is the best drink I have ever _ (drink)."},
	[]string{"think", "I _ (think) of the best idea."},
	[]string{"sold", "My grandmother _ (sell) chocolate when she was young."},
	[]string{"rode", "In Egypt, I _ (ride) camels."},
	[]string{"grew", "My tree _ (grow) so fast that we had to trim it."},
	[]string{"bought", "I _ (buy) lots of DVDs last weekend."},
	[]string{"shot", "A man _ (shoot) at the man, but he missed him."},
	[]string{"fell", "My brother _ (fall) down the stairs and cracked his head open."},
	[]string{"wore", "I _ (wear) my best clothes yesterday."},
}

func handleInputGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != "POST" {
		json.NewEncoder(w).Encode(ReturnCode{Code: -1})
		return 
	}

	json.NewEncoder(w).Encode(struct{ Inputs []string }{ Inputs: GetInputs(Input) })
}

func GetInputs(input [][]string) []string {
	var list []string

	for _, value := range input {
		list = append(list, value[1])
	}

	return list
}

func handleInputCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != "POST" {
		json.NewEncoder(w).Encode(ReturnCode{Code: -1})
		return 
	}

	var data InputCheckApi
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		json.NewEncoder(w).Encode(ReturnCode{Code: 1})
		return
	}

	var count = 0
	var index = 0
	var length = len(data.Check)

	for _, value := range Input {
		if index == length {
			break
		}
		if value[0] == data.Check[index] {
			count++
		}
		index++
	}

	if count >= 7 {
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

	json.NewEncoder(w).Encode(ReturnCode{Code: 2})
}
