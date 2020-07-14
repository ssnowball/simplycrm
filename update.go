package main

import (
	"strconv"
	"time"

	"github.com/HouzuoGuo/tiedot/db"
)

func updateData(tableName string, data map[string]interface{}) ActionReturn {

	myDB, err := db.OpenDB(dbAddress)
	if err != nil {
		return ActionReturn{false, "Can't open database"}
	}
	defer myDB.Close()

	curTable := myDB.Use(tableName)
	if curTable == nil {
		return ActionReturn{false, "Can't open table"}
	}

	data["ModifiedDate"] = time.Now()
	data["ModifiedUsername"] = "admin"

	docID, _ := strconv.Atoi(data["DocID"].(string))

	layout := "2006-01-02T15:04:05.999999999Z07:00"
	pDs, _ := time.Parse(layout, data["CreatedDate"].(string))

	data["CreatedDate"] = pDs

	//get previous document information
	readBack, err := curTable.Read(docID)
	if err != nil {
		return ActionReturn{false, "Can't find item to update"}
	}

	err = curTable.Update(docID, data)
	if err != nil {
		return ActionReturn{false, "Can't update"}
	}

	AuditCreate(docID, tableName, "all", readBack, data, "UPDATE", "ADMIN")

	return ActionReturn{true, ""}
}
