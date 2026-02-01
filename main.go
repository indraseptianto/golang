package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// ===== MODEL =====
type Produk struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Harga int    `json:"harga"`
	Stok  int    `json:"stok"`
}

// ===== DATA (IN-MEMORY) =====
var produk = []Produk{
	{ID: 1, Nama: "indomie goreng", Harga: 3500, Stok: 30},
	{ID: 2, Nama: "indomie rebus", Harga: 4000, Stok: 30},
	{ID: 3, Nama: "indomie kuah", Harga: 2000, Stok: 10},
}

func main() {
	// API
	http.HandleFunc("/api/produk", produkHandler)
	http.HandleFunc("/api/produk/", produkHandler)
	http.HandleFunc("/health", healthHandler)

	fmt.Println("Server running di http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}

// ===== HANDLER PRODUK =====
func produkHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	path := strings.TrimPrefix(r.URL.Path, "/api/produk")

	// =====================
	// /api/produk
	// =====================
	if path == "" || path == "/" {
		switch r.Method {

		case http.MethodGet:
			json.NewEncoder(w).Encode(produk)

		case http.MethodPost:
			var produkBaru Produk
			err := json.NewDecoder(r.Body).Decode(&produkBaru)
			if err != nil {
				http.Error(w, "Invalid request body", http.StatusBadRequest)
				return
			}

			if produkBaru.Nama == "" || produkBaru.Harga <= 0 || produkBaru.Stok < 0 {
				http.Error(w, "Data produk tidak valid", http.StatusBadRequest)
				return
			}

			produkBaru.ID = len(produk) + 1
			produk = append(produk, produkBaru)

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(produkBaru)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
		return
	}

	// =====================
	// /api/produk/{id}
	// =====================
	idStr := strings.TrimPrefix(path, "/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	switch r.Method {

	// ===== GET BY ID =====
	case http.MethodGet:
		for _, p := range produk {
			if p.ID == id {
				json.NewEncoder(w).Encode(p)
				return
			}
		}
		http.Error(w, "Produk tidak ditemukan", http.StatusNotFound)

	// ===== UPDATE =====
	case http.MethodPut:
		var updated Produk
		err := json.NewDecoder(r.Body).Decode(&updated)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		for i, p := range produk {
			if p.ID == id {
				updated.ID = id
				produk[i] = updated
				json.NewEncoder(w).Encode(updated)
				return
			}
		}
		http.Error(w, "Produk tidak ditemukan", http.StatusNotFound)

	// ===== DELETE =====
	case http.MethodDelete:
		for i, p := range produk {
			if p.ID == id {
				deleted := p
				produk = append(produk[:i], produk[i+1:]...)

				json.NewEncoder(w).Encode(map[string]interface{}{
					"message": "Produk berhasil dihapus",
					"data":    deleted,
				})
				return
			}
		}
		http.Error(w, "Produk tidak ditemukan", http.StatusNotFound)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// ===== HEALTH =====
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"message": "api running",
	})
}
