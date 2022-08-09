package controllers

import (
	"esc/ascendaRoyaltyPoint/pkg/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func(server *Server)CreatePromo(c *gin.Context){
	promo := &models.CreatePromotionParams{}
	  if err := c.ShouldBindJSON(promo); err != nil {
		  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		  return
	  }
	  
	  promoCreated,err:=server.store.Queries.CreatePromotion(c,*promo)
	  if err!=nil{
		  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		  return
	  }
	  c.JSON(http.StatusCreated,promoCreated)
  }


func (server *Server) GetPromoList(c *gin.Context){
promos,err:=server.store.Queries.ListPromotion(c)
if err!=nil{
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
}
c.JSON(http.StatusOK,promos)
}

func (server *Server) GetPromoCurrent(c *gin.Context){
	progID, err:=strconv.Atoi(c.Param("progid"))
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	  }
	getPromoParam := models.GetPromotionByDateRangeParams{
		Column1: time.Now().Format("2006-01-02"),
		Program: int32(progID),
	}

	promotions, err := server.store.Queries.GetPromotionByDateRange(c, getPromoParam)
	
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK,promotions)
	}