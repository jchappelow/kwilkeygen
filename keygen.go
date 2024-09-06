package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/kwilteam/kwil-db/core/crypto"
	"github.com/kwilteam/kwil-db/core/crypto/auth"
)

func main() {
	priv, err := crypto.GenerateSecp256k1Key()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	signer := &auth.EthPersonalSigner{Key: *priv}
	acctID := signer.Identity()
	fmt.Println("Key:", priv.Hex())
	fmt.Printf("Pub: %x\n", priv.PubKey().Bytes())
	fmt.Println("Identity:", hex.EncodeToString(acctID))
}
