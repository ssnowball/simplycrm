package main

import (
	"os"

	"github.com/HouzuoGuo/tiedot/db"
)

const dbAddress string = "doNotDelete.db"

func databaseBuild() ActionReturn {

	// create database
	if _, err := os.Stat(dbAddress); !os.IsNotExist(err) {
		return ActionReturn{true, "Database Exists"}
	}

	// create database
	myDB, err := db.OpenDB(dbAddress)
	if err != nil {
		return ActionReturn{false, "Can't open database"}
	}

	// create tables
	err = myDB.Create("Account")
	if err != nil {
		return ActionReturn{false, "Can't create table 'Account'"}
	}

	err = myDB.Create("Activity")
	if err != nil {
		return ActionReturn{false, "Can't create table 'Activity'"}
	}

	err = myDB.Create("Audit")
	if err != nil {
		return ActionReturn{false, "Can't create table 'Audit'"}
	}

	err = myDB.Create("Campaign")
	if err != nil {
		return ActionReturn{false, "Can't create table 'Campaign'"}
	}

	err = myDB.Create("Lead")
	if err != nil {
		return ActionReturn{false, "Can't create table 'Lead'"}
	}

	err = myDB.Create("Product")
	if err != nil {
		return ActionReturn{false, "Can't create table 'Product'"}
	}

	err = myDB.Create("Option")
	if err != nil {
		return ActionReturn{false, "Can't create table 'Option'"}
	}

	// create indexs
	var curTable *db.Col

	curTable = myDB.Use("Activity")
	if curTable == nil {
		return ActionReturn{false, "couldn't open 'activity' table"}
	}

	err = curTable.Index([]string{"ReferenceID"})
	if err != nil {
		return ActionReturn{false, "Can't create index 'ReferenceID'"}
	}

	curTable = myDB.Use("Lead")
	if curTable == nil {
		return ActionReturn{false, "couldn't open 'lead' table"}
	}

	err = curTable.Index([]string{"AccountID"})
	if err != nil {
		return ActionReturn{false, "Can't create index 'AccountID'"}
	}

	return ActionReturn{true, ""}

}
