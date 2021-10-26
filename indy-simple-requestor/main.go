package main

import (
	"fmt"
	"errors"

	"connection"
	"message"
	"did"

	"github.com/hyperledger/indy-vdr/wrappers/golang/vdr"
        "github.com/hyperledger/indy-vdr/wrappers/golang/crypto"

	"github.com/gin-gonic/gin"
	"net/http"
)


func main() {
	r := gin.Default()

	pool1 := r.Group("/pool1")
        {
                client, endorserSeed, err := getClient("pool1")
                if err != nil {
                        return
                }

                defer client.Close()

                pool1.GET("/did/:did", func(c *gin.Context){
                        did := c.Param("did")
                        res, err := getNym(did, client)
                        if err != nil {
							message.Response(c, http.StatusBadRequest, err.Error())
                            return
                        }

						message.Response(c, http.StatusOK, res)
                })
                pool1.GET("/did/issueDid", func(c *gin.Context){
                        rdid, rverkey, err := did.CreateRandomDid()
                        if err != nil {
							message.Response(c, http.StatusBadRequest, err.Error())
                            return
                        }

                        sig, sDid, err := did.CreateDidWithSeed(endorserSeed)
                        if err != nil {
							message.Response(c, http.StatusBadRequest, err.Error())
                            return
                        }

                        err = Nym(rdid, rverkey, sDid, sig, client)
                        if err != nil {
							message.Response(c, http.StatusBadRequest, err.Error())
                            return
                        }

						message.Response(c, http.StatusOK, message.IssueDidRes{Did: rdid, Verkey: rverkey})
                })
                pool1.GET("/status", func(c *gin.Context){
                        PoolStatus, err := client.GetPoolStatus()
                        if err != nil {
							message.Response(c, http.StatusBadRequest, err.Error())
                            return
                        }

						message.Response(c, http.StatusOK, PoolStatus.Nodes)
                })
                pool1.GET("/genesis", func(c *gin.Context){
                        c.File("connection/dev_pool_genesis")
                })
        }

	r.Run()
}

func getClient(pool string) (client *vdr.Client, seed string, err error) {
	genesisFile, endorserSeed, err := connection.GetGenesisFile(pool)
	if err != nil {
			fmt.Println(err)
			return
	}

	defer  genesisFile.Close()

	client, err = vdr.New(genesisFile)
	if err != nil {
			fmt.Println(err)
			return
	}
	fmt.Println("success create client.")

	err = client.RefreshPool()
	if err != nil {
			fmt.Println(err)
			return
	}

	PoolStatus, err := client.GetPoolStatus()
	if err != nil {
			fmt.Println(err)
			return
	}

	fmt.Print("Nodes : ")
	fmt.Println(PoolStatus.Nodes)

	return client, endorserSeed, err
}

func Nym(targetDid string, targetVerkey string, sDid string, sign *crypto.Ed25519Signer, client *vdr.Client) (err error) {
	err = client.CreateNym(targetDid, targetVerkey, "", sDid, sign)

	if err != nil {
		return
	}

	return
}

func getNym(did string, client *vdr.Client) (res interface{}, err error) {
	ReadReply, err := client.GetNym(did)
	if err != nil {
		return
	}

	if ReadReply.Data == nil {
		err = errors.New("not exist did.")
		return
	}

	return ReadReply.Data, nil
}
