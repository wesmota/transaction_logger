package log

// TransactionLogger is an interface for logging transactions for PUT and DELETE events.
type TransactionLogger interface {
	WriteDelete(key string)
	WritePut(key string, value string)
}
