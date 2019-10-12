package escrow

import (
	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/network"
	"github.com/stellar/go/txnbuild"
	"github.com/tevino/log"
	"github.com/vlyl/stellar-escrow-account/config"
	"github.com/vlyl/stellar-escrow-account/util"
	"strconv"
)

func NewAccount() *keypair.Full {
	full, _ := keypair.Random()
	ca := txnbuild.CreateAccount{
		Destination:   full.Address(),
		Amount:        "10",
		SourceAccount: nil,
	}

	tx := txnbuild.Transaction{
		SourceAccount: GetSimpleAccount(config.FaucetAccount.ID),
		Operations:    []txnbuild.Operation{&ca},
		Timebounds:    txnbuild.NewInfiniteTimeout(),
		Network:       network.TestNetworkPassphrase,
	}

	err := tx.Build()
	if err != nil {
		log.Error(err)
		return nil
	}

	// write keypair to file
	err = util.AppendToFile("escrow.txt", full.Address()+"/"+full.Seed())
	if err != nil {
		log.Error(err)
		return nil
	}

	return full
}

func GetSimpleAccount(id string) *txnbuild.SimpleAccount {
	ar := horizonclient.AccountRequest{
		AccountID: id,
		DataKey:   "",
	}

	detail, err := horizonclient.DefaultTestNetClient.AccountDetail(ar)
	if err != nil {
		log.Error(err)
		return nil
	}

	return &txnbuild.SimpleAccount{
		AccountID: id,
		Sequence: func() int64 {
			seq, _ := strconv.ParseInt(detail.Sequence, 10, 64)
			return int64(seq)
		}(),
	}
}

// setup new account for escrow
