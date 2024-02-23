package database

import (
	"context"

	"github.com/wherethacoffe/escuela_API/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
  
//Checks if a record exists in a recordsCollection  
func RecordExists(id primitive.ObjectID) (bool, error) {
    count, err := recordsCollection.CountDocuments(context.TODO(), bson.M{"_id": id})

    if count == 0 && err != nil {
	return false, err
    }

    return true, err
}

//CRUD METHODS FOR recordsCollection

//CREATE
func InsertRecord(record models.Record) error {
    _, err := recordsCollection.InsertOne(context.TODO(),
    bson.D{{Key: "cuatrimestre", 
           Value: record.Cuatrimestre,},
   	   {Key: "mes",
	   Value: record.Mes},
  	   {Key: "colegiatura",
  	   Value: record.Colegiatura,},
  	   {Key: "extra",
  	   Value: record.Extra,},
     	   {Key: "total",
  	   Value: record.Total,},})
    if err != nil {
	return err
    }

    return nil
}

//READ
func GetRecordByID(id primitive.ObjectID) (models.Record, error) {
    var record models.Record
    res := recordsCollection.FindOne(context.TODO(), bson.M{"_id": id})

    if err := res.Err(); err != nil {
	return record, err
    }

    if err := res.Decode(&record); err != nil {
	return record, err
    }

    return record, nil
}

func GetAllRecords() ([]models.Record, error) {
    var records []models.Record

    res, err := recordsCollection.Find(context.TODO(), bson.M{})

    if err != nil {
	return records, err
    }

    for res.Next(context.TODO()) {
	var record models.Record
	res.Decode(&record)
	records = append(records, record)
    }

    return records, err
}

//UPDATE
func UpdateRecordByID(id primitive.ObjectID, record models.Record) error {
    recordExists, err := RecordExists(id)
    if !recordExists && err != nil {
	return err
    }

    _, err = recordsCollection.UpdateOne(context.TODO(), 
         bson.M{"_id": id}, 
	 bson.M{"$set": bson.D{
 		{Key: "cuatrimestre", 
           	Value: record.Cuatrimestre,},
   	   	{Key: "mes",
	   	Value: record.Mes},
  	   	{Key: "colegiatura",
  	   	Value: record.Colegiatura,},
  	   	{Key: "extra",
  	   	Value: record.Extra,},
	     	{Key: "total",
  	   	Value: record.Total,},
	 }})

    if err != nil {
	return err
    }

    return nil
}

//DELETE
func DeleteRecordByID(id primitive.ObjectID) error {
    recordExists, err := RecordExists(id)
    if !recordExists && err != nil {
	return err
    }
    _, err = recordsCollection.DeleteOne(context.TODO(), bson.M{"_id": id})

    if err != nil {
	return err
    }

    return err
}
