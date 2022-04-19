package storage

type (
	// Storage wraps all functions
	// related to storaging.
	// e.g file, Redis, DynamoDB
	Storage interface {
		Put(v interface{}) error
		Get(v interface{}) error
		Update(v interface{}) error
		Delete(v interface{}) error
	}
)
