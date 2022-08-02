package controllers

import (
	"context"
	"esc/ascendaRoyaltyPoint/pkg/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransStatus struct {
	CreditRequestId   int32   `json:"credit_request_id"`
	TransactionStatus string  `json:"transaction_status"`
	Program           string  `json:"program"`
	CreditUsed        float64 `json:"credit_used"`
	CreditToReceive   float64 `json:"credit_to_receive"`
}

// get all transaction status from user
func (server *Server) GetAllCreditRequest(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userRequests, err := server.store.GetCreditRequestByUser(c, int32(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if userRequests == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No request found"})
		return
	}

	result := GetTransStatus(userRequests, server.store.Queries)
	c.JSON(http.StatusOK, result)
}

func GetTransStatus(creditRequests []models.CreditRequest, server *models.Queries) []TransStatus {
	var result []TransStatus
	for _, creditRequest := range creditRequests {
		program, _ := server.GetLoyaltyByID(context.Background(), int64(creditRequest.Program))
		// if err != nil {
		// 	return result
		// }
		result = append(result, TransStatus{CreditRequestId: int32(creditRequest.ReferenceNumber), TransactionStatus: string(creditRequest.TransactionStatus), Program: program.Name, CreditUsed: creditRequest.CreditUsed, CreditToReceive: creditRequest.RewardShouldReceive})
	}
	return result
}
