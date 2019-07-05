package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	//Init Router
	router := mux.NewRouter()

	//Route Handler
	router.HandleFunc("/api/getCurrency/{date}/{base}/{rate}", getCurrency).Methods("GET")
	http.ListenAndServe(":5000", router)
}

func getCurrency(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//parametrelerin ne oldugunu cektik, daha sonra maps ile bu parametrelerin değerlerini aldık
	params := mux.Vars(r)
	//url taslagı oluşturdum
	var raw = "https://api.ratesapi.io/api/dateparam?base=baseacurr&symbols=ratecurr"
	//bu taslağın içerisinden tarihin bulundugu alanı değiştirdim
	newurl := strings.Replace(raw, "dateparam", params["date"], 1)
	//daha sonrasında güncellediğim url i parse edip değişmesi gereken alanları kullanıcının yolladıgı istek neticesinde değiştirdim
	u, err := url.Parse(newurl)
	if err != nil {
		log.Fatal(err)
	}
	//url içeriğindeki parametreleri değiştirdim
	q := u.Query()
	q.Set("base", params["base"])
	q.Set("symbols", params["rate"])
	u.RawQuery = q.Encode()

	//bu urlden verileri çektim
	resp, err2 := http.Get(u.String())
	if err2 != nil {
		fmt.Println(err.Error)
	} else {
		//hata yoksa bir Currency nesnesi oluşturup, gelen response un body kısmını önce byte array e sonrasında unmarshal ile de daha önceden tanımladıgım struct yapısına geri cevirdim
		var getCur Currency
		data, _ := ioutil.ReadAll(resp.Body)
		err3 := json.Unmarshal(data, &getCur)
		if err3 != nil {
			fmt.Println(err3.Error)
		}
		//son olarak bu struct ı json a çevirip return ettim
		json.NewEncoder(w).Encode(getCur)
		return
	}

}

//Response için model
// Bu struct kısmı değişmesi gerekiyor Verileri aldığımız formata uygun bir biçimde olmalı
type Currency struct {
	Base  string `json:"base"`
	Rates struct {
		GBP float64 `json:"GBP,omitempty"`
		HKD float64 `json:"HKD,omitempty"`
		IDR float64 `json:"IDR,omitempty"`
		ILS float64 `json:"ILS,omitempty"`
		DKK float64 `json:"DKK,omitempty"`
		INR float64 `json:"INR,omitempty"`
		CHF float64 `json:"CHF,omitempty"`
		MXN float64 `json:"MXN,omitempty"`
		CZK float64 `json:"CZK,omitempty"`
		SGD float64 `json:"SGD,omitempty"`
		THB float64 `json:"THB,omitempty"`
		HRK float64 `json:"HRK,omitempty"`
		MYR float64 `json:"MYR,omitempty"`
		NOK float64 `json:"NOK,omitempty"`
		CNY float64 `json:"CNY,omitempty"`
		BGN float64 `json:"BGN,omitempty"`
		PHP float64 `json:"PHP,omitempty"`
		SEK float64 `json:"SEK,omitempty"`
		PLN float64 `json:"PLN,omitempty"`
		ZAR float64 `json:"ZAR,omitempty"`
		CAD float64 `json:"CAD,omitempty"`
		ISK float64 `json:"ISK,omitempty"`
		BRL float64 `json:"BRL,omitempty"`
		RON float64 `json:"RON,omitempty"`
		NZD float64 `json:"NZD,omitempty"`
		TRY float64 `json:"TRY,omitempty"`
		JPY float64 `json:"JPY,omitempty"`
		RUB float64 `json:"RUB,omitempty"`
		KRW float64 `json:"KRW,omitempty"`
		USD float64 `json:"USD,omitempty"`
		HUF float64 `json:"HUF,omitempty"`
		AUD float64 `json:"AUD,omitempty"`
	} `json:"rates"`
	Date string `json:"date"`
}
