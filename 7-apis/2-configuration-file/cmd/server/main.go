package main

import (
	"github.com/dmcardoso/go-expert/7-apis/2-configuration-file/configs"
)

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
