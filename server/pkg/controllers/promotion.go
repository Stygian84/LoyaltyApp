package controllers

import (
	"esc/ascendaRoyaltyPoint/pkg/models"
	"net/http"
	

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