package controllers

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func ValidateMembershipNo(regexstr string, membernum string) bool {
	Re := regexp.MustCompile(regexstr)
	return Re.MatchString(membernum)

}

type regexCheckParams struct {
	ProgramId    int64  `json:"programId" binding:"required"`
	StringToTest string `json:"stringToTest" binding:"required"`
}

func (server *Server) CheckLoyaltyRegEx(c *gin.Context) {
	body := &regexCheckParams{}
	err := c.ShouldBindJSON(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	regexform, err := server.store.Queries.GetRegEx(c, body.ProgramId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	valid := ValidateMembershipNo(regexform, body.StringToTest)
	c.JSON(http.StatusOK, gin.H{"valid": valid})
}
