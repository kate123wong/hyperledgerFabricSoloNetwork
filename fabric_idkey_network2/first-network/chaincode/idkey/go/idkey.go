package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {


	function, args := APIstub.GetFunctionAndParameters()
	if function == "queryKey" {
		return s.queryKey(APIstub, args)
	} else if function == "insertIdKeyPair" {
		return s.insertIdKeyPair(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) queryKey(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	fmt.Println("ex02 queryKey")
	var id string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting id to query")
	}

	id = args[0]
	
	keyAsBytes, err := APIstub.GetState(id)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for" + id + "\"}"
		return shim.Error(jsonResp)
	}
	
	if keyAsBytes == nil {
		jsonResp := "{\"Error\":\"Nil key for " + id + "\"}"
		return shim.Error(jsonResp)
	}
	
	jsonResp := "{\"id\":\"" + id + "\",\"key\":\"" + string(keyAsBytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return shim.Success(keyAsBytes)
}

func (s *SmartContract) insertIdKeyPair(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	var id string
	var key int
	var err error
	var e error
	
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	id = args[0]
	idlength := utf8.RuneCountInString(id)
	if idlength != 32 {
		return shim.Error("Incorrect id. Expecting id length is 32")
	}

	keyAsBytes, err := APIstub.GetState(id)
	if err != nil {
		return shim.Error(err.Error())
	}

	if keyAsBytes == nil {
		key, e = strconv.Atoi(args[1])
		if e != nil {
			return shim.Error("Expecting integer value for key")
		}
		fmt.Printf("id = %d, key = %d\n", id, key)

		APIstub.PutState(id, []byte(strconv.Itoa(key)))

		return shim.Success(nil)
	}
	
	jsonResp := "{\"Error\":\"id: "+ id + "aleary exists\"}"
	return shim.Error(jsonResp)

}

func main() {

	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
