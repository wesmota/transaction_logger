package main

import (
	"github.com/gorilla/mux"
	"github.com/wesmota/transaction_logger/http/rest"
	"github.com/wesmota/transaction_logger/log"
	"github.com/wesmota/transaction_logger/store"
)

var h *rest.Handler

func init() {
	store := store.New()
	var logFile log.TransactionLogger
	logFile, err := log.NewFileTransactionLogger("transaction.log")
	if err != nil {
		panic(err)
	}
	h = rest.NewHandler(store, logFile.(*log.FileTransactionLogger))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/keys/{key}", h.Get).Methods("GET")

}
