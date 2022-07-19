package controllers

import(
  "github.com/gin-gonic/gin"
  "net/http"
  "strconv"
  "esc/ascendaRoyaltyPoint/pkg/models"
)
func(server *Server)CreateLoyaltyProg(c *gin.Context){
  prog := &models.CreateLoyaltyParams{}
	if err := c.ShouldBindJSON(prog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	programCreated,err:=server.store.Queries.CreateLoyalty(c,*prog)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated,programCreated)
}
func (server *Server) GetLoyalty(c *gin.Context){
  progs,err:=server.store.Queries.ListLoyalty(c)
  if err!=nil{
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
  }
  c.JSON(http.StatusCreated,progs)
}
func (server *Server) GetLoyaltyId(c *gin.Context){
  id,err := strconv.Atoi(c.Param("id"))
  if err!=nil{
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
  }
  
  prog,err:=server.store.Queries.GetLoyaltyByID(c,int64(id))
  if err!=nil{
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
  }
  c.JSON(http.StatusCreated,prog)
}


