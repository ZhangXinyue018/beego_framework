package impl

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/dbtest"
	"os"
	"testing"
)

type TestMongoServer struct {
	dbtest.DBServer
}

var (
	MongoPath = "./mongodb/"
	//MongoTransactionCollection = "transaction"
)

func (server *TestMongoServer) StartMongo(t *testing.T) (*mgo.Session) {
	defer func() {
		if x := recover(); x != nil {
			t.Fatalf("Error: %s", x.(error).Error())
			server.Stop()
		}
	}()
	os.RemoveAll(MongoPath)
	if _, err := os.Stat(MongoPath); os.IsNotExist(err) {
		err = os.MkdirAll(MongoPath, 0755)
		if err != nil {
			t.Fatalf(err.Error())
			panic(err)
		}
	}
	server.SetPath(MongoPath)
	resultSession := server.Session()
	initData(resultSession)
	return resultSession
}

func (server *TestMongoServer) StopMongo(t *testing.T) {
	server.Stop()
	os.RemoveAll(MongoPath)
}

func initData(mongoSession *mgo.Session) {
	session := mongoSession.Copy()
	defer session.Close()

}
