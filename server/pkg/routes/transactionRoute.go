package routes

import(
  "esc/ascendaRoyaltyPoint/pkg/controllers"
  "github.com/gin-gonic/gin"
)

var RegisterTransactionRoute = func(router gin.IRoutes){
  router.POST("/initTransaction",controllers.CreateTransaction)
}
