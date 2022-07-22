package controllers

import (
	"esc/ascendaRoyaltyPoint/pkg/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransStatus struct {
	CreditRequestId   int32  `json:"credit_request_id"`
	TransactionStatus string `json:"transaction_status"`
	ProgramId         int32  `json:"program"`
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
	result := GetTransStatus(userRequests)
	fmt.Println(result)
	c.JSON(http.StatusOK, result)
}

func GetTransStatus(creditRequests []models.CreditRequest) []TransStatus {
	var result []TransStatus
	for _, creditRequest := range creditRequests {
		result = append(result, TransStatus{CreditRequestId: int32(creditRequest.ReferenceNumber), TransactionStatus: string(creditRequest.TransactionStatus), ProgramId: creditRequest.Program})
	}
	return result
}
