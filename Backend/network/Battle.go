package network

import (
	"time"
	"net/http"
	"math/rand"
	"encoding/hex"
	"encoding/json"
	"../utils"
	"../models"
	"../settings"
)

type BattlePlace struct {
	XP int
	Heroes []models.User
	Enemies []models.User
}
var Battles = make(map[string]BattlePlace)

type BattleCloseApi BattleCreateApi
type BattleDataApi BattleCreateApi
type BattleCreateApi struct {
	Heroes []string
}

type BattleUpdateApi struct {
	Heroes []string
	Attacker string
	Defender string
}

func handleBattleData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != "POST" {
		json.NewEncoder(w).Encode(ReturnCode{Code: -1})
		return 
	}

	var data BattleDataApi
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		json.NewEncoder(w).Encode(ReturnCode{Code: 1})
		return
	}

	Heroes := GetHeroes(data)
	if Heroes == nil {
		json.NewEncoder(w).Encode(ReturnCode{Code: 2})
		return
	}

	hash_sum := HashSumBattle(BattlePlace{Heroes: Heroes})
	json.NewEncoder(w).Encode(Battles[hash_sum])
}

func GetHeroes(data BattleDataApi) []models.User {
	var Heroes []models.User
	var curr_user models.User
	var json_user string

	for index := range data.Heroes {
		row := settings.DataBase.QueryRow("SELECT JSON FROM User WHERE VKID = $1", data.Heroes[index])
		row.Scan(&json_user)

		err := json.Unmarshal([]byte(json_user), &curr_user)
		if err != nil {
			return nil
		}

		Heroes = append(Heroes, curr_user)
	}

	return Heroes
}

func handleBattleClose(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != "POST" {
		json.NewEncoder(w).Encode(ReturnCode{Code: -1})
		return 
	}

	var data BattleDataApi
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		json.NewEncoder(w).Encode(ReturnCode{Code: 1})
		return
	}

	Heroes := GetHeroes(data)
	hash_sum := HashSumBattle(BattlePlace{Heroes: Heroes})
	new_data := Battles[hash_sum]

	health := 0
	for _, enemie := range new_data.Enemies {
		health += enemie.HP
	}

	if len(new_data.Heroes) == 0 {
		json.NewEncoder(w).Encode(ReturnCode{Code: 2})
		return
	}

	if health <= 0 {
		for _, hero := range new_data.Heroes {
			x := upgradeUser(hero.VKID, new_data.XP)
			if x != 0 {
				json.NewEncoder(w).Encode(ReturnCode{Code: x})
			}
		}
	}
	
	delete(Battles, hash_sum)
	json.NewEncoder(w).Encode(ReturnCode{Code: 0})
}

func upgradeUser(hero string, xp int) int {
	var curr_user_json string
	var curr_user models.User

	row := settings.DataBase.QueryRow("SELECT JSON FROM User WHERE VKID = $1", hero)
	row.Scan(&curr_user_json)

	err := json.Unmarshal([]byte(curr_user_json), &curr_user)
	if err != nil {
		return 2
	}

	curr_user.XP += xp
	for curr_user.XP >= 100 {
		curr_user.LVL++
		curr_user.XP -= 100

		// By Class
		if curr_user.Class == 0 && rand.Int() % 100 < 33 {
			curr_user.Damage += 5 // warrior
		}
		if curr_user.Class == 1 && rand.Int() % 100 < 33 {
			curr_user.HP += 5 // healer
		}

		// Default
		if rand.Int() % 100 < 10 {
			curr_user.Damage += 5
		}
		if rand.Int() % 100 < 10 {
			curr_user.HP += 5
		}
	}

	user_json, err := json.MarshalIndent(curr_user, "", "\t")
	if err != nil {
		return 3
	}
	
	_, err = settings.DataBase.Exec(
		"UPDATE User SET JSON = $1 WHERE VKID = $2", 
		string(user_json),
		curr_user.VKID,
	)
	if err != nil {
		return 4
	}

	return 0
}

func handleBattleUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != "POST" {
		json.NewEncoder(w).Encode(ReturnCode{Code: -1})
		return 
	}

	var data BattleUpdateApi
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		json.NewEncoder(w).Encode(ReturnCode{Code: 1})
		return
	}
	hash_sum := HashSumBattle(BattlePlace{Heroes: GetHeroes(BattleDataApi{data.Heroes})})
	
	attacker_is_hero, attacker_damage, attacker_class := FindAttacker(data.Attacker, Battles[hash_sum])
	if attacker_damage == -1 {
		json.NewEncoder(w).Encode(ReturnCode{Code: 2})
		return
	}

	defender_is_hero, defender_is_exist := FindDefender(data.Defender, Battles[hash_sum])
	if !defender_is_exist {
		json.NewEncoder(w).Encode(ReturnCode{Code: 3})
		return
	}

	// Heal hero
	if attacker_is_hero && defender_is_hero && attacker_class == 1 {
		for index, hero := range Battles[hash_sum].Heroes {
			if hero.VKID == data.Defender {
				Battles[hash_sum].Heroes[index].HP += attacker_damage
			}
		}
	}

	// Heal enemy
	if !attacker_is_hero && !defender_is_hero && attacker_class == 1 {
		for index, enemy := range Battles[hash_sum].Enemies {
			if enemy.VKID == data.Defender {
				Battles[hash_sum].Enemies[index].HP += attacker_damage
			}
		}
	}

	// Damage enemy
	if attacker_is_hero && !defender_is_hero {
		for index, enemy := range Battles[hash_sum].Enemies {
			if enemy.VKID == data.Defender {
				Battles[hash_sum].Enemies[index].HP -= attacker_damage
			}
		}
	}

	// Damage hero
	if !attacker_is_hero && defender_is_hero {
		for index, hero := range Battles[hash_sum].Heroes {
			if hero.VKID == data.Defender {
				Battles[hash_sum].Heroes[index].HP -= attacker_damage
			}
		}
	}

	json.NewEncoder(w).Encode(ReturnCode{Code: 0})
}

func FindAttacker(attacker string, place BattlePlace) (bool, int, int) {
	for _, hero := range place.Heroes {
		if hero.VKID == attacker {
			return true, hero.Damage, hero.Class
		}
	}
	for _, enemy := range place.Enemies {
		if enemy.VKID == attacker {
			return false, enemy.Damage, enemy.Class
		}
	}
	return false, -1, -1
}

func FindDefender(defender string, place BattlePlace) (bool, bool) {
	for _, hero := range place.Heroes {
		if hero.VKID == defender {
			return true, true
		}
	}
	for _, enemy := range place.Enemies {
		if enemy.VKID == defender {
			return false, true
		}
	}
	return false, false
}

func handleBattleCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != "POST" {
		json.NewEncoder(w).Encode(ReturnCode{Code: -1})
		return 
	}

	var data BattleDataApi
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		json.NewEncoder(w).Encode(ReturnCode{Code: 1})
		return
	}

	var Heroes = GetHeroes(data)
	if len(Heroes) == 0 {
		json.NewEncoder(w).Encode(ReturnCode{Code: 2})
		return
	}

	var Enemies = GetEnemies(Heroes)

	battle := BattlePlace{rand.Int() % 125 + 1, Heroes, Enemies}
	hash_sum := HashSumBattle(battle)

	go func() {
		time.Sleep(30 * time.Minute)
		delete(Battles, hash_sum)
	}()

	Battles[hash_sum] = battle
	json.NewEncoder(w).Encode(battle)
}

func GetEnemies(Heroes []models.User) []models.User {
	return []models.User{ 
		NewEnemie(Heroes, "1"), 
		NewEnemie(Heroes, "2"), 
		NewEnemie(Heroes, "3"), 
	}
}

func NewEnemie(Heroes []models.User, num string) models.User {
	var avr_hp, avr_lvl, avr_dmg int
	var size = len(Heroes)

	for _, hero := range Heroes {
		avr_hp += hero.HP
		avr_lvl += hero.LVL 
		avr_dmg += hero.Damage
	}

	avr_hp /= size
	avr_lvl /= size
	avr_dmg /= size

	class := 0
	if rand.Int() % 100 < 50 {
		class = 1
	}
	return models.User {
		VKID: "Enemy" + num,
		Class: class,
		Clan: "ENEMIES",
		XP: rand.Int() % 100,
		HP: avr_hp,
		LVL: avr_lvl,
		Damage: avr_dmg,
	}
}

func HashSumBattle(battle BattlePlace) string {
	var summ string
	for _, hero := range battle.Heroes {
		summ += hero.VKID
	}
	return hex.EncodeToString(utils.HashSum([]byte(summ)))
}
