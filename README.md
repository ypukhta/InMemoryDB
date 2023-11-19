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
- On the top of the stack we always have an active (current) transaction.
- Transaction stack implementation is backed by a slice. Alternatively, a linked list could be used.
- Changes (set and delete operations) made in the context of an active transaction are stored in a data storage associated with a transaction. 
- Data storage associated with a transaction is map where value is a string pointer rather than a string. A pointer type allows to use a nil to depict a deleted key.
- On transaction commit all changes made in the context of an active transaction are propagated to the previous one. And then the transaction is popped out from the stack.
- On transaction rollback no changes are propagated, an active transaction is just popped out from the stack.


## Requirements:

- Go version 1.20+

## Run:

- To run, type `go run .`


## Test the project

- To test, type `go test ./...` 
