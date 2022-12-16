package api

import (
	"github.com/FiiLabs/block_explorer/api/rest"
	docs "github.com/FiiLabs/block_explorer/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func Routers(Router *gin.Engine) {
	docs.SwaggerInfo.BasePath = ""
	Router.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})
	blockRouter := Router.Group("block")
	blkCtl(blockRouter)
	txRouter := Router.Group("tx")
	txCtl(txRouter)
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func blkCtl(r *gin.RouterGroup) {
	bctl := rest.BlockController{}
	r.GET("/height/:block_height", bctl.QueryByHeight)
	r.GET("/hash/:block_hash", bctl.QueryByHash)
	r.GET("/blocks", bctl.QueryBlocks)
}

func txCtl(r *gin.RouterGroup) {
	txctl := rest.TxController{}
	r.GET("/hash/:tx_hash", txctl.QueryByHash)
	r.GET("/block_txs/:block_height", txctl.QueryTxsByBlockHeight)
}