package handler

import (
	"net/http"

	"github.com/SssHhhAaaDddOoWww/miniBank/api/services"
	"github.com/gin-gonic/gin"
)

func Deposit(c *gin.Context) {
	var req struct {
		AccountID uint    `json:"account_id"`
		Amount    float64 `json:"amount"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.Deposit(req.Amount, req.AccountID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	c.JSON(http.StatusOK, gin.H{"status": "withdraw successful"})

}

func Withdraw(c *gin.Context) {
	var req struct {
		AccountID uint    `json:"account_id"`
		Amount    float64 `json:"amount"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.Withdraw(req.Amount, req.AccountID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	c.JSON(http.StatusOK, gin.H{"status": "withdrawl succesfull !!"})
}

func Transfer(c *gin.Context) {
	var req struct {
		FromAccountID uint    `json:"from_id"`
		ToAccountID   uint    `json:"to_id"`
		Amount        float64 `json:"amount"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err := services.Transfer(req.FromAccountID, req.ToAccountID, req.Amount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"status": "transferred succesfully !"})
}
