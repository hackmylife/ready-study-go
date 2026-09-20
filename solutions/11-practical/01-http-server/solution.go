//go:build ignore

package koan

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type Note struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

func NewHandler() http.Handler {
	var mu sync.Mutex
	notes := make(map[int]Note)
	nextID := 1
	mux := http.NewServeMux()
	mux.HandleFunc("POST /notes", func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 4096)
		var req struct {
			Text string `json:"text"`
		}
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&req); err != nil || strings.TrimSpace(req.Text) == "" {
			http.Error(w, "invalid note", http.StatusBadRequest)
			return
		}
		var extra any
		if err := decoder.Decode(&extra); err != io.EOF {
			http.Error(w, "invalid note", http.StatusBadRequest)
			return
		}
		mu.Lock()
		note := Note{ID: nextID, Text: req.Text}
		notes[nextID] = note
		nextID++
		mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Location", fmt.Sprintf("/notes/%d", note.ID))
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(note)
	})
	mux.HandleFunc("GET /notes/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		mu.Lock()
		note, ok := notes[id]
		mu.Unlock()
		if !ok {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(note)
	})
	return mux
}
