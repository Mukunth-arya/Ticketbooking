package helpers

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"

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
func (h *HelperModels) CredPayment(Datavalue model.Payments) {

	Datavalue.Id = datas.Id
	Hashacc := h.Hashingacc(Datavalue.Accountnumber)
	Hashpasswd := h.HashingCvv(Datavalue.CVV)

	Data1 := Mongoconnect()

	Data2, err := Data1.InsertOne(context.Background(), model.Payments{Id: Datavalue.Id, Accountnumber: Hashacc, CVV: Hashpasswd})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted Document", Data2.InsertedID)

}
func (h *HelperModels) Hashingacc(accountnum string) string {
	value, err := bcrypt.GenerateFromPassword([]byte(accountnum), 14)
	if err != nil {
		log.Fatal(err)
	}
	return string(value)
}
func (h *HelperModels) HashingCvv(cvv string) string {

	value1, err := bcrypt.GenerateFromPassword([]byte(cvv), 14)
	if err != nil {
		log.Fatal(err)
	}
	return string(value1)
}
