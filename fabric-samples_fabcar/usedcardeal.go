/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing a car
type SmartContract struct {
	contractapi.Contract
}

// Car describes basic details of what makes up a car
type Car struct {
	Make   string `json:"make"`
	Model  string `json:"model"`
	Colour string `json:"colour"`
	Owner  string `json:"owner"`
	ProductionDate string `json:"productiondata"`

	Appraiser string `json:"appraiser"`
	MeasuredPrice int `json:"measuredprice"`
}

type Customer struct {
	Name string `json:"name"`
	Amount int `json:"amout"`
}


// QueryResult structure used for handling result of query
type QueryCarResult struct {
	Key    string `json:"Key"`
	Record *Car
}

type QueryCustomerResult struct {
	Key string `json:"Key"`
	Record *Customer
}

// InitLedger adds a base set of cars to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	cars := []Car{
		Car{Make: "Toyota", Model: "Prius", Colour: "blue", Owner: "Tomoko", ProductionDate: "2014-03-05", Appraiser: nil, MeasuredPrice: 0},
		Car{Make: "Ford", Model: "Mustang", Colour: "red", Owner: "Brad", ProductionDate: "2017-01-08", Appraiser: nil, MeasuredPrice: 0},
		Car{Make: "Hyundai", Model: "Tucson", Colour: "green", Owner: "Jin Soo", ProductionDate: "2008-09-05", Appraiser: nil, MeasuredPrice: 0},
		Car{Make: "Volkswagen", Model: "Passat", Colour: "yellow", Owner: "Max", ProductionDate: "2019-01-05", Appraiser: nil, MeasuredPrice: 0},
		Car{Make: "Tesla", Model: "S", Colour: "black", Owner: "Adriana", ProductionDate: "2001-01-21", Appraiser: nil, MeasuredPrice: 0},
		Car{Make: "Peugeot", Model: "205", Colour: "purple", Owner: "Michel", ProductionDate: "2015-06-21", Appraiser: nil, MeasuredPrice: 0},
		Car{Make: "Holden", Model: "Barina", Colour: "brown", Owner: "Shotaro", ProductionDate: "2019-04-01", Appraiser: nil, MeasuredPrice: 0},
	}

	customers := []Customer{
		Customer{Name:"Tomoko", Amount:0},
		Customer{Name:"Brad", Amount:0},
		Customer{Name:"Jin Soo", Amount:0},
		Customer{Name:"Max", Amount:0},
		Customer{Name:"Adriana", Amount:0},
		Customer{Name:"Michel", Amount:0},
		Customer{Name:"Shotaro", Amount:0},
	}

	for i, car := range cars {
		carAsBytes, _ := json.Marshal(car)
		err := ctx.GetStub().PutState("CAR"+strconv.Itoa(i), carAsBytes)

		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}

	for i, customer := range customers {
		customerAsBytes, _ := json.Marshal(customer)
		err := ctx.GetStub().PutState("Customer"+strconv.Itoa(i), customerAsBytes)

		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}

	return nil
}
/*********************************************************************************************************************************************************/
// CreateCar adds a new car to the world state with given details
func (s *SmartContract) CreateCar(ctx contractapi.TransactionContextInterface, carNumber string, make string, model string, colour string, owner string, productiondata string) error {
	car := Car{
		Make:   make,
		Model:  model,
		Colour: colour,
		Owner:  owner,
		ProductionDate: productiondata,
		Appraiser: nil,
		MeasuredPrice: 0,
	}

	carAsBytes, _ := json.Marshal(car)

	return ctx.GetStub().PutState(carNumber, carAsBytes)
}

// QueryCar returns the car stored in the world state with given id
func (s *SmartContract) QueryCar(ctx contractapi.TransactionContextInterface, carNumber string) (*Car, error) {
	carAsBytes, err := ctx.GetStub().GetState(carNumber)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if carAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", carNumber)
	}

	car := new(Car)
	_ = json.Unmarshal(carAsBytes, car)

	return car, nil
}

// QueryAllCars returns all cars found in world state
func (s *SmartContract) QueryAllCars(ctx contractapi.TransactionContextInterface) ([]QueryResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []QueryResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		car := new(Car)
		_ = json.Unmarshal(queryResponse.Value, car)

		queryResult := QueryResult{Key: queryResponse.Key, Record: car}
		results = append(results, queryResult)
	}

	return results, nil
}

// ChangeCarOwner updates the owner field of car with given id in world state
/*
func (s *SmartContract) ChangeCarOwner(ctx contractapi.TransactionContextInterface, carNumber string, newOwner string) error {
	car, err := s.QueryCar(ctx, carNumber)

	if err != nil {
		return err
	}

	car.Owner = newOwner

	carAsBytes, _ := json.Marshal(car)

	return ctx.GetStub().PutState(carNumber, carAsBytes)
}
*/

func (s *SmartContract) AppraiseCar(ctx contractapi.TransactionContextInterface, carNumber string, Appraiser string, price int) error {
	car, err := s.QueryCar(ctx, carNumber)

	if err != nil {
		return err
	}

	car.Appraiser = Appraiser
	car.MeasuredPrice = price

	carAsBytes, _ := json.Marshal(car)

	return ctx.GetStub().PutState(carNumber, carAsBytes)
}
/*********************************************************************************************************************************************************/
func (s *SmartContract) RegisterUser(ctx contractapi.TransactionContextInterface, customerNumber string, customerName string) error {
	startKey := ""
	endKey := ""
	
	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []QueryResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		customer := new(Customer)
		_ = json.Unmarshal(queryResponse.Value, customer)

		queryResult := QueryResult{Key: queryResponse.Key, Record: customer}
		results = append(results, queryResult)
	}

	for _, value range results {
		if value.Name == customerName{
			return nil, fmt.Errorf("%s is already exist", customerName)
		}
	}

	customer := Customer{
		Name: customerName,
		Amount: 0,
	}

	customerAsBytes, _ := json.Marshal(customer)

	return ctx.GetStub().PutState(customerNumber, customerAsBytes)
}

func (s *SmartContract) QueryCustomer(ctx contractapi.TransactionContextInterface, customerNumber string) (*Customer, error) {
	customerAsBytes, err := ctx.GetStub().GetState(customerNumber)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if customerAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", customerNumber)
	}

	customer := new(Customer)
	_ = json.Unmarshal(customerAsBytes, customer)

	return customer, nil
}




func (s *SmartContract) ChangeCarOwner(ctx contractapi.TransactionContextInterface, carNumber string, buyerNumber string, sellerNumber string) error {
	car, err_car := s.QueryCar(ctx, carNumber)
	buyer, err_buyer := s.QueryCustomer(ctx, buyerNumber)
	seller, err_seller := s.QueryCustomer(ctx, sellerNumber)

	if err_car != nil || err_customer || err_seller {
		return err
	}

	if car.MeasuredPrice == 0 {
		return fmt.Errorf("Checked MeasuredPrice.")
	}
	if buyer.Amount < car.MeasuredPrice {
		return fmt.Errorf("Checked buyer's Amount")
	}

	buyer.Amount -= car.MeasuredPrice
	seller.Amount += car.MeasuredPrice

	car.Owner = buyer.Name

	carAsBytes, _ := json.Marshal(car)
	buyerAsBytes, _ := json.Marshal(buyer)
	sellerAsBytes, _ := json.Marshal(seller)

	ctx.GetStub().PutState(carNumber, carAsBytes)
	ctx.GetStub().PutState(buyerNumber, buyerAsBytes)
	ctx.GetStub().PutState(sellerNumber, sellerAsBytes)
	
	return nil
}

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create fabcar chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting fabcar chaincode: %s", err.Error())
	}
}
