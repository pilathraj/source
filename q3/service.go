package q3

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pilathraj/talentpro/helper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//Connection mongoDB with helper class
var collection = helper.GetCollection()

type User struct {
	ID       primitive.ObjectID `json:"Id,omitempty" bson:"_id,omitempty"`
	LoginID  string             `json:"loginId" bson:"loginId"`
	FullName string             `json:"fullName" bson:"fullName"`
	Enable   bool               `json:"enable" bson:"enable"`
}

//GetUsers - return user list
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{}
	filter := bson.D{{}} // all users
	cur, findError := collection.Find(context.TODO(), filter)

	if findError != nil {
		helper.GetError(findError, w)
		return
	}
	// once exhausted, close the cursor
	defer cur.Close(context.TODO())

	//Map result to slice
	for cur.Next(context.TODO()) {
		u := User{}
		err := cur.Decode(&u)
		if err != nil {
			helper.GetError(err, w)
			return
		}
		users = append(users, u)
	}

	if len(users) == 0 {
		helper.GetError(mongo.ErrNoDocuments, w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

//GetUser - return single user
func GetUser(w http.ResponseWriter, r *http.Request) {
	var user User
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)

	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(user)

}

//DeleteUser - remove user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// string to primitve.ObjectID
	id, err := primitive.ObjectIDFromHex(params["id"])

	// prepare filter.
	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(deleteResult)

}

//CreateUser - creata new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var item User
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		helper.GetError(err, w)
		return
	}
	_, err = collection.InsertOne(context.TODO(), item)
	if err != nil {
		helper.GetError(err, w)
		return
	}
	json.NewEncoder(w).Encode(item)
}

//UpdateUser - Update user detail
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	var params = mux.Vars(r)

	//Get id from parameters
	id, _ := primitive.ObjectIDFromHex(params["id"])

	// Create filter
	filter := bson.M{"_id": id}

	// Read update model from body request
	_ = json.NewDecoder(r.Body).Decode(&user)

	// prepare update model.
	update := bson.D{
		{"$set", bson.D{
			{"loginId", user.LoginID},
			{"fullName", user.FullName},
			{"enable", user.Enable},
		}},
	}

	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&user)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	user.ID = id

	json.NewEncoder(w).Encode(user)

}
