package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type CacheReq struct {
	Key        string      `json:"key"`
	Value      interface{} `json:"value"`
	Expiration int         `json:"expiration"` // in seconds
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get(KEY)
	value, found := cache.Get(key)
	if !found {
		http.Error(w, "Key not found or expired", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"key": key, "value": value})
}

func setHandler(w http.ResponseWriter, r *http.Request) {
	var req CacheReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	expiration := time.Duration(req.Expiration) * time.Second
	cache.Set(req.Key, req.Value, expiration)
	w.WriteHeader(http.StatusOK)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get(KEY)
	cache.Delete(key)
	w.WriteHeader(http.StatusOK)
}
