package controllers

import(
  "github.com/gin-gonic/gin"
  "net/http"
  "esc/ascendaRoyaltyPoint/pkg/models"
  
)
var LoyaltyProg models.LoyaltyProgram
var CreateLoyalty = func(c *gin.Context){
  prog := &models.LoyaltyProgram{}
	if err := c.ShouldBindJSON(prog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// prog.CreateLoyaltyProgram()
	c.JSON(http.StatusCreated,prog)


}
var GetLoyalty = func (c *gin.Context){
  // progs := models.GetAllProg()
  // c.JSON(http.StatusOK,progs)
}

