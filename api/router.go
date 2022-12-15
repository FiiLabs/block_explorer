package api

import (
	"github.com/FiiLabs/block_explorer/api/rest"
	docs "github.com/FiiLabs/block_explorer/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)
// @BasePath /abc

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /abc [get]
func Helloworld(g *gin.Context)  {
	g.JSON(http.StatusOK,"helloworld")
}
func Routers(Router *gin.Engine) {
	docs.SwaggerInfo.BasePath = ""
	Router.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})
	Router.GET("/abc", Helloworld)
	blockRouter := Router.Group("block")
	blkCtl(blockRouter)
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func blkCtl(r *gin.RouterGroup) {
	bctl := rest.BlockController{}
	r.GET("/height/:block_height", bctl.Query)
}