package settings

import (
	"database/sql"
	"../models"
)

var (
	Server models.Server
	DataBase *sql.DB
)
