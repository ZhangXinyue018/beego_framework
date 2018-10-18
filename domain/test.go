package domain

type TestMysql struct {
	Id   int    `xorm:"'id' pk autoincr" json:"id"`
	Test string `xorm:"'test'" json:"test"`
}
