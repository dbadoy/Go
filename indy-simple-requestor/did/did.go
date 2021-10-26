package did

import (
	"github.com/hyperledger/indy-vdr/wrappers/golang/identifiers"
	"github.com/hyperledger/indy-vdr/wrappers/golang/crypto"
	"crypto/ed25519"
	"crypto/rand"

	"fmt"
	"strings"
)

func CreateDidWithSeed(seed string) ( sig *crypto.Ed25519Signer, sDid string, err error) {

	base, err := identifiers.ConvertSeed(seed[0:32])
	if err != nil {
		return 
	}

	var pubkey ed25519.PublicKey
	var privkey ed25519.PrivateKey

	privkey = ed25519.NewKeyFromSeed(base)
	pubkey = privkey.Public().(ed25519.PublicKey)

	did, err := identifiers.CreateDID(&identifiers.MyDIDInfo{PublicKey: pubkey, Cid: true, MethodName: "sov"})
	if err != nil {
		return
	}

	res := strings.Split(did.String(), ":")

	fmt.Println("Did: ", res[2])
	fmt.Println("Verkey: ", did.AbbreviateVerkey())

	mysig := crypto.NewSigner(pubkey, privkey)

	return mysig, res[2], nil
}

func CreateRandomDid() (did string, verkey string, err error) {
	
	someRandomPubkey, _, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return
	}

	someRandomDID, err := identifiers.CreateDID(&identifiers.MyDIDInfo{PublicKey: someRandomPubkey, MethodName: "sov", Cid: true})
	if err != nil {
		return
	}

	return someRandomDID.DIDVal.MethodSpecificID, someRandomDID.Verkey, nil
}
