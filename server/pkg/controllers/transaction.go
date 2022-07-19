package controllers

import (
  "esc/ascendaRoyaltyPoint/pkg/models"
  "context"
  "database/sql"
  "github.com/gin-gonic/gin"
  "net/http"
  "time"
  "fmt"
)

func processReward (promotion models.Promotion,base float64) float64{
  if promotion.EarnRateType=="add"{
    return base+promotion.Constant
  }else{
    return base*promotion.Constant

  }
}

func CalculateReward(c context.Context, query *models.Queries, body models.TransferParams)(result float64 ,err error){
  program,err:= query.GetLoyaltyByID(c,int64(body.ProgramId))
  if err!=nil{
		return 0,err
  }

  getPromoParam := models.GetPromotionByDateRangeParams{
    Column1:time.Now().Format("2006-01-02"),
    Program:int32(program.ID), 
  }

  fmt.Println(getPromoParam)
  promotions,err :=query.GetPromotionByDateRange(c,getPromoParam)
  if err!=nil{
    return 0,err
  }

  user ,err := query.GetUserByID(c,int64(body.UserId))
  if err!=nil{
    return 0,err
  }
  var results =[]float64{}
  var base float64= program.InitialEarnRate*body.CreditToTransfer
  fmt.Println(promotions)
  for _,promotion := range promotions{
    var tempReward float64 = 0
    fmt.Println(promotion)
    if (promotion.PromoType=="onetime" ){
      args:=models.GetCreditRequestByPromoParams {
        Program:int32(program.ID),
        PromoUsed:sql.NullInt32{Valid:true,Int32:int32(promotion.ID)},

      }
      _,err= query.GetCreditRequestByPromo(c,args)
      fmt.Println(err.Error())

      //skip the loop if there is result found
      if err.Error()!="sql: no rows in result set"{
        fmt.Println("no pass request made")
        continue
      }
    }
    
    if promotion.CardTier.Valid && user.CardTier.Valid{
      if promotion.CardTier.Int32==user.CardTier.Int32{
        tempReward = processReward(promotion,base)
      }
    }else{
      tempReward = processReward(promotion,base)

    }
    fmt.Println(tempReward)

    if tempReward!=0{
      results = append(results, tempReward)
    }
  } 
  var max float64=base
  if len(results)>0{
    max = results[0]
    for i:=1;i<len(results);i++{
      if results[i]>max{
        max = results[i]
      }
    }
  }
  return max,nil


}

func (server *Server) CheckRewardRate(c *gin.Context){
  body := models.TransferParams{}
  if err:=c.ShouldBindJSON(&body);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
  }
  
  amount,err:=CalculateReward(c,server.store.Queries,body)
  if err!=nil{
    c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
    return
  }
  c.JSON(http.StatusOK,gin.H{"Amount":amount})


}


func (server *Server)CreateTransaction(c *gin.Context){
  //Reading body for the transaction details
  //We will need the membership program, the amount of credit and userid and membership id
  //validate memebership
  //check that credit balance is more than what is requested if not abort
  //compute the amount of thrid party loyalty points to award base on whether there is a promotion on going
  body := models.TransferParams{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//setting the reward should received in the backend instead of pass from frontend
	program ,err:= server.store.Queries.GetLoyaltyByID(c,int64(body.ProgramId))
  if err!=nil{
    c.JSON(http.StatusNotFound,gin.H{"error":err.Error()})
    return
  }
  valid := ValidateMembershipNo(program.FormatRegex,body.MembershipId)
  if valid==false{
    c.JSON(http.StatusBadRequest,gin.H{"error":"Membership not valid"})
    return
  }
  body.RewardShouldReceive,err = CalculateReward(c,server.store.Queries,body)
	
	creditRequest,err := server.store.CreditTransferOut(c,body)
	if err!=nil{
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
	}
  c.JSON(http.StatusOK,creditRequest)
}

