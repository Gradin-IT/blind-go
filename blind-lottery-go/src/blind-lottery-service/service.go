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
	engine.GET("/blind/:area", getBlindsByArea)
	engine.POST("/blind", addBlind)
	engine.DELETE("/blind", deleteBlind)
	engine.GET("/area", getAreas)
	engine.POST("/lottery", getBlindLotteryResult)
	engine.Static("/web", "./public")
	engine.Run("localhost:9090")
}

func getAreas(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, model.GetAreas())
}

func getBlindsByArea(context *gin.Context) {
	println("getBlindsByArea start ")
	var blinds = []model.Blind{}
	area := context.Param("area")
	println("area: " + area)
	var blindIds []int
	for _, a := range model.GetAreas() {
		id, _ := strconv.Atoi(area)
		if id == a.Id {
			println(a.Id)
			println(a.Name)
			println(a.BlindIds)
			blindIds = a.BlindIds
		}
	}
	println("blinds for area: ", blindIds)
	var blindInfo = model.GetBlinds()
	for _, blind := range blindInfo {
		id, _ := strconv.Atoi(blind.ID)
		if util.Contains(blindIds, id) {
			blinds = append(blinds, blind)
		}

	}
	println("getBlindsByArea end", blinds)
	context.IndentedJSON(http.StatusOK, blinds)
}

type Query struct {
	Hunters []string `json:"hunters"`
	Blinds  []string `json:"blinds"`
}

func getBlindLotteryResult(context *gin.Context) {
	var query Query
	context.BindJSON(&query)
	var hunterInfo = model.GetHunters()
	var blindInfo = model.GetBlinds()
	hunters := util.Randomize(query.Hunters)
	blinds := util.Randomize(query.Blinds)

	var results []model.Result
	for i, hunter := range hunters {
		var result = model.Result{}
		hunterIndex, _ := strconv.Atoi(hunter)
		result.HunterName = hunterInfo[hunterIndex-1].Name
		result.BlindName = getBlindName(blinds[i], blindInfo)
		results = append(results, result)
	}
	model.By(model.Name).Sort(results)
	context.IndentedJSON(http.StatusOK, results)
}

func getBlindName(s string, info []model.Blind) string {
	blindIndex, _ := strconv.Atoi(s)
	return info[blindIndex-1].ID + " ," + info[blindIndex-1].Description
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
