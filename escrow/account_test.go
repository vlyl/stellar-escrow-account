package escrow

import (
	"github.com/vlyl/stellar-escrow-account/config"
	"testing"
)

func TestNewAccount(t *testing.T) {
	config.LoadConfig("")
	full := NewAccount()
	t.Log(full)
}
