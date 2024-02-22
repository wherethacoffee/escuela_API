package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct{
    Id primitive.ObjectID `json:"_id" bson:"_id"`
    Username string `json:"username" bson:"username"`
    Password string `json:"pwd" bson:"pwd"`
    IsAdmin bool `json:"isadmin" bson:"isadmin"`
}
