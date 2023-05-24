package main

import (
	"blind-lottery-go/src/blind-lottery-service/model"
	"blind-lottery-go/src/blind-lottery-service/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	engine.Static("/web", "./public")
	engine.Run("localhost:9090")
}

type Query struct {
	Hunters []int `json:"hunters"`
	Blinds  []int `json:"blinds"`
}

func getBlindLotteryResult(context *gin.Context) {
	var query Query
	context.BindJSON(&query)
	var hunterInfo = model.GetHunters()
	var blindInfo = model.GetBlinds()
	for _, b := range query.Blinds {
		println(b)
	}
	for _, h := range query.Hunters {
		println(h)
	}
	hunters := util.Randomize(query.Hunters)
	blinds := util.Randomize(query.Blinds)
	for _, b := range blinds {
		println(b)
	}
	for _, h := range hunters {
		println(h)
	}
	var results []model.Result
	for i, hunter := range hunters {
		println(results)
		var result = model.Result{}
		result.HunterName = hunterInfo[hunter].Name
		result.BlindName = getBlindName(blinds[i], blindInfo)
		results = append(results, result)
	}
	context.IndentedJSON(http.StatusOK, results)
}

func getBlindName(id int, blinds []model.Blind) string {
	var blindName = strconv.Itoa(blinds[id].ID) + " " + blinds[id].Description
	println(blindName)
	return blindName
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
