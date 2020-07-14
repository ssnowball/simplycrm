package main

import (
	"strconv"

	"github.com/HouzuoGuo/tiedot/db"
	"github.com/HouzuoGuo/tiedot/dberr"
)

func deleteData(tableName string, ID string) ActionReturn {

	docID, err := strconv.Atoi(ID)
	if err != nil {
		return ActionReturn{false, "error converting ID"}
	}

	myDB, err := db.OpenDB(dbAddress)
	if err != nil {
		return ActionReturn{false, "Can't open database"}
	}
	defer myDB.Close()

	curTable := myDB.Use(tableName)
	if curTable == nil {
		return ActionReturn{false, "Can't open table"}
	}

	//get previous document information
	readBack, err := curTable.Read(docID)
	if err != nil {
		return ActionReturn{false, "Can't find item to delete"}
	}

	if err := curTable.Delete(docID); dberr.Type(err) == dberr.ErrorNoDoc {
		return ActionReturn{false, "Can't delete item"}
	}

	preDataIn := make(map[string]interface{})
	preDataIn["changed"] = "none"
	AuditCreate(docID, tableName, "all", readBack, preDataIn, "DELETE", "ADMIN")

	return ActionReturn{true, ""}
}
