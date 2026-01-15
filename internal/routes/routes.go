package routes

import (
	"github.com/SssHhhAaaDddOoWww/miniBank/internal/handler"
	"github.com/SssHhhAaaDddOoWww/miniBank/internal/routes/health"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/health", health.Health)
	router.POST("/accounts", handler.CreateAccount)
	router.GET("/accounts/:id/balance", handler.GetBalance)
	router.POST("/deposit", handler.Deposit)
	router.POST("/withdraw", handler.Withdraw)
	router.POST("/transfer", handler.Transfer)

}
