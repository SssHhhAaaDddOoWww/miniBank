package handler

import (
	"net/http"
	"strconv"

	"github.com/SssHhhAaaDddOoWww/miniBank/internal/services"
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

func GetBalance(c *gin.Context) {
	idParam := c.Param("id")
	accountID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid account id"})
		return
	}
	balance, err := services.GetBalance(uint(accountID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"account_id": accountID,
		"balance":    balance,
	})

}
func GetLedger(c *gin.Context) {
	idParam := c.Param("id")

	accountID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid account id"})
		return
	}

	entries, err := services.GetLedger(uint(accountID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"account_id": accountID,
		"entries":    entries,
	})
}
