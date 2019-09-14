package main

import (
	"os"
	"fmt"
	"time"
	"math/rand"
	"./network"
	"./settings"
)

func main() {
	fmt.Println("[Server is running ...]")
	rand.Seed(time.Now().UnixNano())
	settings.OpenDataBase("database.db")
	settings.InitArgs(os.Args)
	network.ServerHTTP()
}
