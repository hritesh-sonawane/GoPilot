package wallet

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)

	got := wallet.Balance()
	fmt.Printf("Address of balance in test is %v \n", &wallet.balance)
	want := 10

	if got != want {
		t.Errorf("Got: %d | Want: %d", got, want)
	}
}
