package settings

import (
	"../utils"
	"database/sql"
	_"github.com/mattn/go-sqlite3"
)

func OpenDataBase(name string) {
	if !utils.FileIsExist(name) {
		utils.CreateFile(name)
	}

	var err error
	DataBase, err = sql.Open("sqlite3", name)
	if err != nil {
		panic("DataBase not running")
	}

	_, err = DataBase.Exec(`
CREATE TABLE IF NOT EXISTS Task (
	Id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
	Class VARCHAR(128),
	Text VARCHAR(128),
	Answer VARCHAR(128),
	XP INTEGER
);

CREATE TABLE IF NOT EXISTS User (
	Id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
	VKID VARCHAR(128) UNIQUE,
	Class VARCHAR(128),
	Clan VARCHAR(128),
	JSON VARCHAR(1024)
);
`)
	if err != nil {
		panic("DataBase not writeable")
	}
}
