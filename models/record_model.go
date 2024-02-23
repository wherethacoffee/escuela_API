package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Record struct{
    Id primitive.ObjectID `json:"_id" bson:"_id"`
    Cuatrimestre int `json:"cuatrimestre" bson:"cuatrimestre"`
    Mes string `json:"mes" bson:"mes"`
    Colegiatura float64 `json:"colegiatura" bson:"colegiatura"`
    Extra float64 `json:"extra" bson:"extra"`
    Total float64 `json:"total" bson:"total"`
}
