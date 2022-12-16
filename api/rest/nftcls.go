package rest

import (
	"github.com/FiiLabs/block_explorer/api/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type NFTClsController struct {
}

func (nctl *NFTClsController) QueryNFTCls(c *gin.Context) {
	q_page := c.DefaultQuery("page", "0")
	q_size := c.DefaultQuery("size", "10")

	res, e := nftClsService.QueryNFTCls(q_page, q_size)
	if e != nil {
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	}

	c.JSON(http.StatusOK, response.Success(res))
}