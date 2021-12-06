package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jhonpedro/fullcycle-go/adapter/repository"
	"github.com/jhonpedro/fullcycle-go/services"
	"github.com/jhonpedro/fullcycle-go/usecase/process_transaction"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewTransactionRepositoryDb(db)

	unique_identifier_service := services.NewUniqueIdentifierService()

	usecase := process_transaction.NewProcessTransaction(repo, unique_identifier_service)

	input := process_transaction.TransactionInputDto{
		AccountId: "1",
		Amount:    2000,
	}

	output, err := usecase.Execute(input)

	if err != nil {
		fmt.Println(err.Error())
	}

	outputJson, _ := json.Marshal(output)

	fmt.Println(string(outputJson))
}
