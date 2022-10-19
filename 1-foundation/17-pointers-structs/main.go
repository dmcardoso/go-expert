package main

import "fmt"

type Client struct {
	Name string
}
type Account struct {
	Balance float64
}

/**
Unique instance
Like a constructor
*/
func NewAccount() *Account {
	return &Account{Balance: 0}
}

func (a Account) simulate(value float64) float64 {
	a.Balance += value

	return a.Balance
}

func (c *Client) Walked() {
	c.Name = "Daniel Moreira"
	fmt.Printf("O cliente %v andou\n", c.Name)
}

func main() {
	/** Pointers and structs
	 */

	daniel := Client{
		Name: "Daniel",
	}

	daniel.Walked()
	fmt.Printf("O nome da struct é: %v\n", daniel.Name)

	account := Account{Balance: 100}

	account.simulate(200)
	fmt.Printf("O saldo da conta é: %v\n", account.Balance)

	newAccount := NewAccount()

	newAccount.simulate(500)
	newAccount.Balance = 200
	fmt.Printf("O saldo da nova conta é: %v\n", newAccount.Balance)

	newAccount2 := NewAccount()
	fmt.Printf("O saldo da 2ª nova conta é: %v\n", newAccount2.Balance)
}
