package impl

import (
	"beego_framework/domain"
	"reflect"
	"testing"
)

func TestTempMysqlRepository_ListAll(t *testing.T) {
	engine := StartMysql(t)
	defer StopMysql(engine)
	tests := []struct {
		name   string
		want   *[]domain.TestMysql
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &TempMysqlRepository{
				TempMysqlReadEngine:  engine,
				TempMysqlWriteEngine: engine,
			}
			if got := repository.ListAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TempMysqlRepository.ListAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
