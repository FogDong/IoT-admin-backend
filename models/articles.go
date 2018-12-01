package models

import "gopkg.in/mgo.v2/bson"

const (
	// 保存集合名称
	CollectionArticle = "article"
)

type Article struct {
	Id        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`                 // mongodb id
	Title     string        `json:"title" form:"title" binding:"required" bson:"title"` // 标题
	Body      string        `json:"body" form:"body" binding:"required" bson:"body"`    // 内容
	CreatedOn int64         `json:"created_on" bson:"created_on"`                       // 创建时间
	UpdatedOn int64         `json:"updated_on" bson:"updated_on"`                       // 修改时间
}