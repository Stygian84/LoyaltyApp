package main

import (
  "github.com/gin-gonic/gin"
  "esc/ascendaRoyaltyPoint/pkg/routes"
)

func main(){

  r := gin.Default()
  routes.RegisterLoyaltyProgRoutes(r)
  
  r.Run()

}

