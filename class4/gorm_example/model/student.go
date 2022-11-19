package model

type Student struct {
	Id   int    `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` //column指定在数据库中对应的字段
	Name string `gorm:"column:name" json:"name"`
	Sex  string `gorm:"column:sex" json:"sex"`
	Age  string `gorm:"column:age" json:"age"`
}
