package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contracapi"
)

type SmartContract struct {
	contracapi.Contract
}

type Book struct {
	Title string `json:"title"`
	Author string `json:"author"`
	Publisher string `json:"publisher"`
	PubDate string `json:"pubdate"`
	Price int `json:"price"`
	Stock int `json:"stock"`
}

type QueryResult struct {
	Key string `json:"Key"`
	Record *Book
}

type Index struct{
	index int
}

func 


func (s *SmartContract) InitLedger(ctx contracapi.TransactionContextInterface) error {
	books := []Book{
		Book{Title:"", Author:"", Publisher:"", PubDate:"", Price:0},
		Book{Title:"", Author:"", Publisher:"", PubDate:"", Price:0},
		Book{Title:"", Author:"", Publisher:"", PubDate:"", Price:0},
		Book{Title:"", Author:"", Publisher:"", PubDate:"", Price:0},
		Book{Title:"", Author:"", Publisher:"", PubDate:"", Price:0},		
	}

	for i, book := range books {
		bookAsBytes, _ := json.Marshal(book)
		err := ctx.GetStub().PutState("BOOK"+strconv.Itoa(i), bookAsBytes)

		if err != nil {
			return fmt.Errorf("set world state failure. %s", err.Error())
		}
	}
	index := Index{5}
	indexAsBytes, _ := json.Marshal(index)
	err := ctx.GetStub().PutState("index", indexAsBytes)

	return nil
}

func (s *SmartContract) AddBook(ctx contracapi.TransactionContextInterface, bookNumber string, title string, author string, publisher string, pubdate string, price int, stock int  ) error {
	book := Book{
		Title: title,
		Author: author,
		Publisher: publisher,
		PubDate: pubdate,
		Price: price,
		Stock: stock,
	}

	bookAsBytes, _ := json.Marshal(book)

	return ctx.GetStub().PutState(bookNumber, bookAsBytes)
}

func (s *SmartContract) QueryBookByTitle() (*Book, error) {

}

func (s *SmartContract) QueryAllBooks() ([]QueryResult, error) {

}

func (s *SmartContract) DelBook() error {

}
