/*
Copyright GenomeChain Inc.
Open Source License: Creative Commons 4.0
Developer : Ghanendra
*/

// ====CHAINCODE EXECUTION SAMPLES (Python Jupyter) ==================
// ==== Invoke genes, pass genomic private data as encoded bytes in transient map ====
// sample_gene = b'{"id":11,"association":"significant","population":"French","variant":"APOB","gene":"ADRB2","name":"Ron",price":99}'
// transient_map = {'gene_owner':sample_gene.encode()}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

type gene struct {
	ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database
	Id         int    `json:"id"`      //the fieldtags are needed to keep case from bouncing around
	Name       string `json:name`
	Population string `json:"population"`
	Gene       string `json:"gene"`
	Size       int    `json:"size"`
}

type genePrivateDetails struct {
	ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database
	Name       string `json:"name"`    //the fieldtags are needed to keep case from bouncing around
	Age        int    `json:"age"`
	Varient    string `json:"varient"`
	Price      int    `json:"price"`
}

// ===================================================================================
// Main
// ===================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init initializes chaincode
// ===========================
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke - Our entry point for Invocations
// ========================================
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

	// Handle different functions
	switch function {
	case "initGene":
		//create a new gene
		return t.initGene(stub, args)
	case "readGene":
		//read a gene data
		return t.readGene(stub, args)
	case "readGenePrivateDetails":
		//read a gene private details
		return t.readGenePrivateDetails(stub, args)
	case "transferGene":
		//change owner of a specific gene information
		return t.transferGene(stub, args)
	case "delete":
		//delete a gene
		return t.delete(stub, args)
	case "queryLongetivityMapByGene":
		//find genes from LongetivityMap for user X using rich query
		return t.queryLongetivityMapByGene(stub, args)
	case "queryAgeingDrugs":
		//find Ageing Drugs based on Compound name using rich query
		return t.queryAgeingDrugs(stub, args)
	case "getGenesByRange":
		//get genes based on range query
		return t.getGenesByRange(stub, args)
	default:
		//error
		fmt.Println("invoke did not find func: " + function)
		return shim.Error("Received unknown function invocation")
	}
}

// ============================================================
// initGene - create a new gene, store into chaincode state
// ============================================================
func (t *SimpleChaincode) initGene(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	type geneTransientInput struct {
		Id         int    `json:"id"` //the fieldtags are needed to keep case from bouncing around
		Name       string `json:name`
		Population string `json:"population"`
		Gene       string `json:"gene"`
		Size       int    `json:"size"`
		Age        int    `json:"age"`
		Varient    string `json:"varient"`
		Price      int    `json:"price"`
	}

	// ==== Input sanitation ====
	fmt.Println("- start init gene")

	if len(args) != 0 {
		return shim.Error("Incorrect number of arguments. Private gene data must be passed in transient map.")
	}

	transMap, err := stub.GetTransient()
	if err != nil {
		return shim.Error("Error getting transient: " + err.Error())
	}

	if _, ok := transMap["gene"]; !ok {
		return shim.Error("gene must be a key in the transient map")
	}

	if len(transMap["gene"]) == 0 {
		return shim.Error("gene value in the transient map must be a non-empty JSON string")
	}

	var geneInput geneTransientInput
	err = json.Unmarshal(transMap["gene"], &geneInput)
	if err != nil {
		return shim.Error("Failed to decode JSON of: " + string(transMap["gene"]))
	}

	if geneInput.Id <= 0 {
		return shim.Error("id field must be a positive integer")
	}
	if len(geneInput.Name) == 0 {
		return shim.Error("name field must be a non-empty string")
	}
	if len(geneInput.Population) == 0 {
		return shim.Error("population field must be a non-empty string")
	}
	if len(geneInput.Gene) == 0 {
		return shim.Error("gene field must be a non-empty string")
	}
	if geneInput.Size <= 0 {
		return shim.Error("size field must be a positive integer")
	}
	if geneInput.Age <= 0 {
		return shim.Error("Age field must be a positive integer")
	}
	if len(geneInput.Varient) == 0 {
		return shim.Error("varient field must be a non-empty string")
	}
	if geneInput.Price <= 0 {
		return shim.Error("Price field must be a positive integer")
	}
	// ==== Check if gene already exists ====
	geneAsBytes, err := stub.GetPrivateData("collectionGenes", geneInput.Name)
	if err != nil {
		return shim.Error("Failed to get gene: " + err.Error())
	} else if geneAsBytes != nil {
		fmt.Println("This gene already exists: " + geneInput.Name)
		return shim.Error("This gene already exists: " + geneInput.Name)
	}

	// ==== Create gene object, marshal to JSON, and save to state ====
	gene := &gene{
		ObjectType: "gene",
		Id:         geneInput.Id,
		Name:       geneInput.Name,
		Population: geneInput.Population,
		Gene:       geneInput.Gene,
		Size:       geneInput.Size,
	}
	geneJSONasBytes, err := json.Marshal(gene)
	if err != nil {
		return shim.Error(err.Error())
	}

	// === Save gene to state ===
	err = stub.PutPrivateData("collectionGenes", geneInput.Name, geneJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ==== Create gene private details object with price, marshal to JSON, and save to state ====
	genePrivateDetails := &genePrivateDetails{
		ObjectType: "genePrivateDetails",
		Name:       geneInput.Name,
		Age:        geneInput.Age,
		Varient:    geneInput.Varient,
		Price:      geneInput.Price,
	}
	genePrivateDetailsBytes, err := json.Marshal(genePrivateDetails)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutPrivateData("collectionGenesPrivateDetails", geneInput.Name, genePrivateDetailsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	//  ==== Index the gene to enable name-based range queries, e.g. return all genes ====
	//  An 'index' is a normal key/value entry in state.
	indexName := "gene~name"
	geneNameIndexKey, err := stub.CreateCompositeKey(indexName, []string{gene.Gene, gene.Name})
	if err != nil {
		return shim.Error(err.Error())
	}
	//  Save index entry to state. Only the key name is needed, no need to store a duplicate copy of the gene.
	//  Note - passing a 'nil' value will effectively delete the key from state, therefore we pass null character as value
	value := []byte{0x00}
	stub.PutPrivateData("collectionGenes", geneNameIndexKey, value)

	// ==== Gene saved and indexed. Return success ====
	fmt.Println("- end init gene")
	return shim.Success(nil)
}

// ===============================================
// readGene - read a gene from chaincode state
// ===============================================
func (t *SimpleChaincode) readGene(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var name, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the gene to query")
	}

	name = args[0]
	valAsbytes, err := stub.GetPrivateData("collectionGenes", name) //get the gene from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + name + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Gene does not exist: " + name + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}

// ===============================================
// readGenePrivateDetails - read a gene private details from chaincode state
// ===============================================
func (t *SimpleChaincode) readGenePrivateDetails(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var name, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the gene to query")
	}

	name = args[0]
	valAsbytes, err := stub.GetPrivateData("collectionGenesPrivateDetails", name) //get the gene private details from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get private details for " + name + ": " + err.Error() + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Gene private details does not exist: " + name + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}

// ==================================================
// delete - remove a gene key/value pair from state
// ==================================================
func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("- start delete gene")

	type geneDeleteTransientInput struct {
		Name string `json:"name"`
	}

	if len(args) != 0 {
		return shim.Error("Incorrect number of arguments. Private gene name must be passed in transient map.")
	}

	transMap, err := stub.GetTransient()
	if err != nil {
		return shim.Error("Error getting transient: " + err.Error())
	}

	if _, ok := transMap["gene_delete"]; !ok {
		return shim.Error("gene_delete must be a key in the transient map")
	}

	if len(transMap["gene_delete"]) == 0 {
		return shim.Error("gene_delete value in the transient map must be a non-empty JSON string")
	}

	var geneDeleteInput geneDeleteTransientInput
	err = json.Unmarshal(transMap["gene_delete"], &geneDeleteInput)
	if err != nil {
		return shim.Error("Failed to decode JSON of: " + string(transMap["gene_delete"]))
	}

	if len(geneDeleteInput.Name) == 0 {
		return shim.Error("name field must be a non-empty string")
	}

	// to maintain the id~name index, we need to read the gene first and get its color
	valAsbytes, err := stub.GetPrivateData("collectionGenes", geneDeleteInput.Name) //get the gene from chaincode state
	if err != nil {
		return shim.Error("Failed to get state for " + geneDeleteInput.Name)
	} else if valAsbytes == nil {
		return shim.Error("Gene does not exist: " + geneDeleteInput.Name)
	}

	var geneToDelete gene
	err = json.Unmarshal([]byte(valAsbytes), &geneToDelete)
	if err != nil {
		return shim.Error("Failed to decode JSON of: " + string(valAsbytes))
	}

	// delete the gene from state
	err = stub.DelPrivateData("collectionGenes", geneDeleteInput.Name)
	if err != nil {
		return shim.Error("Failed to delete state:" + err.Error())
	}

	// Also delete the gene from the color~name index
	indexName := "color~name"
	geneNameIndexKey, err := stub.CreateCompositeKey(indexName, []string{geneToDelete.Gene, geneToDelete.Name})
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.DelPrivateData("collectionGenes", geneNameIndexKey)
	if err != nil {
		return shim.Error("Failed to delete state:" + err.Error())
	}

	// Finally, delete private details of gene
	err = stub.DelPrivateData("collectionGenesPrivateDetails", geneDeleteInput.Name)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// ===========================================================
// transfer a gene by setting a new owner name on the gene
// ===========================================================
func (t *SimpleChaincode) transferGene(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println("- start transfer gene")

	type geneTransferTransientInput struct {
		Gene string `json:"gene"`
		Name string `json:"name"`
	}

	if len(args) != 0 {
		return shim.Error("Incorrect number of arguments. Private gene data must be passed in transient map.")
	}

	transMap, err := stub.GetTransient()
	if err != nil {
		return shim.Error("Error getting transient: " + err.Error())
	}

	if _, ok := transMap["gene_name"]; !ok {
		return shim.Error("gene_name must be a key in the transient map")
	}

	if len(transMap["gene_name"]) == 0 {
		return shim.Error("gene_name value in the transient map must be a non-empty JSON string")
	}

	var geneTransferInput geneTransferTransientInput
	err = json.Unmarshal(transMap["gene_name"], &geneTransferInput)
	if err != nil {
		return shim.Error("Failed to decode JSON of: " + string(transMap["gene_name"]))
	}

	if len(geneTransferInput.Gene) == 0 {
		return shim.Error("gene field must be a non-empty string")
	}
	if len(geneTransferInput.Name) == 0 {
		return shim.Error("name field must be a non-empty string")
	}

	geneAsBytes, err := stub.GetPrivateData("collectionGenes", geneTransferInput.Name)
	if err != nil {
		return shim.Error("Failed to get gene:" + err.Error())
	} else if geneAsBytes == nil {
		return shim.Error("Name does not exist: " + geneTransferInput.Name)
	}

	geneToTransfer := gene{}
	err = json.Unmarshal(geneAsBytes, &geneToTransfer) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	geneToTransfer.Gene = geneTransferInput.Gene //change the name

	geneJSONasBytes, _ := json.Marshal(geneToTransfer)
	err = stub.PutPrivateData("collectionGenes", geneToTransfer.Name, geneJSONasBytes) //rewrite the gene
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end transferGene (success)")
	return shim.Success(nil)
}

// ===========================================================================================
// getGenesByRange performs a range query based on the start and end keys provided.

// ===========================================================================================
func (t *SimpleChaincode) getGenesByRange(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	startKey := args[0]
	endKey := args[1]

	resultsIterator, err := stub.GetPrivateDataByRange("collectionGenes", startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getGenesByRange queryResult:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

// ===== Rich query =================================================
// queryGenesByOwner queries for genes based on a passed in owner.
// Only available on state databases (e.g. CouchDB)
// ==================================================================
func (t *SimpleChaincode) queryAgeingDrugs(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//   0
	// "bob"
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	owner := strings.ToLower(args[0])

	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"gene\",\"owner\":\"%s\"}}", owner)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

// ===== rich query ========================================================
// queryGenes uses a query string to perform a query for genes.
// Only available on state databases that support rich query (e.g. CouchDB)
// =========================================================================
func (t *SimpleChaincode) queryLongetivityMapByGene(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//   0
	// "queryString"
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	queryString := args[0]

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

// =========================================================================================
// getQueryResultForQueryString executes the passed in query string.
// Result set is built and returned as a byte array containing the JSON results.
// =========================================================================================
func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetPrivateDataQueryResult("collectionGenes", queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}
