package rest

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wesmota/transaction_logger/log"
	"github.com/wesmota/transaction_logger/store"
)

type Handler struct {
	store *store.Store
	log   *log.FileTransactionLogger
}

func NewHandler(store *store.Store, log *log.FileTransactionLogger) *Handler {
	return &Handler{
		store: store,
		log:   log,
	}
}

func (h *Handler) GetKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	value, found := h.store.Get(key)
	if !found {
		http.Error(w, errors.New("key not found").Error(), http.StatusNotFound)
	}
	item := Item{
		Key:   key,
		Value: value,
	}
	res, err := json.Marshal(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (h *Handler) PutKey(w http.ResponseWriter, r *http.Request) {
	var item Item
	req, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = json.Unmarshal(req, &item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	h.store.Put(item.Key, item.Value)
	h.log.WritePut(item.Key, item.Value)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DeleteKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	h.store.Delete(key)
	h.log.WriteDelete(key)
	w.WriteHeader(http.StatusOK)
}
