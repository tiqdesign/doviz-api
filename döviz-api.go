package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//Response için model
// Bu struct kısmı değişmesi gerekiyor Verileri aldığımız formata uygun bir biçimde olmalı
type Currency struct {
	ID   string `json:"id"`
	Curr string `json:"currency"`
	Date string `json:"date"`
}

// burada sıkıntı var ben array in indexini veremem onun yerine slice a bakmalıyım
// daha sonrasında verilen tarihe göre önce hepsini cekip istediği dövizi bu slice icerisinden vermem gerek
var currencyResp [3]Currency

func main() {

	//Init Router
	router := mux.NewRouter()

	//Route Handler
	router.HandleFunc("/api/getCurrency/{date}/{currency}", getCurrency).Methods("GET")
	http.ListenAndServe(":8000", router)
}

func getCurrency(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//parametrelerin ne oldugunu cektik, daha sonra maps ile bu parametrelerin değerlerini aldık
	params := mux.Vars(r)
	//c# taki foreach gibi (_ index değeri oluyor array ın kacıncı elemanda oldugunu gösteriyor)
	for _, item := range currencyResp {
		if item.Curr == params["currency"] && item.Date == params["date"] {
			//Eğer sağlanıyorsa elindeki struct ı json formatına cevir ve return et
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}
