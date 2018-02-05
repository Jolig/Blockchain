package main

import (
        "errors"
        "fmt"
        "strconv"
        "encoding/json"
        "github.com/hyperledger/fabric/core/chaincode/shim"
)
type SampleChaincode struct {
}

type UserLog struct {
	
	EmailId string `json:"emailid"`
	Username string `json:"username"`
	Role string `json:"role"`
	Action string `json:"action"`
}

func (t *SampleChaincode) Init(stub shim.ChaincodeStubInterface,function string, args []string) ([]byte, error) {
        var err error

        if len(args) != 2 {
                return nil, errors.New("Incorrect number of arguments. Expecting 2.")
        }

     if err!=nil {
                        return nil, err
                }
        record := UserLog{}
        record.EmailId="devyani.bajaj@iiitb.org"
        record.Username = "Devyani"
        record.Role = "Student";
	record.Action = "Enrollment"
		
        
	newrecordByte, err := json.Marshal(record);
        if err!=nil {

            return nil, err
        }
                err=stub.PutState("default",newrecordByte);
         if err!=nil {
                        return nil, err
                }

        return nil, nil
}

func (t *SampleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

var account string

        var err error

        if len(args) != 4 {
                return nil, errors.New("Incorrect number of arguments. Expecting 2.")
        }
          account = args[0]

         recordByte, err := stub.GetState(account);
        fmt.Println(recordByte);
        if err != nil {

            return nil, err
        }
        record := UserLog{}
        if recordByte != nil {
        errrecordmarshal := json.Unmarshal(recordByte,&record);
        if errrecordmarshal != nil {
            return nil, errrecordmarshal
        }            
        }
        record.EmailId   =args[0];
        record.Username =args[1];
	record.Role =args[2];
	record.Action =args[3];
      
	
	newrecordByte, err := json.Marshal(record);
        if err!=nil {

            return nil, err
        }
        err =stub.PutState(account,newrecordByte);
        if err != nil {

            return nil, err;
        } 
        return nil, nil
}


func (t *SampleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
  if function != "read" {
                return nil, errors.New("Invalid query function name. Expecting \"query\".")
        }
        var err error

         if len(args) != 1 {
                return nil, errors.New("Incorrect number of arguments. Expecting name of the state variable to query.")
        }

     var   account = args[0]
        accountValueBytes ,err := stub.GetState(account)
        if err != nil {
                 return nil, err
        }
        return accountValueBytes, nil
}


func main() {
        err := shim.Start(new(SampleChaincode))

        if err != nil {
                fmt.Printf("Error starting CrowdFundChaincode: %s", err)
        }
}








