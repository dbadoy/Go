package indydidchecker

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/hyperledger/indy-vdr/wrappers/golang/vdr"
)

type ReadReply struct {
	Type          string      `json:"type"`
	Identifier    string      `json:"identifier,omitempty"`
	ReqID         uint32      `json:"reqId"`
	SeqNo         uint32      `json:"seqNo"`
	TxnTime       uint32      `json:"TxnTime"`
	Data          interface{} `json:"data"`
	SignatureType string      `json:"signature_type,omitempty"`
	Origin        string      `json:"origin,omitempty"`
	Dest          string      `json:"dest,omitempty"`
	Ref           uint32      `json:"ref,omitempty"`
	Tag           string      `json:"tag,omitempty"`
}

func main() {
	genesisFile, err := getGenesisFile()
	if err != nil {
		fmt.Print(err)
		return
	}

	defer genesisFile.Close()

	client, err := vdr.New(genesisFile)
	if err != nil {
		fmt.Print(err)
		return
	}

	err = client.RefreshPool()
	if err != nil {
		fmt.Print(err)
		return
	}

	rply, err := client.GetNym("4zRGZz6aGepQvqSbeguWp9")
	if err != nil {
		fmt.Print(err)
		return
	}

	toByte, _ := json.Marshal(rply)

	var response ReadReply
	err = json.Unmarshal(toByte, &response)

	if response.Data == nil {
		fmt.Print("non exists did")
		return
	}

	fmt.Println("Exist!")
	fmt.Println(response.Data)
}

func getGenesisFile() (genesisFile io.ReadCloser, err error) {
	genesisFile, err = os.Open("./pool_transactions_genesis")
	if err != err {
		return
	}
	return
}
