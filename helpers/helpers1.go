package helpers

import (
	"context"
	"fmt"
	"log"

	model "github.com/Mukunth-arya/Ticketbooking/models"
)

type HelperModels struct{}

func LocationGet() map[string]string {

	var loc = map[string]string{
		"INDIA":   "BANGLORE",
		"CHINA":   "HONG-KONG",
		"UAE":     "DUBAI",
		"ENGLANG": "LONDON",
	}

	return loc
}
func (h *HelperModels) Datasenter(datas model.Models) {
	Data1 := Mongoconnect()
	Data2, err := Data1.InsertOne(context.Background(), datas)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Document is Inserted", Data2.InsertedID)
}
