package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Deposit struct{
    Id primitive.ObjectID `json:"_id" bson:"_id"`
    Concepto string `json:"concepto" bson:"concepto"`
    Cantidad float64 `json:"cantidad" bson:"cantidad"`
    Fecha string `json:"fecha" bson:"fecha"`
}
