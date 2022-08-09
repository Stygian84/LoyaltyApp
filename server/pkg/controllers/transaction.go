package controllers

import (
	"context"
	"database/sql"
	"esc/ascendaRoyaltyPoint/pkg/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func processReward(promotion models.Promotion, base float64) float64 {
	if promotion.EarnRateType == "add" {
		return base + promotion.Constant
	} else {
		return base * promotion.Constant

	}
}

func CalculateReward(c context.Context, query *models.Queries, body models.TransferParams) (result float64, promoUsed sql.NullInt32, err error) {
	if body.CreditToTransfer < 0 {
		// return 0,error
	}
	program, err := query.GetLoyaltyByID(c, int64(body.ProgramId))
	if err != nil {
		return 0, sql.NullInt32{Valid: false}, nil
	}

	getPromoParam := models.GetPromotionByDateRangeParams{
		Column1: time.Now().Format("2006-01-02"),
		Program: int32(program.ID),
	}

	promotions, err := query.GetPromotionByDateRange(c, getPromoParam)
	if err != nil {
		return 0, sql.NullInt32{Valid: false}, nil
	}

	user, err := query.GetUserByID(c, int64(body.UserId))
	if err != nil {
		return 0, sql.NullInt32{Valid: false}, nil
	}
	var base float64 = program.InitialEarnRate * body.CreditToTransfer
	var promoIdUsed int32 = 0
	var max float64 = base
	for _, promotion := range promotions {
		var tempReward float64 = 0
		if promotion.PromoType == "onetime" {
			args := models.GetCreditRequestByPromoParams{
				Program:   int32(program.ID),
				PromoUsed: sql.NullInt32{Valid: true, Int32: int32(promotion.ID)},
			}
			request, err := query.GetCreditRequestByPromo(c, args)
			if len(request) > 0 {
				continue
			}
			//skip the loop if there is result found

			if err != nil {
				fmt.Println(err)
			}

		}

		fmt.Println(promotion.CardTier + ' ' + user.CardTier)
		if promotion.CardTier!=0{
		if promotion.CardTier == user.CardTier {
			tempReward = processReward(promotion, base)
		} else {
			continue

		}}
		tempReward = processReward(promotion, base)
		if tempReward != 0 {
			if tempReward > max {
				max = tempReward
				promoIdUsed = int32(promotion.ID)
			}
		}
	}
	if promoIdUsed != 0 {
		return max, sql.NullInt32{Int32: promoIdUsed, Valid: true}, nil
	} else {
		return max, sql.NullInt32{Valid: false}, nil
	}
}

func (server *Server) CheckRewardRate(c *gin.Context) {
	body := models.TransferParams{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	amount, _, err := CalculateReward(c, server.store.Queries, body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Amount": amount})
}

func (server *Server) CreateTransaction(c *gin.Context) {
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
	program, err := server.store.Queries.GetLoyaltyByID(c, int64(body.ProgramId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	valid := ValidateMembershipNo(program.FormatRegex, body.MembershipId)
	if valid == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Membership not valid"})
		return
	}
	var promoUsed sql.NullInt32
	body.RewardShouldReceive, promoUsed, err = CalculateReward(c, server.store.Queries, body)

	creditRequest, err := server.store.CreditTransferOut(c, body, promoUsed)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, creditRequest)
}
