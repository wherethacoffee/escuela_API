package database

import (
	"context"

	"github.com/wherethacoffe/escuela_API/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
  
//Checks if a deposit exists in a depositsCollection  
func DepositExists(id primitive.ObjectID) (bool, error) {
    count, err := depositsCollection.CountDocuments(context.TODO(), bson.M{"_id": id})

    if count == 0 && err != nil {
	return false, err
    }

    return true, err
}

//CRUD METHODS FOR depositsCollection

//CREATE
func InsertDeposit(deposit models.Deposit) error {
    _, err := depositsCollection.InsertOne(context.TODO(),
    bson.D{{Key: "concepto", 
           Value: deposit.Concepto,},
   	   {Key: "cantidad",
	   Value: deposit.Cantidad},
  	   {Key: "fecha",
  	   Value: deposit.Fecha,},})
    if err != nil {
	return err
    }

    return nil
}

//READ
func GetDepositByID(id primitive.ObjectID) (models.Deposit, error) {
    var deposit models.Deposit
    res := depositsCollection.FindOne(context.TODO(), bson.M{"_id": id})

    if err := res.Err(); err != nil {
	return deposit, err
    }

    if err := res.Decode(&deposit); err != nil {
	return deposit, err
    }

    return deposit, nil
}

func GetAllDeposits() ([]models.Deposit, error) {
    var deposits []models.Deposit

    res, err := depositsCollection.Find(context.TODO(), bson.M{})

    if err != nil {
	return deposits, err
    }

    for res.Next(context.TODO()) {
	var deposit models.Deposit
	res.Decode(&deposit)
	deposits = append(deposits, deposit)
    }

    return deposits, err
}

//UPDATE
func UpdateDepositByID(id primitive.ObjectID, deposit models.Deposit) error {
    depositExists, err := DepositExists(id)
    if !depositExists && err != nil {
	return err
    }

    _, err = depositsCollection.UpdateOne(context.TODO(), 
         bson.M{"_id": id}, 
	 bson.M{"$set": bson.D{
 		{Key: "concepto", 
           	Value: deposit.Concepto,},
   	   	{Key: "cantidad",
	   	Value: deposit.Cantidad},
  	   	{Key: "fecha",
  	   	Value: deposit.Fecha,},
	 }})

    if err != nil {
	return err
    }

    return nil
}

//DELETE
func DeleteDepositByID(id primitive.ObjectID) error {
    depositExists, err := DepositExists(id)
    if !depositExists && err != nil {
	return err
    }
    _, err = depositsCollection.DeleteOne(context.TODO(), bson.M{"_id": id})

    if err != nil {
	return err
    }

    return err
}
