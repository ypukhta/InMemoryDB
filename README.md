# In-memory database

In-memory key-value database with nested transaction support.

Operations:
- get(key)
- set(key, value)
- delete(key)
- start_transaction()
- commit()
- roll_back()

## Approach
- A FIFO stack data structure is used to store nested transactions.
- The top of the stack always holds the active (current) transaction.
- The implementation of the transaction stack is backed by a slice. Alternatively, a linked list could be used.
- Changes (set and delete operations) made within the context of an active transaction are stored in a data storage associated with that transaction.
- The data storage associated with a transaction is a map where the value is a string pointer rather than a string itself. A pointer type allows us to use a nil value to represent a deleted key.
- Upon transaction commit, all changes made within the context of the active transaction are propagated to the previous transaction in the stack. Then, the committed transaction is popped from the stack.
- Upon transaction rollback, no changes are propagated, the active transaction is simply popped from the stack.

## Requirements:

- Go version 1.20+

## Run:

- To run, type `go run .`


## Test the project

- To test, type `go test ./...` 
