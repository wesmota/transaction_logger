package rest

import (
	"github.com/wesmota/transaction_logger/log"
	"github.com/wesmota/transaction_logger/store"
)

type Handler struct {
	store *store.Store
	log   *log.Log
}

func NewHandler(store *store.Store, log *log.Log) *Handler {
	return &Handler{
		store: store,
		log:   log,
	}
}
