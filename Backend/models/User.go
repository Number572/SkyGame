package models

type User struct {
	VKID string
	Class int // 0 - warrior, 1 - healer;
	Clan string 
	XP int
	HP int
	LVL int
	Damage int
}
