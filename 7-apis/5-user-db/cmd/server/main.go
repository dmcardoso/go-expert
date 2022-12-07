package main

import (
	"github.com/dmcardoso/go-expert/7-apis/5-user-db/configs"
)

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
