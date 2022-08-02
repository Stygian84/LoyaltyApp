package controllers

import (
	"esc/ascendaRoyaltyPoint/pkg/models"
	"net/http"
	

	"github.com/gin-gonic/gin"
)

func(server *Server)CreateUser(c *gin.Context){
	user := &models.CreateUserParams{}
	  if err := c.ShouldBindJSON(user); err != nil {
		  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		  return
	  }
	  
	  userCreated,err:=server.store.Queries.CreateUser(c,*user)
	  if err!=nil{
		  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		  return
	  }
	  c.JSON(http.StatusCreated,userCreated)
  }


func (server *Server) GetUserByUserName(c *gin.Context){
	username := c.Param("username")
	
	user,err:=server.store.Queries.GetUserByUserName(c,username)
	if err!=nil{
	  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		  return
	}
	
	c.JSON(http.StatusOK,user)
  }

  func (server *Server) GetUserByEmail(c *gin.Context){
	email := c.Param("email")
	
	users,err:=server.store.Queries.GetUserByUserEmail(c,email)
	if err!=nil{
	  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		  return
	}
	
	c.JSON(http.StatusOK,users)
  }
  
