package main

import (
	"github.com/dmcardoso/go-expert/7-apis/3-user/configs"
)

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
