package helpers

import (
	"encoding/json"
	"log"
	"net/http"

	model "github.com/Mukunth-arya/Ticketbooking/models"
)

type LoggerValue struct {
	logger *log.Logger
}

func Datafunc(logger* log.Logger) *LoggerValue {

	return &LoggerValue{logger}
}
func (l *LoggerValue) Displayget(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("Content-Type", "application/x-www-form-urlencode")

	json.NewEncoder(rw).Encode(LocationGet())
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
func (l *LoggerValue) LivenessProbe(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("Content-Type", "application/x-www-form-urlencode")

	json.NewEncoder(rw).Encode("Hello welcome!!!!")

}
