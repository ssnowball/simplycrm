package main

import (
	"encoding/json"
	"strconv"

	"github.com/HouzuoGuo/tiedot/db"
	guuid "github.com/google/uuid"
)

func readSubData(tableName string) readReturn {

	queryResult := make(map[int]struct{}) // query result (document IDs) goes into map keys
	var query interface{}
	var NewResults []map[string]interface{}
	var tableOrder []TableOrder

	json.Unmarshal([]byte(`["all"]`), &query)

	switch tableName {
	case "Lead":
		tableOrder = append(tableOrder, TableOrder{Prop: "Name", Label: "Name"})
		tableOrder = append(tableOrder, TableOrder{Prop: "SType", Label: "SType"})
		tableOrder = append(tableOrder, TableOrder{Prop: "CreatedDate", Label: "CreatedDate"})
		tableOrder = append(tableOrder, TableOrder{Prop: "Description", Label: "Description"})
		tableOrder = append(tableOrder, TableOrder{Prop: "Status", Label: "Status"})
		tableOrder = append(tableOrder, TableOrder{Prop: "AccountName", Label: "AccountName"})
	case "Account":
		tableOrder = append(tableOrder, TableOrder{Prop: "Name", Label: "Name"})
		tableOrder = append(tableOrder, TableOrder{Prop: "CType", Label: "CType"})
		tableOrder = append(tableOrder, TableOrder{Prop: "CreatedDate", Label: "CreatedDate"})

	case "Campaign":
		tableOrder = append(tableOrder, TableOrder{Prop: "Name", Label: "Name"})
		tableOrder = append(tableOrder, TableOrder{Prop: "Description", Label: "Description"})
		tableOrder = append(tableOrder, TableOrder{Prop: "TypeChannel", Label: "TypeChannel"})
		tableOrder = append(tableOrder, TableOrder{Prop: "TypeMethod", Label: "TypeMethod"})
		tableOrder = append(tableOrder, TableOrder{Prop: "Status", Label: "Status"})
		tableOrder = append(tableOrder, TableOrder{Prop: "StartDate", Label: "StartDate"})
		tableOrder = append(tableOrder, TableOrder{Prop: "EndDate", Label: "EndDate"})
		tableOrder = append(tableOrder, TableOrder{Prop: "CreatedDate", Label: "CreatedDate"})

	case "Product":
		tableOrder = append(tableOrder, TableOrder{Prop: "Name", Label: "Name"})
		tableOrder = append(tableOrder, TableOrder{Prop: "Description", Label: "Description"})
		tableOrder = append(tableOrder, TableOrder{Prop: "Cost", Label: "Cost"})
		tableOrder = append(tableOrder, TableOrder{Prop: "Category", Label: "Category"})
		tableOrder = append(tableOrder, TableOrder{Prop: "SubCategory", Label: "SubCategory"})
		tableOrder = append(tableOrder, TableOrder{Prop: "CreatedDate", Label: "CreatedDate"})

	default:
		return readReturn{nil, nil, "table name not matched"}
	}

	myDB, err := db.OpenDB(dbAddress)
	if err != nil {
		return readReturn{nil, nil, "Can't open database"}
	}
	defer myDB.Close()

	curTable := myDB.Use(tableName)
	if curTable == nil {
		return readReturn{nil, nil, "Can't open table"}
	}

	if err := db.EvalQuery(query, curTable, &queryResult); err != nil {
		return readReturn{nil, nil, "Can't build query"}
	}

	if len(queryResult) == 0 {
		return readReturn{nil, nil, "No data"}
	}

	for docID := range queryResult {
		readBack, err := curTable.Read(docID)
		if err != nil {
			return readReturn{nil, nil, "Can't read results"}
		}

		for k, v := range readBack {
			readBack[k] = convertData("READ", k, v)
		}

		readBack["DocID"] = strconv.Itoa(docID)
		readBack["Table"] = tableName
		readBack["UUID"] = guuid.New()

		NewResults = append(NewResults, readBack)

	}

	return readReturn{NewResults, tableOrder, ""}
}
