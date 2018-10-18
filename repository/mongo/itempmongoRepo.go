package mongo

type ITempMongoRepo interface {
	Get() (*TestMongo)
}
