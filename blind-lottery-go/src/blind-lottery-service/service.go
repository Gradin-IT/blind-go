package main

import (
	"blind-lottery-go/src/blind-lottery-service/model"
	"blind-lottery-go/src/blind-lottery-service/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.Use(cors.Default())
	engine.GET("/hunter", getHunters)
	engine.POST("/hunter", addHunter)
	engine.DELETE("/hunter", deleteHunter)
	engine.PUT("/hunter", updateHunter)
	engine.GET("/blind", getBlinds)
	engine.POST("/blind", addBlind)
	engine.DELETE("/blind", deleteBlind)
	engine.POST("/lottery", getBlindLotteryResult)
	engine.Run("localhost:9090")
}

func getBlindLotteryResult(context *gin.Context) {

	context.IndentedJSON(http.StatusOK, nil)
}

func getHunters(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, model.GetHunters())
}
func updateHunter(context *gin.Context) {
	var hunter model.Hunter
	var hunters = model.GetHunters()
	context.BindJSON(&hunter)
	if index, err := util.IndexOf(hunters, hunter); err == nil {
		hunters[index] = hunter
	}
	context.IndentedJSON(http.StatusOK, hunters)
}
func deleteHunter(context *gin.Context) {
	var hunter model.Hunter
	var hunters = model.GetHunters()
	context.BindJSON(&hunter)
	if index, err := util.IndexOf(hunters, hunter); err == nil {
		hunters = util.RemoveIndex(hunters, index)
	}
	context.IndentedJSON(http.StatusOK, hunters)
}
func addHunter(context *gin.Context) {
	var hunter model.Hunter
	var hunters = model.GetHunters()
	context.BindJSON(&hunter)
	hunters = append(hunters, hunter)
}

func getBlinds(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, model.GetBlinds())
}
func deleteBlind(context *gin.Context) {
	var blind model.Blind
	var blinds = model.GetBlinds()
	context.BindJSON(&blind)
	if index, err := util.IndexOf(blinds, blind); err == nil {
		blinds = util.RemoveIndex(blinds, index)
	}
	context.IndentedJSON(http.StatusOK, blinds)
}
func addBlind(context *gin.Context) {
	var blind model.Blind
	var blinds = model.GetBlinds()
	context.BindJSON(&blind)
	blinds = append(blinds, blind)
	context.IndentedJSON(http.StatusOK, blinds)
}
