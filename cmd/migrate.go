package main

import (
	"final-project/db"
	"fmt"
)

func main() {

	gormDb, error := db.NewMySqlGormDB()

	if error != nil {
		panic(error)
	}
	if err := db.MigrateProductTable(gormDb); err != nil {

		panic(err)
	}

	fmt.Println("Table crate")

}
