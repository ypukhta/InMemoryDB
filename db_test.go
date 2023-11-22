package main

import "testing"
import "github.com/stretchr/testify/assert"

// Example 1 for commit a transaction
func TestTransactionCommit(t *testing.T) {
	db := NewDB()
	db.Set("key1", "value1")
	db.StartTransaction()
	db.Set("key1", "value2")
	_ = db.Commit()
	val, err := db.Get("key1")
	assert.NoError(t, err)
	assert.Equal(t, "value2", val)
}

// Example 2 for roll_back().
func TestTransactionRollback(t *testing.T) {
	db := NewDB()
	db.Set("key1", "value1")

	db.StartTransaction()
	val, err := db.Get("key1")
	assert.NoError(t, err)
	assert.Equal(t, "value1", val)
	db.Set("key1", "value2")
	val, err = db.Get("key1")
	assert.NoError(t, err)
	assert.Equal(t, "value2", val)
	_ = db.Rollback()

	val, err = db.Get("key1")
	assert.NoError(t, err)
	assert.Equal(t, "value1", val)
}

// Example 3 for nested transactions
func TestNestedTransactions(t *testing.T) {
	db := NewDB()
	db.Set("key1", "value1")

	db.StartTransaction()
	db.Set("key1", "value2")
	val, err := db.Get("key1")
	assert.NoError(t, err)
	assert.Equal(t, "value2", val)

	db.StartTransaction()
	val, err = db.Get("key1")
	assert.NoError(t, err)
	assert.Equal(t, "value2", val)
	db.Delete("key1")
	_ = db.Commit()

	val, err = db.Get("key1")
	assert.Error(t, err)
	_ = db.Commit()

	val, err = db.Get("key1")
	assert.Error(t, err)
}

// Example 4 for nested transactions with roll_back()
func TestNestedTransactionsWithRollback(t *testing.T) {
	db := NewDB()
	db.Set("key1", "value1")

	db.StartTransaction()
	db.Set("key1", "value2")
	val, err := db.Get("key1")
	assert.NoError(t, err)
	assert.Equal(t, "value2", val)

	db.StartTransaction()
	val, err = db.Get("key1")
	assert.NoError(t, err)
	assert.Equal(t, "value2", val)
	db.Delete("key1")
	_ = db.Rollback()

	assert.NoError(t, err)
	assert.Equal(t, "value2", val)
	_ = db.Commit()

	val, err = db.Get("key1")
	assert.NoError(t, err)
	assert.Equal(t, "value2", val)
}

func TestGetCanReturnValueForUncommittedTTransaction(t *testing.T) {
	db := NewDB()

	db.Set("key1", "value1")

	// Case 1 - Get value after set for an uncommitted transaction
	db.StartTransaction()
	db.Set("key1", "value2")
	val, err := db.Get("key1")
	assert.NoError(t, err)
	assert.Equal(t, "value2", val)

	// Case 2 - Get value after delete for an uncommitted transaction
	db.StartTransaction()
	db.Delete("key1")
	val, err = db.Get("key1")
	assert.Error(t, err)
}
