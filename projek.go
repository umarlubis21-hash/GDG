package main

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

type User struct {
	ID 	 int    `json:"id"`
	Nama string `json:"name"`
	Role string `json:"role"` // untuk admin atau user biasa
}

type Event struct {
	ID 	  			int       `json:"id"`
	Title 			string    `json:"title"`
	KapasitasTotal 	int       `json:"kapasitas_total"`
	SisaStok      	int       `json:"sisa_stok"`
}

type Transaksi struct {
	ID       		int       `json:"id"`
	UserID   		int       `json:"user_id"`
	EventID  		int       `json:"event_id"`
	Jumlah 			int       `json:"jumlah"`
	Waktu 			time.Time `json:"waktu"`
}

var (
	users	  	= []User{}
	events	  	= []Event{}
	transaksis 	= []Transaksi{}

	mu sync.Mutex
)

func CreateUser (w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {return}
	var u User
	json.NewDecoder(r.Body).Decode(&u)
	u.ID = len(users) + 1
	users = append(users, u)
	json.NewEncoder(w).Encode(u)
}

func CreateEvent (w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {return}
	var e Event
	json.NewDecoder(r.Body).Decode(&e)
	e.ID = len(events) + 1
	events = append(events, e)
	json.NewEncoder(w).Encode(e)
}

func BeliTiket (w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {return}
	var req struct {
		UserID  int `json:"user_id"`
		EventID int `json:"event_id"`
		Jumlah  int `json:"jumlah"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	mu.Lock()
	defer mu.Lock()

	found := false
	for i := range events {
		if events[i].ID == req.EventID {
			found = true
			if events[i].SisaStok >= req.Jumlah {
				events[i].SisaStok -= req.Jumlah
				t := Transaksi{
					ID:	 len(transaksis) + 1,
					UserID: req.UserID,
					EventID: req.EventID,
					Jumlah: req.Jumlah,
					Waktu: time.Now(),
				}
				transaksis = append(transaksis, t)

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(map[string]string{"message": "Pembelian tiket berhasil"})
				return
			}
		}
	}
	if !found {
		http.Error(w, "Event tidak ditemukan", http.StatusNotFound)
	} else {
		http.Error(w, "Stok tiket tidak mencukupi", http.StatusBadRequest)
	}
}