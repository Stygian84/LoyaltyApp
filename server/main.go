package main

import (
  "github.com/gin-gonic/gin"
  "fmt"
  "esc/ascendaRoyaltyPoint/pkg/routes"
)

func main(){
  fmt.Println("Hello")
  r := gin.Default()
  routes.RegisterLoyaltyProgRoutes(r)
  
  r.Run()

}

