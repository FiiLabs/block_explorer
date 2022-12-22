package rest

import (
	"github.com/FiiLabs/block_explorer/api/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BlockController struct {
}

func (bctl *BlockController) QueryByHeight(c *gin.Context) {
	blkHeight := c.Param("block_height")
	if blkHeight == "" {
		c.JSON(http.StatusBadRequest, response.FailBadRequest("parameter block_height is required"))
		return
	}

	res, e := blockService.QueryBlockByHeight(blkHeight)
	if e != nil {
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	}

	c.JSON(http.StatusOK, response.Success(res))
}

func (bctl *BlockController) QueryByHash(c *gin.Context) {
	blkHash := c.Param("block_hash")
	if blkHash == "" {
		c.JSON(http.StatusBadRequest, response.FailBadRequest("parameter block_hash is required"))
		return
	}

	res, e := blockService.QueryBlockByHash(blkHash)
	if e != nil {
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	}

	c.JSON(http.StatusOK, response.Success(res))
}

func (bctl *BlockController) QueryBlocks(c *gin.Context) {
	q_num := c.DefaultQuery("num", "10")

	res, e := blockService.QueryBlocks(q_num)
	if e != nil {
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	}

	c.JSON(http.StatusOK, response.Success(res))
}

func (bctl *BlockController) QueryLBlocks(c *gin.Context) {

	res, e := blockService.QueryLBlocks()
	if e != nil {
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	}

	c.JSON(http.StatusOK, response.Success(res))
}
