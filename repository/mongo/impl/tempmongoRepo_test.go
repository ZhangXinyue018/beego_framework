package impl

import (
	"reflect"
	"testing"
)

func TestTempMongoRepository_Get(t *testing.T) {
	var server TestMongoServer
	defer server.StopMongo(t)
	mongoSession := server.StartMongo(t)
	defer mongoSession.Close()
	tests := []struct {
		name string
		want *TestMongo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &TempMongoRepository{
				TempMongoSession:   mongoSession,
				TempDBName:         "db",
				TempCollectionName: "collection",
			}
			if got := repository.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TempMongoRepository.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
