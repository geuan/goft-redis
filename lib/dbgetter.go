package lib

import (
	"goft-redis/gedis"
	"log"
)

func NewsDBGetter(id string)  gedis.DBGetterFunc {  // 装饰器模式
	return func() interface{} {
		log.Println("get from db")
		newsModel := NewNewsModel()
		Gorm.Table("mynews").Where("id=?",id).Find(newsModel)
		return newsModel

	}
}