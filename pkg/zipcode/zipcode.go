package zipcode

import (
	"database/sql"
	"fmt"
	"net/http"
)

const defautlAddr = ":8000"

func New() (*http.Server, error) {
	mux, err := newMux()
	if err != nil {
		return nil, err
	}
	return &http.Server{
		Addr:    defautlAddr,
		Handler: mux,
	}, nil
}

func newMux() (*http.ServeMux, error) {
	h, err := newHandler()
	if err != nil {
		return nil, err
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/reset", h.reset)
	mux.HandleFunc("/orders", h.orders)
	mux.HandleFunc("/cities", h.cities)
	return mux, nil
}

type handler struct {
	db *sql.DB
}

func newHandler() (*handler, error) {
	db, err := newPostgreSQLDB()
	if err != nil {
		return nil, err
	}
	return &handler{db: db}, nil
}

func (h *handler) reset(w http.ResponseWriter, r *http.Request) {
	if err := resetOrders(r.Context(), h.db); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *handler) orders(w http.ResponseWriter, r *http.Request) {
	zips := r.URL.Query()["zip"]
	if len(zips) != 0 {
		h.insertOrders(w, r, zips)
	} else {
		h.getOrders(w, r)
	}
}

func (h *handler) getOrders(w http.ResponseWriter, r *http.Request) {
	os, err := listOrders(r.Context(), h.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	for _, o := range os {
		fmt.Fprintln(w, o)
	}
}

func (h *handler) insertOrders(w http.ResponseWriter, r *http.Request, zips []string) {
	err := insertOrders(r.Context(), h.db, zips...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.getOrders(w, r)
}

func (h *handler) cities(w http.ResponseWriter, r *http.Request) {
	cities, err := orderCountByCity(r.Context(), h.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	for city, count := range cities {
		fmt.Fprintf(w, "%d Order in %s\n", count, city)
	}
}
