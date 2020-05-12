package models

import (
	"github.com/jaydeep87/poc-go-microservice/src/mgo-microservices/db"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID    bson.ObjectId `bson:"_id" json:"id"`
	Name  string `bson:"name" json:"name"`
	Email string `bson:"email" json:"email"`
}

type UserModel struct{}

var server = "mongodb://localhost:27017/"

var dbConnect = db.NewConnection(server)

func (m *UserModel) Create(data User) error {
	collection := dbConnect.Use("test-mgo", "users")
	err := collection.Insert(bson.M{"name": data.Name, "email": data.Email})
	return err
}

func (m *UserModel) Find() (list []User, err error) {
	collection := dbConnect.Use("test-mgo", "users")
	err = collection.Find(bson.M{}).All(&list)
	return list, err
}

func (m *UserModel) Get(id string) (user User, err error) {
	collection := dbConnect.Use("test-mgo", "users")
	err = collection.FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

func (m *UserModel) Update(id string, data User) (err error) {
	collection := dbConnect.Use("test-mgo", "users")
	err = collection.UpdateId(bson.ObjectIdHex(id), data)

	return err
}

func (m *UserModel) Delete(id string) (err error) {
	collection := dbConnect.Use("test-mgo", "users")
	err = collection.RemoveId(bson.ObjectIdHex(id))

	return err
}
