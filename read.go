package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/HouzuoGuo/tiedot/db"
)

// ReadOptions struct
type ReadOptions struct {
	Convert string
}

// TableOrder struct
type TableOrder struct {
	Prop  string `json:"prop"`
	Label string `json:"label"`
}

// CalDate function
func CalDate(in string) string {

	layout := "2006-01-02"
	t, _ := time.Parse(layout, in[:10])

	outLayout := "02 Jan 2006"

	return t.UTC().Format(outLayout)
}

// PickerDate function
func PickerDate(in string) string {

	layout := "2006-01-02"
	t, _ := time.Parse(layout, in[:10])

	outLayout := "02/01/2006"

	return t.UTC().Format(outLayout)
}

// Reduce function
func Reduce(in []interface{}) (out string) {
	out = strconv.Itoa(len(in))
	return
}

// OutDate function
func OutDate(in string) string {

	layout := "2006-01-02T15:04:05.999999999Z07:00"
	t, _ := time.Parse(layout, in)

	outLayout := "02 Jan 2006 @ 3:04pm (MST)"

	return t.UTC().Format(outLayout)
}

type readReturn struct {
	Data      []map[string]interface{}
	Titles    []TableOrder
	ReturnErr string
}

func convertData(option string, key string, value interface{}) string {

	convertList := make(map[string]ReadOptions)

	if option == "READ" {
		convertList["CreatedDate"] = ReadOptions{Convert: "OutDate"}
		convertList["ModifiedDate"] = ReadOptions{Convert: "OutDate"}
		convertList["Products"] = ReadOptions{Convert: "Reduce"}
		convertList["Campaigns"] = ReadOptions{Convert: "Reduce"}
		// convertList["StartDate"] = ReadOptions{Convert: "CalDate"}
		// convertList["EndDate"] = ReadOptions{Convert: "CalDate"}
	} else {
		convertList["Products"] = ReadOptions{Convert: "json"}
		convertList["Campaigns"] = ReadOptions{Convert: "json"}
		// convertList["StartDate"] = ReadOptions{Convert: "CalDate"}
		// convertList["EndDate"] = ReadOptions{Convert: "PickerDate"}
	}

	if k1, found := convertList[key]; found {
		switch k1.Convert {
		case "json":
			dataExport, _ := json.Marshal(value)
			return string(dataExport)
		case "CalDate":
			return CalDate(value.(string))
		case "OutDate":
			return OutDate(value.(string))
		case "PickerDate":
			return PickerDate(value.(string))
		case "Reduce":
			return Reduce(value.([]interface{}))
		}
	}

	switch v := value.(type) {
	// case []interface {}:
	// 	return
	case int:
		return strconv.Itoa(v)
	case float64:
		return fmt.Sprintf("%.2f", v)
	}

	return value.(string)

}

func readData(tableName string) readReturn {

	runActivities := false
	runLeads := false

	queryResult := make(map[int]struct{}) // query result (document IDs) goes into map keys

	var query interface{}
	var NewResults []map[string]interface{}
	var tableOrder []TableOrder

	json.Unmarshal([]byte(`["all"]`), &query)

	switch tableName {
	case "Activity":
		tableOrder = append(tableOrder, TableOrder{Prop: "Activity", Label: "Activity"})
		tableOrder = append(tableOrder, TableOrder{Prop: "Detail", Label: "Detail"})
		tableOrder = append(tableOrder, TableOrder{Prop: "Type", Label: "Type"})
		tableOrder = append(tableOrder, TableOrder{Prop: "ReferenceName", Label: "ReferenceName"})
		tableOrder = append(tableOrder, TableOrder{Prop: "CreatedDate", Label: "CreatedDate"})
		tableOrder = append(tableOrder, TableOrder{Prop: "ModifiedDate", Label: "ModifiedDate"})
	case "Lead":
		runActivities = true
		tableOrder = append(tableOrder, TableOrder{Prop: "Name", Label: "Name"})
		tableOrder = append(tableOrder, TableOrder{Prop: "Description", Label: "Description"})
		tableOrder = append(tableOrder, TableOrder{Prop: "SType", Label: "SType"})
		tableOrder = append(tableOrder, TableOrder{Prop: "SSubType", Label: "SSubType"})
		tableOrder = append(tableOrder, TableOrder{Prop: "AccountName", Label: "AccountName"})
		tableOrder = append(tableOrder, TableOrder{Prop: "Status", Label: "Status"})
		tableOrder = append(tableOrder, TableOrder{Prop: "Stage", Label: "Stage"})
		tableOrder = append(tableOrder, TableOrder{Prop: "LeadsourceType", Label: "LeadsourceType"})
		tableOrder = append(tableOrder, TableOrder{Prop: "LeadsourceChannel", Label: "LeadsourceChannel"})
		tableOrder = append(tableOrder, TableOrder{Prop: "SaleValue", Label: "SaleValue"})
		tableOrder = append(tableOrder, TableOrder{Prop: "EstCloseDate", Label: "EstCloseDate"})
		tableOrder = append(tableOrder, TableOrder{Prop: "ActualCloseDate", Label: "ActualCloseDate"})
		tableOrder = append(tableOrder, TableOrder{Prop: "CloseReason", Label: "CloseReason"})
		tableOrder = append(tableOrder, TableOrder{Prop: "InvoiceID", Label: "InvoiceID"})
		tableOrder = append(tableOrder, TableOrder{Prop: "ActivityCount", Label: "ActivityCount"})
		tableOrder = append(tableOrder, TableOrder{Prop: "Products", Label: "Products"})
		tableOrder = append(tableOrder, TableOrder{Prop: "Campaigns", Label: "Campaigns"})

	case "Account":
		runActivities = true
		runLeads = true
		tableOrder = append(tableOrder, TableOrder{Prop: "Name", Label: "Name"})
		tableOrder = append(tableOrder, TableOrder{Prop: "CType", Label: "CType"})
		tableOrder = append(tableOrder, TableOrder{Prop: "CreatedDate", Label: "CreatedDate"})
		tableOrder = append(tableOrder, TableOrder{Prop: "ContactName", Label: "ContactName"})
		tableOrder = append(tableOrder, TableOrder{Prop: "Postcode", Label: "Postcode"})
		tableOrder = append(tableOrder, TableOrder{Prop: "LeadCount", Label: "LeadCount"})
		tableOrder = append(tableOrder, TableOrder{Prop: "ActivityCount", Label: "ActivityCount"})

	case "Audit":
		tableOrder = append(tableOrder, TableOrder{Prop: "CreatedDate", Label: "CreatedDate"})
		tableOrder = append(tableOrder, TableOrder{Prop: "AuditType", Label: "AuditType"})
		tableOrder = append(tableOrder, TableOrder{Prop: "Doc", Label: "Doc"})
		tableOrder = append(tableOrder, TableOrder{Prop: "PreV", Label: "PreV"})
		tableOrder = append(tableOrder, TableOrder{Prop: "NewV", Label: "NewV"})

	case "Campaign":
		tableOrder = append(tableOrder, TableOrder{Prop: "Name", Label: "Name"})
		tableOrder = append(tableOrder, TableOrder{Prop: "Description", Label: "Description"})
		tableOrder = append(tableOrder, TableOrder{Prop: "TypeChannel", Label: "TypeChannel"})
		tableOrder = append(tableOrder, TableOrder{Prop: "TypeMethod", Label: "TypeMethod"})
		tableOrder = append(tableOrder, TableOrder{Prop: "Budget", Label: "Budget"})
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

	case "Option":
		tableOrder = append(tableOrder, TableOrder{Prop: "CreatedDate", Label: "CreatedDate"})
		tableOrder = append(tableOrder, TableOrder{Prop: "Name", Label: "Name"})
		tableOrder = append(tableOrder, TableOrder{Prop: "Company", Label: "Company"})

	default:
		return readReturn{nil, nil, "table name not matched - " + tableName}
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

		if runLeads {
			// add count of leads to NewResults
			queryLeadResult := make(map[int]struct{})

			var queryLead interface{}
			// take docID and lookup in activity table
			leadTable := myDB.Use("Lead")
			if leadTable == nil {
				return readReturn{nil, nil, "Can't open lead sub table"}
			}

			json.Unmarshal([]byte(`[{"eq": "`+strconv.Itoa(docID)+`", "in": ["AccountID"]}]`), &queryLead)

			if err := db.EvalQuery(queryLead, leadTable, &queryLeadResult); err != nil {
				return readReturn{nil, nil, "Can't build lead query"}
			}

			readBack["LeadCount"] = strconv.Itoa(len(queryLeadResult))

		}

		if runActivities {
			// add count of activities to NewResults
			queryActivityResult := make(map[int]struct{})
			var queryActivity interface{}
			// take docID and lookup in activity table
			activityTable := myDB.Use("Activity")
			if activityTable == nil {
				return readReturn{nil, nil, "Can't open activty sub table"}
			}

			json.Unmarshal([]byte(`[{"eq": "`+strconv.Itoa(docID)+`", "in": ["ReferenceID"]}]`), &queryActivity)

			if err := db.EvalQuery(queryActivity, activityTable, &queryActivityResult); err != nil {
				return readReturn{nil, nil, "Can't build activity query"}
			}

			readBack["ActivityCount"] = strconv.Itoa(len(queryActivityResult))

		}

		NewResults = append(NewResults, readBack)

	}

	return readReturn{NewResults, tableOrder, ""}
}

type updateReturn struct {
	Data      map[string]interface{}
	ReturnErr string
}

func getUpdateData(tableName string, ID string) updateReturn {

	docID, err := strconv.Atoi(ID)
	if err != nil {
		return updateReturn{nil, "error converting ID"}
	}

	myDB, err := db.OpenDB(dbAddress)
	if err != nil {
		return updateReturn{nil, "Can't open database"}
	}
	defer myDB.Close()

	curTable := myDB.Use(tableName)
	if curTable == nil {
		return updateReturn{nil, "Can't open table"}
	}

	readBack, err := curTable.Read(docID)
	if err != nil {
		return updateReturn{nil, "Can't find record"}
	}

	readBack["DocID"] = strconv.Itoa(docID)

	runActivities := false
	runLeads := false

	switch tableName {
	case "Account":
		runActivities = true
		runLeads = true
	case "Lead":
		runActivities = true
	}

	var query interface{}
	queryActivityResult := make(map[int]struct{})
	queryLeadResult := make(map[int]struct{})

	if runLeads {
		// add count of leads to NewResults
		// take docID and lookup in activity table
		leadTable := myDB.Use("Lead")
		if leadTable == nil {
			return updateReturn{nil, "Can't open lead sub table"}
		}

		json.Unmarshal([]byte(`[{"eq": "`+strconv.Itoa(docID)+`", "in": ["AccountID"]}]`), &query)

		if err := db.EvalQuery(query, leadTable, &queryLeadResult); err != nil {
			return updateReturn{nil, "Can't build lead query"}
		}

		readBack["LeadCount"] = strconv.Itoa(len(queryLeadResult))

	}

	if runActivities {
		// add count of activities to NewResults
		// take docID and lookup in activity table
		activityTable := myDB.Use("Activity")
		if activityTable == nil {
			return updateReturn{nil, "Can't open activty sub table"}
		}

		json.Unmarshal([]byte(`[{"eq": "`+strconv.Itoa(docID)+`", "in": ["ReferenceID"]}]`), &query)

		if err := db.EvalQuery(query, activityTable, &queryActivityResult); err != nil {
			return updateReturn{nil, "Can't build activity query"}
		}

		readBack["ActivityCount"] = strconv.Itoa(len(queryActivityResult))

	}

	for k, v := range readBack {
		readBack[k] = convertData("UPDATE", k, v)
	}

	return updateReturn{readBack, ""}
}
