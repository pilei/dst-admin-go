package api

import (
	"dst-admin-go/service"
	"dst-admin-go/utils/clusterUtils"
	"dst-admin-go/vo"
	"dst-admin-go/vo/level"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GameApi struct {
}

var gameService = service.GameService{}

func (g *GameApi) UpdateGame(ctx *gin.Context) {

	log.Println("正在更新游戏。。。。。。")
	cluster := clusterUtils.GetClusterFromGin(ctx)
	clusterName := cluster.ClusterName

	err := gameService.UpdateGame(clusterName)
	if err != nil {
		log.Panicln("更新游戏失败: ", err)
	}

	ctx.JSON(http.StatusOK, vo.Response{
		Code: 200,
		Msg:  "update dst success",
		Data: nil,
	})
}

func (g *GameApi) StartGame(ctx *gin.Context) {

	opType, _ := strconv.Atoi(ctx.DefaultQuery("type", "0"))

	cluster := clusterUtils.GetClusterFromGin(ctx)
	bin := cluster.Bin
	beta := cluster.Beta
	clusterName := cluster.ClusterName
	log.Println("正在启动指定游戏服务 type:", clusterName, opType, "bin:", bin, "beta: ", beta)
	gameService.StartGame(clusterName, bin, beta, opType)

	ctx.JSON(http.StatusOK, vo.Response{
		Code: 200,
		Msg:  "start " + clusterName + " level success",
		Data: nil,
	})
}

func (g *GameApi) StopGame(ctx *gin.Context) {

	opType, _ := strconv.Atoi(ctx.DefaultQuery("type", "0"))
	cluster := clusterUtils.GetClusterFromGin(ctx)
	clusterName := cluster.ClusterName
	log.Println("正在停止指定游戏服务 type:", clusterName, opType)

	gameService.StopGame(clusterName, opType)

	ctx.JSON(http.StatusOK, vo.Response{
		Code: 200,
		Msg:  "stop " + clusterName + " level success",
		Data: nil,
	})
}

func (g *GameApi) GetDashboardInfo(ctx *gin.Context) {

	cluster := clusterUtils.GetClusterFromGin(ctx)
	clusterName := cluster.ClusterName
	ctx.JSON(http.StatusOK, vo.Response{
		Code: 200,
		Msg:  "success",
		Data: gameService.GetClusterDashboard(clusterName),
	})
}

func (g *GameApi) GetGameConfig(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, vo.Response{
		Code: 200,
		Msg:  "success",
		Data: gameService.GetGameConfig(ctx),
	})
}

func (g *GameApi) SaveGameConfig(ctx *gin.Context) {

	gameConfig := level.GameConfig{}
	ctx.ShouldBind(&gameConfig)
	gameService.SaveGameConfig(ctx, &gameConfig)

	ctx.JSON(http.StatusOK, vo.Response{
		Code: 200,
		Msg:  "success",
		Data: nil,
	})
}
