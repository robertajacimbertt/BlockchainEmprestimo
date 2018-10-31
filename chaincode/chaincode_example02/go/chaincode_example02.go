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
	fmt.Printf("-----------> ex02 Init")
	_, args := stub.GetFunctionAndParameters()
	var fin1, fin2, fin3, fin4, fin5, fin6, fin7, fin8, fin9, fin10 string 
	var fin1val, fin2val, fin3val, fin4val, fin5val, fin6val, fin7val, fin8val, fin9val, fin10val int

	var rede string    
	var redeVal int 

	var dev1, dev2, dev3, dev4, dev5, dev6, dev7 string 
	var dev1val, dev2val, dev3val, dev4val, dev5val, dev6val, dev7val int

	// type Person struct {
	// 	name string
	// 	value  string
	// }
	// var fins = []Person{
	// 	Person{
	// 		name: args[0],
	// 		value, err: strconv.Atoi(args[1]),
	// 	},
	// }
	var err error

	// if len(args) != 6 {
	// 	return shim.Error("Incorrect number of arguments. Expecting 1997")
	// }

	// for _,x := range fins {
	// 	fmt.Printf("%s has %d", x.name, x.value)
	// }

	// Initialize the chaincode

	// FINANCIADORES
	fin1 = args[0]
	fin1val, err = strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	fin2 = args[2]
	fin2val, err = strconv.Atoi(args[3])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	fin3 = args[4]
	fin3val, err = strconv.Atoi(args[5])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	fin4 = args[6]
	fin4val, err = strconv.Atoi(args[7])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	fin5 = args[8]
	fin5val, err = strconv.Atoi(args[9])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	fin6 = args[10]
	fin6val, err = strconv.Atoi(args[11])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	fin7 = args[12]
	fin7val, err = strconv.Atoi(args[13])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	fin8 = args[14]
	fin8val, err = strconv.Atoi(args[15])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	fin9 = args[16]
	fin9val, err = strconv.Atoi(args[17])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	fin10 = args[18]
	fin10val, err = strconv.Atoi(args[19])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}

	fmt.Printf("********* fin1val = %d\n", fin1val)

	err = stub.PutState(fin1, []byte(strconv.Itoa(fin1val)))
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(fin2, []byte(strconv.Itoa(fin2val)))
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(fin3, []byte(strconv.Itoa(fin3val)))
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(fin4, []byte(strconv.Itoa(fin4val)))
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(fin5, []byte(strconv.Itoa(fin5val)))
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(fin6, []byte(strconv.Itoa(fin6val)))
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(fin7, []byte(strconv.Itoa(fin7val)))
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(fin8, []byte(strconv.Itoa(fin8val)))
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(fin9, []byte(strconv.Itoa(fin9val)))
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(fin10, []byte(strconv.Itoa(fin10val)))
	if err != nil {
		return shim.Error(err.Error())
	}

	// REDE
	rede = args[20]
	redeVal, err = strconv.Atoi(args[21])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding rede")
	}

	err = stub.PutState(rede, []byte(strconv.Itoa(redeVal)) )
	if err != nil {
		return shim.Error(err.Error())
	}


	// DEVEDORES
	dev1 = args[22]
	dev1val, err = strconv.Atoi(args[23])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding d1")
	}
	dev2 = args[24]
	dev2val, err = strconv.Atoi(args[25])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding d2")
	}
	dev3 = args[26]
	dev3val, err = strconv.Atoi(args[27])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding d3")
	}
	dev4 = args[28]
	dev4val, err = strconv.Atoi(args[29])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding d4")
	}
	dev5 = args[30]
	dev5val, err = strconv.Atoi(args[31])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding d5")
	}
	dev6 = args[32]
	dev6val, err = strconv.Atoi(args[33])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding d6")
	}
	dev7 = args[34]
	dev7val, err = strconv.Atoi(args[35])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding d7")
	}

	err = stub.PutState(dev1, []byte(strconv.Itoa(dev1val)))
	if err != nil {
		return shim.Error(err.Error())
	}	
	err = stub.PutState(dev2, []byte(strconv.Itoa(dev2val)))
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(dev3, []byte(strconv.Itoa(dev3val)))
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(dev4, []byte(strconv.Itoa(dev4val)))
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(dev5, []byte(strconv.Itoa(dev5val)))
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(dev6, []byte(strconv.Itoa(dev6val)))
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(dev7, []byte(strconv.Itoa(dev7val)))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("-------> ex02 Invoke")
	function, args := stub.GetFunctionAndParameters()
	if function == "invokeFinanciador" {
		fmt.Println("-----> escolheu invoke financiador")
		// Make payment of X units from A to network
		return t.invokeFinanciador(stub, args)
	} else if function == "invokeDevedor" {
		fmt.Println("-----> escolheu invoke devedor")
		// Make payment of X units from network to A
		return t.invokeDevedor(stub, args)
	} else if function == "invokePagamentoDevedor" {
		fmt.Println("-----> escolheu invoke pagamento do devedor")
		// Implementa as regras de pagamento para o devedor (nao pode ficar com saldo menor que 0. pode pagar até zerar - em parcelas. )
		return t.invokePagamentoDevedor(stub, args)
	} else if function == "delete" {
		fmt.Println("-----> escolheu delete")
		// Deletes an entity from its state
		return t.delete(stub, args)
	} else if function == "query" {
		fmt.Println("-----> escolheu query")
		// the old "Query" is now implemtned in invoke
		return t.query(stub, args)
	} else if function == "GetLoanApplication" {
		fmt.Println("-----> escolheu GetLoanApplication")
		return GetLoanApplication(stub, args)
	} else if function == "getHistory" {
		fmt.Println("-----> escolheu getHistory")
		return getHistory(stub, args)
	}

	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}

// Transaction makes payment of X units from A to network
func (t *SimpleChaincode) invokeFinanciador(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// somente valor de A. pois vai sempre mandar pra rede
	fmt.Println("-----> entrou no Invoke")
	var A, Rede string    // Entities
	var Aval, Redeval int // Asset holdings
	var X int          // Transaction value
	var err error

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2 on invoke Financiador")
	}

	A = args[0]
	Rede = "rede"

	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		return shim.Error("Failed to get state from user")
	}
	if Avalbytes == nil {
		return shim.Error("Entity user not found")
	}
	Aval, _ = strconv.Atoi(string(Avalbytes))

	if Aval > 0 { //significa que o usuário já possui um emprestimo
		return shim.Error("Usuário já possui um financiamento realizado, não será possivel realizar outro no momento.")
	}

	Redevalbytes, err := stub.GetState(Rede)
	if err != nil {
		return shim.Error("Failed to get state from network")
	}
	if Redevalbytes == nil {
		return shim.Error("Entity network not found")
	}
	Redeval, _ = strconv.Atoi(string(Redevalbytes))

	// Perform the execution
	X, err = strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("Invalid transaction amount, expecting a integer value")
	}

	Aval = Aval + X
	Redeval = Redeval + X
	fmt.Println("User A val = %d, Network val = %d\n", Aval, Redeval)

	// Write the state back to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(Rede, []byte(strconv.Itoa(Redeval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// Transaction makes payment of X units from A to network
func (t *SimpleChaincode) invokeDevedor(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// somente valor de A. pois vai sempre mandar pra rede
	// desconta o valor da rede
	fmt.Println("-----> entrou no Invoke devedor")
	var A, Rede string    // Entities
	var Aval, Redeval int // Asset holdings
	var X int          // Transaction value
	var err error

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2 on invoke Devedor")
	}

	A = args[0]
	Rede = "rede"

	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		return shim.Error("Failed to get state from user")
	}
	if Avalbytes == nil {
		return shim.Error("Entity user not found")
	}
	Aval, _ = strconv.Atoi(string(Avalbytes))

	if Aval < 0 { //significa que o usuário já possui um emprestimo
		return shim.Error("Usuário já possui um empréstimo realizado, não será possivel realizar outro no momento.")
	}

	Redevalbytes, err := stub.GetState(Rede)
	if err != nil {
		return shim.Error("Failed to get state from network")
	}
	if Redevalbytes == nil {
		return shim.Error("Entity network not found")
	}
	Redeval, _ = strconv.Atoi(string(Redevalbytes))

	// Perform the execution
	X, err = strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("Invalid transaction amount, expecting a integer value")
	}

	Aval = Aval - X
	Redeval = Redeval - X
	fmt.Println("User A val = %d, Network val = %d\n", Aval, Redeval)

	// Write the state back to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(Rede, []byte(strconv.Itoa(Redeval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *SimpleChaincode) invokePagamentoDevedor(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// somente valor de A. pois vai sempre mandar pra rede
	// desconta o valor da rede
	fmt.Println("-----> entrou no Invoke pagamento do devedor")
	var A, Rede string    // Entities
	var Aval, Redeval int // Asset holdings
	var X int          // Transaction value
	var err error

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2 on invoke Devedor")
	}

	A = args[0]
	Rede = "rede"

	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		return shim.Error("Failed to get state from user")
	}
	if Avalbytes == nil {
		return shim.Error("Entity user not found")
	}
	Aval, _ = strconv.Atoi(string(Avalbytes))

	Redevalbytes, err := stub.GetState(Rede)
	if err != nil {
		return shim.Error("Failed to get state from network")
	}
	if Redevalbytes == nil {
		return shim.Error("Entity network not found")
	}
	Redeval, _ = strconv.Atoi(string(Redevalbytes))

	// Perform the execution
	X, err = strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("Invalid transaction amount, expecting a integer value")
	}

	if X < ( Aval * 1 ) {//se meu valor pago for menor que o valor que eu devo...
		Aval = Aval + X
		Redeval = Redeval + X
	} else if X === ( Aval * 1 ) {
		Aval = Aval + X
		Redeval = Redeval + X
	}//falta verificar se for ir abaixo de zero

	fmt.Println("User A val = %d, Network val = %d\n", Aval, Redeval)

	// Write the state back to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(Rede, []byte(strconv.Itoa(Redeval)))
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
	errors.New("-------> entra no query ....")
	var A string // Entities
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	A = args[0] // este argumento é o nome da pessoa, ela acessa o asset de mesmo nome

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