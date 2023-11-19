package main

import (
	"errors"
)

var NoActiveTransactionError = errors.New("no active transaction")
var NotFoundError = errors.New("not found")

// Transaction acts as a storage for key-value items associated with that transactions
type Transaction = map[string]*string

type DB struct {
	transactions []Transaction // Stack of transactions
}

func NewDB() *DB {
	db := DB{}
	// Start the default transaction
	db.StartTransaction()
	return &db
}

// Get the value associated with the given key.
func (db *DB) Get(key string) (string, error) {
	// Iterate on transactions in a reverse order (from the active one to the default one)
	// and search for a specified key's value
	for i := len(db.transactions) - 1; i >= 0; i-- {
		val, ok := db.transactions[i][key]
		if ok {
			if val == nil {
				// key has been deleted
				return "", NotFoundError
			}
			return *val, nil
		}
	}
	return "", NotFoundError
}

// Set stores a key-value pair in the database.
func (db *DB) Set(key string, value string) {
	activeTransaction := db.peekTransaction()
	activeTransaction[key] = &value
}

// Delete removes the key-value pair associated with the given key.
func (db *DB) Delete(key string) {
	activeTransaction := db.peekTransaction()
	activeTransaction[key] = nil
}

// StartTransaction starts a new transaction. All operations within this transaction are isolated from others.
func (db *DB) StartTransaction() {
	newTransaction := make(Transaction)
	db.pushTransaction(newTransaction)
}

// Rollback reverts all changes made within the current transaction and discard them.
func (db *DB) Rollback() error {
	if len(db.transactions) == 1 {
		return NoActiveTransactionError
	}

	db.popTransaction()
	return nil
}

// Commit pushes all changes made within the current transaction to the database.
func (db *DB) Commit() error {
	if len(db.transactions) == 1 {
		return NoActiveTransactionError
	}

	activeTransaction := db.popTransaction()
	previousTransaction := db.peekTransaction()
	// Move all items from the active transaction to the previous one
	for key, value := range activeTransaction {
		previousTransaction[key] = value
		// TODO clean-up deleted items (nilled) if the previousTransaction is the default one
	}
	return nil
}

// pushTransaction put a new transaction onto the stack
func (db *DB) pushTransaction(storage Transaction) {
	db.transactions = append(db.transactions, storage) // Simply append the new value to the end of the stack
}

// popTransaction removes and return the top element of the stack
func (db *DB) popTransaction() Transaction {
	topMostIndex := len(db.transactions) - 1
	topMostElement := (db.transactions)[topMostIndex]
	// Remove the top element from the stack by slicing it off.
	db.transactions = (db.transactions)[:topMostIndex]
	return topMostElement
}

// peekTransaction returns (but doesn't remove) the top element of the stack
func (db *DB) peekTransaction() Transaction {
	topMostIndex := len(db.transactions) - 1
	topMostElement := (db.transactions)[topMostIndex]
	return topMostElement
}
