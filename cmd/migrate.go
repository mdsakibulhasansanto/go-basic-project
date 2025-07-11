package main

import (
	"final-project/db"
	"final-project/models"
	"flag"
	"fmt"
)

func main() {
	gormDb, err := db.NewMySqlGormDB()
	if err != nil {
		panic(err)
	}

	only := flag.String("only", "", "Run migration for a specific model: users, products")
	flag.Parse()

	switch *only {
	case "users":
		err := gormDb.AutoMigrate(&models.User{})
		if err != nil {
			panic(err)
		}
		fmt.Println("Users table migrated.")
	case "products":
		err := gormDb.AutoMigrate(&models.Product{})
		if err != nil {
			panic(err)
		}
		fmt.Println("Products table migrated.")
	case "":
		if err := db.RunMigration(gormDb); err != nil {
			panic(err)
		}
		fmt.Println("All tables migrated.")
	default:
		fmt.Println("Unknown migration type. Use: users or products")
	}
}
