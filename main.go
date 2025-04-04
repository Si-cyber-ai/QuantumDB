package main

import (
	"fmt"
	"QuantumDB/executor"
	"QuantumDB/index"
	"QuantumDB/parser"
	"QuantumDB/storage"
	"QuantumDB/transaction"
)

func main() {
	// Create Table
	db, _ := storage.CreateTable("users")

	// Insert Data
	executor.ExecuteCommand(parser.ParseCommand("INSERT | users | Alice"), db)
	executor.ExecuteCommand(parser.ParseCommand("INSERT | users | Bob"), db)

	// Secure Transaction
	transaction.SecureTransaction("Alice sending $100 to Bob")

	// Quantum Search Optimization
	record, found := executor.QuantumSearch(db.Index.Data, "Alice")
	if found {
		fmt.Println("Quantum Optimized Lookup:", record)
	} else {
		fmt.Println("User Not Found")
	}
}
