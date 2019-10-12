package escrow

import "github.com/stellar/go/keypair"

func GenerateAccount() *keypair.Full {
	full, _ := keypair.Random()
	return full
}
