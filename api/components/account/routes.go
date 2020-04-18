package account

import "github.com/gin-gonic/gin"

func Routes(route *gin.RouterGroup) {
	route.GET("/", HandleGetAccount)
	accountRouter := route.Group("/transaction")
	accountRouter.GET("/", HandleGetTransactions)
	accountRouter.POST("/", HandleModification)
	accountRouter.GET("/:transactionId", HandleGetTransaction)
}
