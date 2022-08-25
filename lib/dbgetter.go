package lib

import (
	"encoding/json"
	"goft-redis/gedis"
	"log"
)

func NewsDBGetter(id string)  gedis.DBGetterFunc {  // 装饰器模式
	log.Println("get from db")
	return func() string {
		newsModel := NewNewsModel()
		Gorm.Table("mynews").Where("id=?",id).Find(newsModel)
		b, _ := json.Marshal(newsModel)
		return string(b)
	}
}