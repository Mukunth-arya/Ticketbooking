package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Mukunth-arya/Ticketbooking/helpers"
	model "github.com/Mukunth-arya/Ticketbooking/models"
)

type LoggerValue struct {
	log.Logger
}

func Datafunc(Datas log.Logger) *LoggerValue {

	return &LoggerValue{Datas}
}
func (l *LoggerValue) Displayget(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("Content-Type", "application/x-www-form-urlencode")

	json.NewEncoder(rw).Encode(helpers.LocationGet())
}
func (l *LoggerValue) Dataenter(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("Content-Type", "application/x-www-form-urlencode")
	var Datagetvalue model.Models
	err := json.NewDecoder(r.Body).Decode(&Datagetvalue)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(rw).Encode(Datagetvalue)
}
