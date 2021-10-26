package connection

import ( 
	"io"
	"os"
)

const (
	ENDORSERSEED1 = "CHANGEME"
	ENDORSERSEED2 = "CHANGEME"
	ENDORSERSEED3 = "CHANGEME"
)

func GetGenesisFile(pool string) (genesisFile io.ReadCloser, endorserSeed string, err error) {
	switch pool {
	case "pool1":
		genesisFile, err = os.Open("./connection/pool_transactions_genesis1")
		endorserSeed = ENDORSERSEED1
	case "pool2":
		genesisFile, err = os.Open("./connection/pool_transactions_genesis2")
		endorserSeed = ENDORSERSEED2
	case "pool3":
		genesisFile, err = os.Open("./connection/pool_transactions_genesis3")
		endorserSeed = ENDORSERSEED3
	default:
		return
	}

	if err != err {
		return
	}

	return
}