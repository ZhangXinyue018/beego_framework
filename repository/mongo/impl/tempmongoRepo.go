package impl

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TempMongoRepository struct {
	TempMongoSession   *mgo.Session
	TempDBName         string
	TempCollectionName string
}

type TestMongo struct {
	Id_ bson.ObjectId `bson:"_id"`
}

func (repository *TempMongoRepository) Get() (*TestMongo) {
	session := repository.TempMongoSession.Clone()
	defer session.Close()
	collection := repository.TempMongoSession.DB(repository.TempDBName).C(repository.TempCollectionName)
	result := &TestMongo{}
	collection.Find(bson.M{"test": "test"}).One(result)
	return result
}
