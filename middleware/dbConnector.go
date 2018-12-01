package middleware

import (
	"IoT-admin-backend/db"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 数据库连接中间件：克隆每一个数据库会话，并且确保 `db` 属性在每一个 handler 里均有效
func Connect(context *gin.Context) {
	s := db.Session.Clone()
	defer s.Clone()

	context.Set("db", s.DB(db.Mongo.Database))
	context.Next()

}

const (
	APPLICATION_JSON = "application/json"
)

// 错误处理中间件
func ErrorHandler(context *gin.Context) {
	context.Next()

	// TODO
	if len(context.Errors) > 0 {
		ct := context.Request.Header.Get("Content-Type")
		if strings.Contains(ct, APPLICATION_JSON) {
			context.JSON(http.StatusBadRequest, gin.H{"error": context.Errors})
		} else {
			context.HTML(http.StatusBadRequest, "400", gin.H{"error": context.Errors})
		}
	}
}
