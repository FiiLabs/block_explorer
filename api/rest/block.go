package rest

import (
	"github.com/FiiLabs/block_explorer/api/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BlockController struct {
}

// @Summary 通过高度查询区块
// @Description 通过高度查询区块
// @Param block_height path string true "height"
// @Tags block
// @Success 200 {object} vo.BlockResp
// @Router /block/height/{block_height} [get]
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
// @Summary 通过hash查询区块
// @Description 通过hash查询区块
// @Param block_hash path string true "hash"
// @Tags block
// @Success 200 {object} vo.BlockResp
// @Router /block/hash/{block_hash} [get]
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
// @Summary 查询最近产生的区块
// @Description 查询最近产生的区块,最大可查到50个区块
// @Param num query		int	false	"number"
// @Tags block
// @Success 200 {object} vo.BlocksResp
// @Router /block/blocks [get]
func (bctl *BlockController) QueryBlocks(c *gin.Context) {
	q_num := c.DefaultQuery("num", "10")

	res, e := blockService.QueryBlocks(q_num)
	if e != nil {
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	}

	c.JSON(http.StatusOK, response.Success(res))
}
// @Summary 查询最近10个区块
// @Description 查询最近10个区块，用在首页展示
// @Tags block
// @Success 200 {object} vo.BlocksResp
// @Router /block/lblocks [get]
func (bctl *BlockController) QueryLBlocks(c *gin.Context) {

	res, e := blockService.QueryLBlocks()
	if e != nil {
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	}

	c.JSON(http.StatusOK, response.Success(res))
}
