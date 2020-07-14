package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/HouzuoGuo/tiedot/db"
)

// ActionReturn struct
type ActionReturn struct {
	Success   bool
	ReturnErr string
}

func createData(tableName string, data map[string]interface{}) ActionReturn {

	myDB, err := db.OpenDB(dbAddress)
	if err != nil {
		return ActionReturn{false, "Can't open database"}
	}
	defer myDB.Close()

	curTable := myDB.Use(tableName)
	if curTable == nil {
		return ActionReturn{false, "Can't open table"}
	}

	data["CreatedDate"] = time.Now()
	data["CreatedUsername"] = "admin"
	data["ModifiedDate"] = time.Now()
	data["ModifiedUsername"] = "admin"

	var returnedID int
	returnedID, err = curTable.Insert(data)
	if err != nil {
		return ActionReturn{false, "didn't create entry"}
	}

	preData := make(map[string]interface{})
	preData["changed"] = "none"
	AuditCreate(returnedID, tableName, "all", preData, data, "CREATE", "ADMIN")

	return ActionReturn{true, ""}
}

// AuditCreate function
func AuditCreate(docID int, doc string, field string, preV map[string]interface{}, newV map[string]interface{}, auditType string, user string) {

	fmt.Printf("HERE IS DOCID %v\n", docID)

	var errorText string
	var preVC string
	var newVC string

	dataExportPre, err := json.Marshal(preV)
	if err != nil {
		errorText = "didn't convert"
	} else {
		preVC = string(dataExportPre)
	}

	dataExportNew, err := json.Marshal(newV)
	if err != nil {
		errorText = "didn't convert"
	} else {
		newVC = string(dataExportNew)
	}

	dataIn := make(map[string]interface{})
	dataIn["DocID"] = strconv.Itoa(docID)
	dataIn["Doc"] = doc
	dataIn["Field"] = field
	dataIn["PreV"] = preVC
	dataIn["NewV"] = newVC
	dataIn["AuditType"] = auditType
	dataIn["User"] = user
	dataIn["CreatedDate"] = time.Now()

	myDB, err := db.OpenDB(dbAddress)
	if err != nil {
		errorText = "Didn't open database"
	}
	defer myDB.Close()

	curTable := myDB.Use("Audit")
	if curTable == nil {
		errorText = "Can't open table"
	}

	_, err = curTable.Insert(dataIn)
	if err != nil {
		errorText = "didn't create entry"
	}

	if errorText != "" {
		var out string
		dataExport, _ := json.Marshal(dataIn)
		if err != nil {
			out = "didn't convert"
		} else {
			out = string(dataExport)
		}

		file, err := os.OpenFile("error.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		log.SetOutput(file)
		log.Printf("Audit: %v with error %v", out, errorText)

	}

}
