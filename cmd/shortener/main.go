package main

import (
	"encoding/base64"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

var dBEmul map[string]string

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/{id}", returnUrlById)
	myRouter.HandleFunc("/", shortLongUrl).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func returnUrlById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	baseUrl := dBEmul[key]
	w.Header().Set("Location", baseUrl)
	w.WriteHeader(307)
	w.Write([]byte(baseUrl))
}

func shortLongUrl(w http.ResponseWriter, r *http.Request) {
	d, err := io.ReadAll(r.Body)
	// обрабатываем ошибку
	if err != nil || string([]byte(d)) == "" {
		http.Error(w, err.Error(), 400)
		return
	}
	shortUrl := UrlShortener(string(d))
	dBEmul[shortUrl] = string(d)
	w.Write([]byte(shortUrl))
	w.WriteHeader(201)
}

func UrlShortener(url string) string {
	shortUrl := []byte(url)
	urlEnc := base64.StdEncoding.EncodeToString(shortUrl)
	return urlEnc
}

func main() {
	dBEmul = make(map[string]string)
	handleRequests()
	http.ListenAndServe(":8080", nil)
}
