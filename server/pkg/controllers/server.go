package controllers

import(
  "esc/ascendaRoyaltyPoint/pkg/models"
  "github.com/gin-contrib/cors"
  "github.com/gin-gonic/gin"
)
type Server struct{
  store *models.Store
  router *gin.Engine
}

func (server *Server) Start(address string)error{
  return server.router.Run(address)

}

func NewServer(store *models.Store) *Server{
  server := &Server{
    store:store,
  }
  router := gin.Default()

  config := cors.DefaultConfig()
  config.AllowAllOrigins = true
  router.Use(cors.New(config))
  router.POST("/initTransaction",server.CreateTransaction)
  router.POST("/loyalty/validateMembership", server.CheckLoyaltyRegEx)
  router.POST("/loyalty", server.CreateLoyaltyProg)
  router.GET("/loyalty", server.GetLoyalty)
  router.GET("/loyalty/:id", server.GetLoyaltyId)
  router.POST("/checkReward",server.CheckRewardRate)
  server.router=router
  return server
}
