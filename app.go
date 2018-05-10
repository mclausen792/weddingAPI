package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	. "weddingAPI/config"
	. "weddingAPI/dao"
	. "weddingAPI/models"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var config = Config{}
var dao = GuestsDAO{}

func AllGuestsEndPoint(w http.ResponseWriter, r *http.Request) {
	guests, err := dao.FindAllGuests()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, guests)
}

func CreateGuestEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var guest Guests
	if err := json.NewDecoder(r.Body).Decode(&guest); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	guest.ID = bson.NewObjectId()
	if err := dao.Insert(guest); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, guest)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func init() {
	config.Read()
	dao.DialInfo = &mgo.DialInfo{
		Addrs:    []string{config.Server},
		Database: config.Database,
		Username: config.Username,
		Password: config.Password,
	}
	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	r := mux.NewRouter()
	r.HandleFunc("/guests", AllGuestsEndPoint).Methods("GET")
	r.HandleFunc("/guests", CreateGuestEndPoint).Methods("POST")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"hello\": \"world\"}"))
	})

	fmt.Println("running on Port" + port)
	handler := cors.Default().Handler(r)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal(err)
	}
}
