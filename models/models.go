package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Models struct {
	Id                 primitive.ObjectID `json:"_id",bson:"_id`
	Name               string             `json:"name",bson:"name"`
	Age                int                `json:"age",bson:"age"`
	EmailId            string             `json:"email",bson:"email"`
	Location           string             `json:"location",bson:"location"`
	EventParticipation string             `json:"eventparticipation",bson:"eventparticipation"`
	MembersAttending   int                `json:"membersattending",bson:"membersattending"`
}
