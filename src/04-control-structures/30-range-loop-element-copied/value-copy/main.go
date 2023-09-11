package main

import (
	"fmt"
	"strings"
)

type account struct {
	balance float32
}

func main() {
	accounts := createAccounts()
	for _, a := range accounts {
		a.balance += 1000
	}
	fmt.Println(accounts)

	accounts = createAccounts()
	for i := range accounts {
		accounts[i].balance += 1000
	}
	fmt.Println(accounts)

	accounts = createAccounts()
	for i := 0; i < len(accounts); i++ {
		accounts[i].balance += 1000
	}
	fmt.Println(accounts)

	accountsPtr := createAccountsPtr()
	for _, a := range accountsPtr {
		a.balance += 1000
	}
	printAccountsPtr(accountsPtr)
}

func createAccounts() []account {
	return []account{
		{balance: 100.},
		{balance: 200.},
		{balance: 300.},
	}
}

func createAccountsPtr() []*account {
	return []*account{
		{balance: 100.},
		{balance: 200.},
		{balance: 300.},
	}
}

func printAccountsPtr(accounts []*account) {
	sb := strings.Builder{}
	sb.WriteString("[")
	s := make([]string, len(accounts))
	for i, account := range accounts {
		s[i] = fmt.Sprintf("{%.0f}", account.balance)
	}
	sb.WriteString(strings.Join(s, " "))
	sb.WriteString("]")
	fmt.Println(sb.String())
}
