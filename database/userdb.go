package database

import (
	"context"

	"github.com/wherethacoffe/escuela_API/models"
	"go.mongodb.org/mongo-driver/bson"
)

//Checks if a user exists in a usersCollection  
func DocumentExists(username string) (bool, error) {
    count, err := usersCollection.CountDocuments(context.TODO(), bson.M{"username": username})

    if count == 0 && err != nil {
	return false, err
    }

    return true, err
}

//CRUD METHODS FOR usersCollection

//CREATE
func InsertUser(user models.User) error {
    _, err := usersCollection.InsertOne(context.TODO(),
    bson.D{{Key: "username", 
           Value: user.Username,},
   	   {Key: "pwd",
	   Value: user.Password},
  	   {Key: "isadmin",
  	   Value: user.IsAdmin,}})
    if err != nil {
	return err
    }

    return nil
}

//READ
func GetUserByUsername(username string) (models.User, error) {
    var user models.User
    res := usersCollection.FindOne(context.TODO(), bson.M{"username": username})

    if err := res.Err(); err != nil {
	return user, err
    }

    if err := res.Decode(&user); err != nil {
	return user, err
    }

    return user, nil
}

func GetAllUsers() ([]models.User, error) {
    var users []models.User

    res, err := usersCollection.Find(context.TODO(), bson.M{})

    if err != nil {
	return users, err
    }

    for res.Next(context.TODO()) {
	var user models.User
	res.Decode(&user)
	users = append(users, user)
    }

    return users, err
}

//UPDATE
func UpdateUserByUsername(username string, user models.User) error {
    userExists, err := DocumentExists(username)
    if !userExists && err != nil {
	return err
    }

    _, err = usersCollection.UpdateOne(context.TODO(), 
         bson.M{"username": username}, 
	 bson.M{"$set": bson.D{
 		{Key: "username", 
           	Value: user.Username,},
   	   	{Key: "pwd",
	   	Value: user.Password},
  	   	{Key: "isadmin",
  	   	Value: user.IsAdmin,},
	 }})

    if err != nil {
	return err
    }

    return nil
}

//DELETE
func DeleteUserByUsername(username string) error {
    userExists, err := DocumentExists(username)
    if !userExists && err != nil {
	return err
    }
    _, err = usersCollection.DeleteOne(context.TODO(), bson.M{"username": username})

    if err != nil {
	return err
    }

    return err
}
