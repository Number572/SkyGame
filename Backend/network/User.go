package network

import (
	"net/http"
	"encoding/json"
	"../models"
	"../settings"
)

type UserDataApi struct {
	VKID string
}

type UserCreateApi struct {
	VKID string
	Class int
	Clan string
}

func handleUserAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != "POST" {
		json.NewEncoder(w).Encode(ReturnCode{Code: -1})
		return 
	}

	rows, err := settings.DataBase.Query("SELECT JSON FROM User")
	if err != nil {
		json.NewEncoder(w).Encode(ReturnCode{Code: 1})
		return
	}

	var all_users []models.User
	var curr_user_json string
	var curr_user models.User

	for rows.Next() {
		rows.Scan(&curr_user_json)
		err := json.Unmarshal([]byte(curr_user_json), &curr_user)
		if err != nil {
			json.NewEncoder(w).Encode(ReturnCode{Code: 2})
			return
		}
		all_users = append(all_users, curr_user)
	}

	json.NewEncoder(w).Encode(struct{ AllUsers []models.User }{ AllUsers: all_users })
}

func handleUserData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != "POST" {
		json.NewEncoder(w).Encode(ReturnCode{Code: -1})
		return 
	}

	var data UserDataApi
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		json.NewEncoder(w).Encode(ReturnCode{Code: 1})
		return
	}

	var curr_user_json string
	var curr_user models.User
	
	row := settings.DataBase.QueryRow("SELECT JSON FROM User WHERE VKID = $1", data.VKID)
	row.Scan(&curr_user_json)

	err = json.Unmarshal([]byte(curr_user_json), &curr_user)
	if err != nil {
		json.NewEncoder(w).Encode(ReturnCode{Code: 2})
		return
	}

	json.NewEncoder(w).Encode(curr_user)
}

func handleUserCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != "POST" {
		json.NewEncoder(w).Encode(ReturnCode{Code: -1})
		return 
	}

	var data UserCreateApi
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		json.NewEncoder(w).Encode(ReturnCode{Code: 1})
		return
	}

	new_user := models.User {
		VKID: data.VKID,
		Class: data.Class,
		Clan: data.Clan,
		XP: 0,
		HP: 100,
		LVL: 1,
		Damage: 25,
	}

	new_user_json, err := json.MarshalIndent(new_user, "", "\t")
	if err != nil {
		json.NewEncoder(w).Encode(ReturnCode{Code: 2})
		return
	}

	_, err = settings.DataBase.Exec(
		"INSERT INTO User (VKID, Class, Clan, JSON) VALUES ($1, $2, $3, $4)",
		new_user.VKID,
		new_user.Class,
		new_user.Clan,
		string(new_user_json),
	)
	if err != nil {
		json.NewEncoder(w).Encode(ReturnCode{Code: 3})
		return
	}

	json.NewEncoder(w).Encode(ReturnCode{Code: 0})
}
