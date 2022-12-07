package main

import (
	"github.com/dmcardoso/go-expert/7-apis/4-product/configs"
)

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
