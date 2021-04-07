usedcardeal.go, library.go
==============
Refered to fabcar.go

Added 
=====

##usedcardeal.go
#### Struct
```
//Car
ProductionDate string
Appraiser string
MeasuredPrice int
```      
      
```
//Customer
Name string
Amount int
```

#### Function

```
AppraiseCar(ctx contractapi.TransactionContextInterface, carNumber string, Appraiser string, price int) err
RegisterUser(ctx contractapi.TransactionContextInterface, customerNumber string, customerName string) err
QueryCustomer(ctx contractapi.TransactionContextInterface, customerNumber string) (*Customer, error)
ChangeCarOwner(ctx contractapi.TransactionContextInterface, carNumber string, buyerNumber string, sellerNumber string) err
```
           
           
##library.go

