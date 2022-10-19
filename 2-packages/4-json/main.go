package main

import (
	"encoding/json"
	"os"
)

type Account struct {
	Number  int `json:"number"`
	Balance int `json:"balance"`
	Hidden  int `json:"-"`
}

func main() {
	/* Json

	 */

	account := Account{Number: 1, Balance: 200, Hidden: 5}

	res, err := json.Marshal(account)

	if err != nil {
		println(err)
	}

	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(account)

	if err != nil {
		println(err)
	}

	// unmarshal json

	jsonPure := []byte(`{"Number":2,"Balance":300}`)
	var accountX Account
	err = json.Unmarshal(jsonPure, &accountX)

	if err != nil {
		println(err)
	}

	err = json.NewEncoder(os.Stdout).Encode(accountX)

	if err != nil {
		println(err)
	}

	// tags format

	jsonPure = []byte(`{"number":2,"balance":300}`)
	var accountY Account
	err = json.Unmarshal(jsonPure, &accountY)

	if err != nil {
		println(err)
	}

	err = json.NewEncoder(os.Stdout).Encode(accountY)

	if err != nil {
		println(err)
	}
}
