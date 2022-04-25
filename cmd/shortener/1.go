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
	myRouter.HandleFunc("/{id}", returnURLByID)
	myRouter.HandleFunc("/", shortLongURL).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func returnURLByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	baseURL := dBEmul[key]
	w.Header().Set("Location", baseURL)
	w.WriteHeader(307)
	w.Write([]byte(baseURL))
}

func shortLongURL(w http.ResponseWriter, r *http.Request) {
	d, err := io.ReadAll(r.Body)
	// обрабатываем ошибку
	if err != nil || string([]byte(d)) == "" {
		http.Error(w, err.Error(), 400)
		return
	}
	shortURL := URLShortener(string(d))
	dBEmul[shortURL] = string(d)
	w.Write([]byte(shortURL))
	w.WriteHeader(201)
}

func URLShortener(url string) string {
	shortURL := []byte(url)
	urlEnc := base64.StdEncoding.EncodeToString(shortURL)
	return urlEnc
}

func main() {
	dBEmul = make(map[string]string)
	handleRequests()
	http.ListenAndServe(":8080", nil)
}
