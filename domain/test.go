package domain

type TestMysql struct {
	Id   int    `xorm:"pk" json:"id"`
	Test string `xorm:"'test'" json:"test"`
}
