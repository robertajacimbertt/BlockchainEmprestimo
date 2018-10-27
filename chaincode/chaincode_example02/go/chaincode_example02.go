/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

//WARNING - this chaincode's ID is hard-coded in chaincode_example04 to illustrate one way of
//calling chaincode from a chaincode. If this example is modified, chaincode_example04.go has
//to be modified as well with the new ID of chaincode_example02.
//chaincode_example05 show's how chaincode ID can be passed in as a parameter instead of
//hard-coding.

import (
	"fmt"
	"errors"
	"strconv"
	"encoding/json"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	errors.New("-----------> ex02 Init")
	_, args := stub.GetFunctionAndParameters()
	var A, B string    // Entities
	var Aval, Bval int // Asset holdings
	var err error

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	// Initialize the chaincode
	A = args[0]
	Aval, err = strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	B = args[2]
	Bval, err = strconv.Atoi(args[3])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	fmt.Println("Aval = %d, Bval = %d\n", Aval, Bval)

	// Write the state to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	errors.New("-------> ex02 Invoke")
	function, args := stub.GetFunctionAndParameters()
	if function == "invoke" {
		errors.New("-----> escolheu invoke")
		// Make payment of X units from A to B
		return t.invoke(stub, args)
	} else if function == "delete" {
		errors.New("-----> escolheu delete")
		// Deletes an entity from its state
		return t.delete(stub, args)
	} else if function == "query" {
		errors.New("-----> escolheu query")
		// the old "Query" is now implemtned in invoke
		return t.query(stub, args)
	} else if function == "GetLoanApplication" {
		errors.New("-----> escolheu GetLoanApplication")
		return GetLoanApplication(stub, args)
	} else if function == "getHistory" {
		errors.New("-----> escolheu getHistory")
		return getHistory(stub, args)
	}

	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}

// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	errors.New("-----> entrou no Invoke")
	var A, B string    // Entities
	var Aval, Bval int // Asset holdings
	var X int          // Transaction value
	var err error

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	A = args[0]
	B = args[1]

	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Avalbytes == nil {
		return shim.Error("Entity not found")
	}
	Aval, _ = strconv.Atoi(string(Avalbytes))

	Bvalbytes, err := stub.GetState(B)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Bvalbytes == nil {
		return shim.Error("Entity not found")
	}
	Bval, _ = strconv.Atoi(string(Bvalbytes))

	// Perform the execution
	X, err = strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("Invalid transaction amount, expecting a integer value")
	}
	Aval = Aval - X
	Bval = Bval + X
	fmt.Println("Aval = %d, Bval = %d\n", Aval, Bval)

	// Write the state back to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// Deletes an entity from state
func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	A := args[0]

	// Delete the key from the state in ledger
	err := stub.DelState(A)
	if err != nil {
		return shim.Error("Failed to delete state")
	}

	return shim.Success(nil)
}

// func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
// 	fmt.Println("-------> Query....")
// 	if function == "GetLoanApplication" {
// 		fmt.Println("-------> IF....")
// 		return GetLoanApplication(stub, args)
// 	}
// 	return nil, nil
// }

// query callback representing the query of a chaincode
func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	errors.New("-------> entra no query minusculo....")
	var A string // Entities
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	A = args[0]

	// Get the state from the ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	fmt.Println("Query Response:%s\n", jsonResp)
	return shim.Success(Avalbytes)
}

// func GetState(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
// 	fmt.Println("Entering GetLoanApplication")
// 	if len(args) < 1 {
// 	fmt.Println("Invalid number of arguments")
// 	return nil, errors.New("Missing loan application ID")
// 	}
// 	var loanApplicationId = args[0]
// 	bytes, err := stub.GetState(loanApplicationId)
// 	if err != nil {
// 	fmt.Println("Could not fetch loan application with id "+loanApplicationId+" from ledger", err)
// 	return nil, err
// 	}
// 	fmt.Println("Bytes: ", bytes)
// 	return bytes, nil
// }

// Considerar outra implementacao
// func (stub *ChaincodeStub) GetHistoryForKey(key string) (HistoryQueryIteratorInterface, error) {
// 	response, err := stub.handler.handleGetHistoryForKey(key, stub.ChannelId, stub.TxID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &HistoryQueryIterator{CommonIterator: &CommonIterator{stub.handler, stub.ChannelId, stub.TxID, response, 0}}, nil
// }

func GetLoanApplication(stub shim.ChaincodeStubInterface, args []string) pb.Response  {
	fmt.Println("------> Entering GetLoanApplication")

	// if len(args) < 1 {
	// 	errors.New("--------> Invalid number of arguments")
	// 	return shim.Error()
	// }

	var loanApplicationId = args[0]
	bytes, err := stub.GetState(loanApplicationId)
	if err != nil {
	// 	fmt.Println("---------> Could not fetch loan application with id "+loanApplicationId+" from ledger", err)
	// 	return  err
	}
	return shim.Success(bytes)
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Println("Error starting Simple chaincode: %s", err)
	}
}

func getHistory(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	type Pessoa struct {
		TxId      string `json:"txId"`
		Value     int `json:"value"`
		Timestamp string `json:"timestamp"`
	}
	var history []Pessoa
	var ticket int

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	id_pessoa := args[0]
	fmt.Println("----> start getHistoryForPerson: %s\n", id_pessoa)

	// Get History
	resultsIterator, err := stub.GetHistoryForKey(id_pessoa)
	fmt.Println("----> Result iterator", resultsIterator)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		pointer, err := resultsIterator.Next()
		txID, historicValue := pointer.GetTxId(), pointer.GetValue()
		fmt.Println("----> Iterator 1 ", historicValue)
		if err != nil {
			return shim.Error(err.Error())
		}

		var tx Pessoa
		tx.TxId = txID                         
		json.Unmarshal(historicValue, &ticket) 
		if historicValue == nil {              
			var emptyTicket int
			tx.Value = emptyTicket //copy nil marble
		} else {
			json.Unmarshal(historicValue, &ticket) //un stringify it aka JSON.parse()
			// if ticket.ObjectType == TYPE_TICKET {
				tx.Value = ticket //copy ticket over
			// }
		}
		var ts = time.Unix(pointer.Timestamp.Seconds, int64(pointer.Timestamp.Nanos)).String()
		tx.Timestamp = ts
		history = append(history, tx) //add this tx to the list
	}
	fmt.Println("- getHistoryForTicket returning:\n%s", history)

	//change to array of bytes
	historyAsBytes, _ := json.Marshal(history) //convert to array of bytes
	return shim.Success(historyAsBytes)
}