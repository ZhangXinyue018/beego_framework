package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TestMongoRepository struct {
	TestMongoSession   *mgo.Session
	TestDBName         string
	TestCollectionName string
}

type TestMongo struct {
	Id_ bson.ObjectId `bson:"_id"`
}

func (repository *TestMongoRepository) Get() (*TestMongo) {
	session := repository.TestMongoSession.Clone()
	defer session.Close()
	collection := repository.TestMongoSession.DB(repository.TestDBName).C(repository.TestCollectionName)
	result := &TestMongo{}
	collection.Find(bson.M{"test": "test"}).One(result)
	return result
}
