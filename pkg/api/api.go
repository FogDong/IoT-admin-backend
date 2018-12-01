package api

import (
	"github.com/gin-gonic/gin"
)

func RunHTTPServer(engine *gin.Engine) {
	GetUserAPI(engine)
}
