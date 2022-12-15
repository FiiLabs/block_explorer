package rest

import (
	"github.com/FiiLabs/block_explorer/api/response"
	"github.com/FiiLabs/block_explorer/model/vo"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BlockController struct {
}

func (bctl *BlockController) Query(c *gin.Context) {
	txHash := c.Param("block_height")
	if txHash == "" {
		c.JSON(http.StatusBadRequest, response.FailBadRequest("parameter block_height is required"))
		return
	}

	var req vo.BlockReq
	res, e := blockService.Query(txHash, req)
	if e != nil {
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	}

	c.JSON(http.StatusOK, response.Success(res))
}